package subfunction

import (
	api_input_reader "data-platform-api-delivery-document-items-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-delivery-document-items-creates-subfunc/API_Processing_Data_Formatter"
	"sort"
	"strings"
	"time"

	"golang.org/x/xerrors"
)

func (f *SubFunction) OrdersHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrdersHeader, error) {
	args := make([]interface{}, 0)

	orderItem := psdc.OrderItem
	repeat := strings.Repeat("?,", len(orderItem)-1) + "?"
	for _, v := range orderItem {
		args = append(args, v.OrderID)
	}

	rows, err := f.db.Query(
		`SELECT OrderID, OrderType, SupplyChainRelationshipID, SupplyChainRelationshipBillingID,
		SupplyChainRelationshipPaymentID, Buyer, Seller, BillToParty, BillFromParty, BillToCountry,
		BillFromCountry, Payer, Payee, ContractType, OrderValidityStartDate, OrderValidityEndDate,
		InvoicePeriodStartDate, InvoicePeriodEndDate, TransactionCurrency, Incoterms, IsExportImport, HeaderText
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE OrderID IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrdersHeader(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) CalculateDeliveryDocument(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.CalculateDeliveryDocument, error) {
	metaData := psdc.MetaData
	dataKey := psdc.ConvertToCalculateDeliveryDocumentKey()

	dataKey.ServiceLabel = metaData.ServiceLabel

	rows, err := f.db.Query(
		`SELECT ServiceLabel, FieldNameWithNumberRange, LatestNumber
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
		WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, dataKey.ServiceLabel, dataKey.FieldNameWithNumberRange,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dataQueryGets, err := psdc.ConvertToCalculateDeliveryDocumentQueryGets(rows)
	if err != nil {
		return nil, err
	}

	if dataQueryGets.DeliveryDocumentLatestNumber == nil {
		return nil, xerrors.Errorf("'data_platform_number_range_latest_number_data'テーブルのLatestNumberがNULLです。")
	}

	deliverPlants := make([]api_processing_data_formatter.DeliverPlant, 0)
	for _, orderItem := range psdc.OrderItem {
		deliverFromPlant := orderItem.DeliverFromPlant
		deliverToPlant := orderItem.DeliverToPlant
		orderID := orderItem.OrderID
		orderItem := orderItem.OrderItem

		if deliverFromPlant == nil || deliverToPlant == nil {
			continue
		}

		if deliverPlantContain(deliverPlants, *deliverFromPlant, *deliverToPlant) {
			continue
		}

		deliverPlants = append(deliverPlants, api_processing_data_formatter.DeliverPlant{
			DeliverFromPlant: *deliverFromPlant,
			DeliverToPlant:   *deliverToPlant,
			OrderID:          orderID,
			OrderItem:        orderItem,
		})
	}

	data := make([]*api_processing_data_formatter.CalculateDeliveryDocument, 0)
	for i, deliverPlant := range deliverPlants {
		deliveryDocumentLatestNumber := dataQueryGets.DeliveryDocumentLatestNumber
		deliveryDocument := *dataQueryGets.DeliveryDocumentLatestNumber + i + 1
		deliverFromPlant := deliverPlant.DeliverFromPlant
		deliverToPlant := deliverPlant.DeliverToPlant
		orderID := deliverPlant.OrderID
		orderItem := deliverPlant.OrderItem

		datum := psdc.ConvertToCalculateDeliveryDocument(deliveryDocumentLatestNumber, deliveryDocument, orderID, orderItem, deliverFromPlant, deliverToPlant)
		data = append(data, datum)
	}

	return data, err
}

func deliverPlantContain(deliverPlants []api_processing_data_formatter.DeliverPlant, deliverFromPlant, deliverToPlant string) bool {
	for _, deliverPlant := range deliverPlants {
		if deliverFromPlant == deliverPlant.DeliverFromPlant && deliverToPlant == deliverPlant.DeliverToPlant {
			return true
		}
	}
	return false
}

func (f *SubFunction) DocumentDate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.DocumentDate {
	documentDate := getStringPtr(getSystemDate())

	if sdc.Header.DocumentDate != nil {
		if *sdc.Header.DocumentDate != "" {
			documentDate = sdc.Header.DocumentDate
		}
	}

	data := psdc.ConvertToDocumentDate(documentDate)

	return data
}

func (f *SubFunction) InvoiceDocumentDate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.InvoiceDocumentDate, error) {
	rows, err := f.db.Query(
		`SELECT PaymentTerms, BaseDate, BaseDateCalcAddMonth, BaseDateCalcFixedDate, PaymentDueDateCalcAddMonth, PaymentDueDateCalcFixedDate
			FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_terms_payment_terms_data
			WHERE PaymentTerms = ?;`, psdc.OrdersItem[0].PaymentTerms,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	psdc.PaymentTerms, err = psdc.ConvertToPaymentTerms(rows)
	if err != nil {
		return nil, err
	}

	if sdc.Header.InvoiceDocumentDate != nil {
		if *sdc.Header.InvoiceDocumentDate != "" {
			data := psdc.ConvertToInvoiceDocumentDate(sdc)
			return data, nil
		}
	}

	plannedGoodsIssueDate, err := psdc.ConvertToPlannedGoodsIssueDate(sdc)
	if err != nil {
		return nil, err
	}

	invoiceDocumentDate, err := calculateInvoiceDocumentDate(psdc, plannedGoodsIssueDate.PlannedGoodsIssueDate, psdc.PaymentTerms)
	if err != nil {
		return nil, err
	}

	data := psdc.ConvertToCaluculateInvoiceDocumentDate(sdc, invoiceDocumentDate)

	return data, err
}

func calculateInvoiceDocumentDate(
	psdc *api_processing_data_formatter.SDC,
	plannedGoodsIssueDate string,
	paymentTerms []*api_processing_data_formatter.PaymentTerms,
) (*string, error) {
	format := "2006-01-02"
	t, err := time.Parse(format, plannedGoodsIssueDate)
	if err != nil {
		return nil, err
	}

	sort.Slice(paymentTerms, func(i, j int) bool {
		return paymentTerms[i].BaseDate < paymentTerms[j].BaseDate
	})

	day := t.Day()
	for i, v := range paymentTerms {
		if day <= v.BaseDate {
			t = time.Date(t.Year(), t.Month()+time.Month(*v.BaseDateCalcAddMonth)+1, 0, 0, 0, 0, 0, time.UTC)
			if *v.BaseDateCalcFixedDate == 31 {
				t = time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC)
			} else {
				t = time.Date(t.Year(), t.Month(), *v.BaseDateCalcFixedDate, 0, 0, 0, 0, time.UTC)
			}
			break
		}
		if i == len(paymentTerms)-1 {
			return nil, xerrors.Errorf("'data_platform_payment_terms_payment_terms_data'テーブルが不適切です。")
		}
	}

	res := getStringPtr(t.Format(format))

	return res, nil
}

func (f *SubFunction) HeaderGrossWeight(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.HeaderGrossWeight {
	headerGrossWeight := new(float32)

	for _, v := range psdc.OrdersItem {
		if v.ItemGrossWeight != nil {
			*headerGrossWeight += *v.ItemGrossWeight
		}
	}

	data := psdc.ConvertToHeaderGrossWeight(headerGrossWeight)

	return data
}

func (f *SubFunction) HeaderNetWeight(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.HeaderNetWeight {
	headerNetWeight := new(float32)

	for _, v := range psdc.OrdersItem {
		if v.ItemNetWeight != nil {
			*headerNetWeight += *v.ItemNetWeight
		}
	}

	data := psdc.ConvertToHeaderNetWeight(headerNetWeight)

	return data
}

func (f *SubFunction) CreationDateHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.CreationDate {
	data := psdc.ConvertToCreationDateHeader(getStringPtr(getSystemDate()))

	return data
}

func (f *SubFunction) LastChangeDateHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.LastChangeDate {
	data := psdc.ConvertToLastChangeDateHeader(getStringPtr(getSystemDate()))

	return data
}

func (f *SubFunction) CreationTimeHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.CreationTime {
	data := psdc.ConvertToCreationTimeHeader(getStringPtr(getSystemTime()))

	return data
}

func (f *SubFunction) LastChangeTimeHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.LastChangeTime {
	data := psdc.ConvertToLastChangeTimeHeader(getStringPtr(getSystemTime()))

	return data
}

func getSystemDate() string {
	day := time.Now()
	return day.Format("2006-01-02")
}

func getSystemTime() string {
	day := time.Now()
	return day.Format("15:04:05")
}

func getStringPtr(s string) *string {
	return &s
}
