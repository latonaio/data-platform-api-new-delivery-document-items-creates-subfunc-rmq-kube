package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-delivery-document-items-creates-subfunc/API_Input_Reader"
	"data-platform-api-delivery-document-items-creates-subfunc/DPFM_API_Caller/requests"
	"database/sql"

	"golang.org/x/xerrors"
)

// Initializer
func (psdc *SDC) ConvertToMetaData(sdc *api_input_reader.SDC) *MetaData {
	pm := &requests.MetaData{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
	}

	data := pm
	res := MetaData{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
	}

	return &res
}

func (psdc *SDC) ConvertToProcessType() *ProcessType {
	pm := &requests.ProcessType{}

	data := pm
	res := ProcessType{
		BulkProcess:       data.BulkProcess,
		IndividualProcess: data.IndividualProcess,
		ArraySpec:         data.ArraySpec,
		RangeSpec:         data.RangeSpec,
	}

	return &res
}

func (psdc *SDC) ConvertToOrderItemKey() *OrderItemKey {
	pm := &requests.OrderItemKey{
		ItemCompleteDeliveryIsDefined: false,
		ItemDeliveryStatus:            "CL",
		ItemBlockStatus:               false,
		ItemDeliveryBlockStatus:       false,
	}

	data := pm
	res := OrderItemKey{
		DeliverToParty:                data.DeliverToParty,
		DeliverToPartyFrom:            data.DeliverToPartyFrom,
		DeliverToPartyTo:              data.DeliverToPartyTo,
		DeliverFromParty:              data.DeliverFromParty,
		DeliverFromPartyFrom:          data.DeliverFromPartyFrom,
		DeliverFromPartyTo:            data.DeliverFromPartyTo,
		DeliverToPlant:                data.DeliverToPlant,
		DeliverToPlantFrom:            data.DeliverToPlantFrom,
		DeliverToPlantTo:              data.DeliverToPlantTo,
		DeliverFromPlant:              data.DeliverFromPlant,
		DeliverFromPlantFrom:          data.DeliverFromPlantFrom,
		DeliverFromPlantTo:            data.DeliverFromPlantTo,
		ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
		ItemDeliveryStatus:            data.ItemDeliveryStatus,
		ItemBlockStatus:               data.ItemBlockStatus,
		ItemDeliveryBlockStatus:       data.ItemDeliveryBlockStatus,
	}

	return &res
}

func (psdc *SDC) ConvertToOrderItem(rows *sql.Rows) ([]*OrderItem, error) {
	defer rows.Close()
	res := make([]*OrderItem, 0)

	i := 0
	for rows.Next() {
		pm := &requests.OrderItem{}
		i++

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.ItemBlockStatus,
			&pm.ItemDeliveryBlockStatus,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrderItem{
			OrderID:                       data.OrderID,
			OrderItem:                     data.OrderItem,
			DeliverToParty:                data.DeliverToParty,
			DeliverFromParty:              data.DeliverFromParty,
			DeliverToPlant:                data.DeliverToPlant,
			DeliverFromPlant:              data.DeliverFromPlant,
			ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:            data.ItemDeliveryStatus,
			ItemBlockStatus:               data.ItemBlockStatus,
			ItemDeliveryBlockStatus:       data.ItemDeliveryBlockStatus,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_orders_item_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrderItemInIndividualProcessKey() *OrderItemKey {
	pm := &requests.OrderItemKey{
		ItemCompleteDeliveryIsDefined: false,
		ItemDeliveryStatus:            "CL",
		ItemDeliveryBlockStatus:       false,
	}

	data := pm
	res := OrderItemKey{
		OrderID:                       data.OrderID,
		OrderItem:                     data.OrderItem,
		ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
		ItemDeliveryStatus:            data.ItemDeliveryStatus,
		ItemDeliveryBlockStatus:       data.ItemDeliveryBlockStatus,
	}

	return &res
}

func (psdc *SDC) ConvertToOrderItemInIndividualProcess(rows *sql.Rows) ([]*OrderItem, error) {
	defer rows.Close()
	res := make([]*OrderItem, 0)

	i := 0
	for rows.Next() {
		pm := &requests.OrderItem{}
		i++

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.ItemDeliveryBlockStatus,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrderItem{
			OrderID:                       data.OrderID,
			OrderItem:                     data.OrderItem,
			ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:            data.ItemDeliveryStatus,
			ItemDeliveryBlockStatus:       data.ItemDeliveryBlockStatus,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_orders_item_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrdersItemScheduleLineKey() *OrdersItemScheduleLineKey {
	pm := &requests.OrdersItemScheduleLineKey{
		ItemScheduleLineDeliveryBlockStatus: false,
		OpenConfirmedQuantityInBaseUnit:     0,
	}

	data := pm
	res := OrdersItemScheduleLineKey{
		OrderID:                             data.OrderID,
		OrderItem:                           data.OrderItem,
		ConfirmedDeliveryDate:               data.ConfirmedDeliveryDate,
		ConfirmedDeliveryDateFrom:           data.ConfirmedDeliveryDateFrom,
		ConfirmedDeliveryDateTo:             data.ConfirmedDeliveryDateTo,
		ItemScheduleLineDeliveryBlockStatus: data.ItemScheduleLineDeliveryBlockStatus,
		OpenConfirmedQuantityInBaseUnit:     data.OpenConfirmedQuantityInBaseUnit,
	}

	return &res
}

func (psdc *SDC) ConvertToOrdersItemScheduleLine(rows *sql.Rows) ([]*OrdersItemScheduleLine, error) {
	defer rows.Close()
	res := make([]*OrdersItemScheduleLine, 0)

	i := 0
	for rows.Next() {
		pm := &requests.OrdersItemScheduleLine{}
		i++

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ScheduleLine,
			&pm.RequestedDeliveryDate,
			&pm.ConfirmedDeliveryDate,
			&pm.OrderQuantityInBaseUnit,
			&pm.ConfirmedOrderQuantityByPDTAvailCheck,
			&pm.OpenConfirmedQuantityInBaseUnit,
			&pm.StockIsFullyConfirmed,
			&pm.ItemScheduleLineDeliveryBlockStatus,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrdersItemScheduleLine{
			OrderID:                               data.OrderID,
			OrderItem:                             data.OrderItem,
			ScheduleLine:                          data.ScheduleLine,
			RequestedDeliveryDate:                 data.RequestedDeliveryDate,
			ConfirmedDeliveryDate:                 data.ConfirmedDeliveryDate,
			OrderQuantityInBaseUnit:               data.OrderQuantityInBaseUnit,
			ConfirmedOrderQuantityByPDTAvailCheck: data.ConfirmedOrderQuantityByPDTAvailCheck,
			OpenConfirmedQuantityInBaseUnit:       data.OpenConfirmedQuantityInBaseUnit,
			StockIsFullyConfirmed:                 data.StockIsFullyConfirmed,
			ItemScheduleLineDeliveryBlockStatus:   data.ItemScheduleLineDeliveryBlockStatus,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_orders_item_schedule_line_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

// Header
func (psdc *SDC) ConvertToOrdersHeader(rows *sql.Rows) ([]*OrdersHeader, error) {
	defer rows.Close()
	res := make([]*OrdersHeader, 0)

	i := 0
	for rows.Next() {
		pm := &requests.OrdersHeader{}
		i++

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderType,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.SupplyChainRelationshipPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.Payer,
			&pm.Payee,
			&pm.ContractType,
			&pm.OrderValidityStartDate,
			&pm.OrderValidityEndDate,
			&pm.InvoicePeriodStartDate,
			&pm.InvoicePeriodEndDate,
			&pm.TransactionCurrency,
			&pm.Incoterms,
			&pm.IsExportImport,
			&pm.HeaderText,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrdersHeader{
			OrderID:                          data.OrderID,
			OrderType:                        data.OrderType,
			SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
			SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
			SupplyChainRelationshipPaymentID: data.SupplyChainRelationshipPaymentID,
			Buyer:                            data.Buyer,
			Seller:                           data.Seller,
			BillToParty:                      data.BillToParty,
			BillFromParty:                    data.BillFromParty,
			BillToCountry:                    data.BillToCountry,
			BillFromCountry:                  data.BillFromCountry,
			Payer:                            data.Payer,
			Payee:                            data.Payee,
			ContractType:                     data.ContractType,
			OrderValidityStartDate:           data.OrderValidityStartDate,
			OrderValidityEndDate:             data.OrderValidityEndDate,
			InvoicePeriodStartDate:           data.InvoicePeriodStartDate,
			InvoicePeriodEndDate:             data.InvoicePeriodEndDate,
			TransactionCurrency:              data.TransactionCurrency,
			Incoterms:                        data.Incoterms,
			IsExportImport:                   data.IsExportImport,
			HeaderText:                       data.HeaderText,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_orders_header_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToCalculateDeliveryDocumentKey() *CalculateDeliveryDocumentKey {
	pm := &requests.CalculateDeliveryDocumentKey{
		FieldNameWithNumberRange: "DeliveryDocument",
	}

	data := pm
	res := CalculateDeliveryDocumentKey{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
	}

	return &res
}

func (psdc *SDC) ConvertToCalculateDeliveryDocumentQueryGets(rows *sql.Rows) (*CalculateDeliveryDocumentQueryGets, error) {
	defer rows.Close()
	pm := &requests.CalculateDeliveryDocumentQueryGets{}

	i := 0
	for rows.Next() {
		i++
		err := rows.Scan(
			&pm.ServiceLabel,
			&pm.FieldNameWithNumberRange,
			&pm.DeliveryDocumentLatestNumber,
		)
		if err != nil {
			return nil, err
		}
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_number_range_latest_number_data'テーブルに対象のレコードが存在しません。")
	}

	data := pm
	res := CalculateDeliveryDocumentQueryGets{
		ServiceLabel:                 data.ServiceLabel,
		FieldNameWithNumberRange:     data.FieldNameWithNumberRange,
		DeliveryDocumentLatestNumber: data.DeliveryDocumentLatestNumber,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCalculateDeliveryDocument(deliveryDocumentLatestNumber *int, deliveryDocument int) *CalculateDeliveryDocument {
	pm := &requests.CalculateDeliveryDocument{}

	pm.DeliveryDocumentLatestNumber = deliveryDocumentLatestNumber
	pm.DeliveryDocument = deliveryDocument

	data := pm
	res := CalculateDeliveryDocument{
		DeliveryDocumentLatestNumber: data.DeliveryDocumentLatestNumber,
		DeliveryDocument:             data.DeliveryDocument,
	}

	return &res
}

func (psdc *SDC) ConvertToDocumentDate(documentDate *string) *DocumentDate {
	pm := &requests.DocumentDate{}

	pm.DocumentDate = documentDate

	data := pm
	res := DocumentDate{
		DocumentDate: data.DocumentDate,
	}

	return &res
}

func (psdc *SDC) ConvertToPaymentTerms(rows *sql.Rows) ([]*PaymentTerms, error) {
	defer rows.Close()
	res := make([]*PaymentTerms, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PaymentTerms{}

		err := rows.Scan(
			&pm.PaymentTerms,
			&pm.BaseDate,
			&pm.BaseDateCalcAddMonth,
			&pm.BaseDateCalcFixedDate,
			&pm.PaymentDueDateCalcAddMonth,
			&pm.PaymentDueDateCalcFixedDate,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &PaymentTerms{
			PaymentTerms:                data.PaymentTerms,
			BaseDate:                    data.BaseDate,
			BaseDateCalcAddMonth:        data.BaseDateCalcAddMonth,
			BaseDateCalcFixedDate:       data.BaseDateCalcFixedDate,
			PaymentDueDateCalcAddMonth:  data.PaymentDueDateCalcAddMonth,
			PaymentDueDateCalcFixedDate: data.PaymentDueDateCalcFixedDate,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_payment_terms_payment_terms_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToInvoiceDocumentDate(sdc *api_input_reader.SDC) *InvoiceDocumentDate {
	pm := &requests.InvoiceDocumentDate{}

	pm.InvoiceDocumentDate = sdc.Header.InvoiceDocumentDate
	data := pm

	res := InvoiceDocumentDate{
		PlannedGoodsIssueDate: data.PlannedGoodsIssueDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &res
}

func (psdc *SDC) ConvertToPlannedGoodsIssueDate(sdc *api_input_reader.SDC) (*InvoiceDocumentDate, error) {
	if sdc.Header.PlannedGoodsIssueDate == nil {
		return nil, xerrors.Errorf("PlannedGoodsIssueDateがnullです。")
	}

	pm := &requests.InvoiceDocumentDate{
		PlannedGoodsIssueDate: *sdc.Header.PlannedGoodsIssueDate,
	}

	data := pm
	res := InvoiceDocumentDate{
		PlannedGoodsIssueDate: data.PlannedGoodsIssueDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCaluculateInvoiceDocumentDate(sdc *api_input_reader.SDC, invoiceDocumentDate *string) *InvoiceDocumentDate {
	pm := &requests.InvoiceDocumentDate{
		PlannedGoodsIssueDate: *sdc.Header.PlannedGoodsIssueDate,
	}

	pm.InvoiceDocumentDate = invoiceDocumentDate

	data := pm
	res := InvoiceDocumentDate{
		PlannedGoodsIssueDate: data.PlannedGoodsIssueDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &res
}

func (psdc *SDC) ConvertToHeaderGrossWeight(headerGrossWeight *float32) *HeaderGrossWeight {
	pm := &requests.HeaderGrossWeight{}

	pm.HeaderGrossWeight = headerGrossWeight

	data := pm
	res := HeaderGrossWeight{
		HeaderGrossWeight: data.HeaderGrossWeight,
	}

	return &res
}

func (psdc *SDC) ConvertToHeaderNetWeight(headerNetWeight *float32) *HeaderNetWeight {
	pm := &requests.HeaderNetWeight{}

	pm.HeaderNetWeight = headerNetWeight

	data := pm
	res := HeaderNetWeight{
		HeaderNetWeight: data.HeaderNetWeight,
	}

	return &res
}

func (psdc *SDC) ConvertToCreationDateHeader(systemDate *string) *CreationDate {
	pm := &requests.CreationDate{}

	pm.CreationDate = systemDate

	data := pm
	res := CreationDate{
		CreationDate: data.CreationDate,
	}

	return &res
}

func (psdc *SDC) ConvertToLastChangeDateHeader(systemDate *string) *LastChangeDate {
	pm := &requests.LastChangeDate{}

	pm.LastChangeDate = systemDate

	data := pm
	res := LastChangeDate{
		LastChangeDate: data.LastChangeDate,
	}

	return &res
}

func (psdc *SDC) ConvertToCreationTimeHeader(systemTime *string) *CreationTime {
	pm := &requests.CreationTime{}

	pm.CreationTime = systemTime

	data := pm
	res := CreationTime{
		CreationTime: data.CreationTime,
	}

	return &res
}

func (psdc *SDC) ConvertToLastChangeTimeHeader(systemTime *string) *LastChangeTime {
	pm := &requests.LastChangeTime{}

	pm.LastChangeTime = systemTime

	data := pm
	res := LastChangeTime{
		LastChangeTime: data.LastChangeTime,
	}

	return &res
}

// Item
func (psdc *SDC) ConvertToDeliveryDocumentItem(sdc *api_input_reader.SDC) []*DeliveryDocumentItem {
	res := make([]*DeliveryDocumentItem, 0)
	if psdc.ProcessType.BulkProcess {
		for i := range psdc.OrdersItemScheduleLine {
			pm := &requests.DeliveryDocumentItem{}

			pm.DeliveryDocumentItemNumber = i + 1

			data := pm
			res = append(res, &DeliveryDocumentItem{
				DeliveryDocumentItemNumber: data.DeliveryDocumentItemNumber,
			})
		}
	} else if psdc.ProcessType.IndividualProcess {
		for i := range sdc.Header.Item {
			pm := &requests.DeliveryDocumentItem{}

			pm.DeliveryDocumentItemNumber = i + 1

			data := pm
			res = append(res, &DeliveryDocumentItem{
				DeliveryDocumentItemNumber: data.DeliveryDocumentItemNumber,
			})
		}
	}

	return res
}

func (psdc *SDC) ConvertToOrdersItem(rows *sql.Rows) ([]*OrdersItem, error) {
	defer rows.Close()
	res := make([]*OrdersItem, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrdersItem{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.OrderItemCategory,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.OrderItemText,
			&pm.OrderItemTextByBuyer,
			&pm.OrderItemTextBySeller,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.BaseUnit,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverToPlantTimeZone,
			&pm.DeliverToPlantStorageLocation,
			&pm.ProductIsBatchManagedInDeliverToPlant,
			&pm.BatchMgmtPolicyInDeliverToPlant,
			&pm.DeliverToPlantBatch,
			&pm.DeliverToPlantBatchValidityStartDate,
			&pm.DeliverToPlantBatchValidityEndDate,
			&pm.DeliverFromPlant,
			&pm.DeliverFromPlantTimeZone,
			&pm.DeliverFromPlantStorageLocation,
			&pm.ProductIsBatchManagedInDeliverFromPlant,
			&pm.BatchMgmtPolicyInDeliverFromPlant,
			&pm.DeliverFromPlantBatch,
			&pm.DeliverFromPlantBatchValidityStartDate,
			&pm.DeliverFromPlantBatchValidityEndDate,
			&pm.DeliveryUnit,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantTimeZone,
			&pm.ProductIsBatchManagedInStockConfirmationPlant,
			&pm.BatchMgmtPolicyInStockConfirmationPlant,
			&pm.StockConfirmationPlantBatch,
			&pm.StockConfirmationPlantBatchValidityStartDate,
			&pm.StockConfirmationPlantBatchValidityEndDate,
			&pm.OrderQuantityInBaseUnit,
			&pm.OrderQuantityInDeliveryUnit,
			&pm.StockConfirmationPolicy,
			&pm.StockConfirmationStatus,
			&pm.ConfirmedOrderQuantityInBaseUnit,
			&pm.ItemWeightUnit,
			&pm.ProductGrossWeight,
			&pm.ItemGrossWeight,
			&pm.ProductNetWeight,
			&pm.ItemNetWeight,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantTimeZone,
			&pm.ProductionPlantStorageLocation,
			&pm.ProductIsBatchManagedInProductionPlant,
			&pm.BatchMgmtPolicyInProductionPlant,
			&pm.ProductionPlantBatch,
			&pm.ProductionPlantBatchValidityStartDate,
			&pm.ProductionPlantBatchValidityEndDate,
			&pm.Incoterms,
			&pm.TransactionTaxClassification,
			&pm.ProductTaxClassificationBillToCountry,
			&pm.ProductTaxClassificationBillFromCountry,
			&pm.DefinedTaxClassification,
			&pm.AccountAssignmentGroup,
			&pm.ProductAccountAssignmentGroup,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.Project,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrdersItem{
			OrderID:                                       data.OrderID,
			OrderItem:                                     data.OrderItem,
			OrderItemCategory:                             data.OrderItemCategory,
			SupplyChainRelationshipID:                     data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:             data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:        data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipStockConfPlantID:       data.SupplyChainRelationshipStockConfPlantID,
			SupplyChainRelationshipProductionPlantID:      data.SupplyChainRelationshipProductionPlantID,
			OrderItemText:                                 data.OrderItemText,
			OrderItemTextByBuyer:                          data.OrderItemTextByBuyer,
			OrderItemTextBySeller:                         data.OrderItemTextBySeller,
			Product:                                       data.Product,
			ProductStandardID:                             data.ProductStandardID,
			ProductGroup:                                  data.ProductGroup,
			BaseUnit:                                      data.BaseUnit,
			DeliverToParty:                                data.DeliverToParty,
			DeliverFromParty:                              data.DeliverFromParty,
			DeliverToPlant:                                data.DeliverToPlant,
			DeliverToPlantTimeZone:                        data.DeliverToPlantTimeZone,
			DeliverToPlantStorageLocation:                 data.DeliverToPlantStorageLocation,
			ProductIsBatchManagedInDeliverToPlant:         data.ProductIsBatchManagedInDeliverToPlant,
			BatchMgmtPolicyInDeliverToPlant:               data.BatchMgmtPolicyInDeliverToPlant,
			DeliverToPlantBatch:                           data.DeliverToPlantBatch,
			DeliverToPlantBatchValidityStartDate:          data.DeliverToPlantBatchValidityStartDate,
			DeliverToPlantBatchValidityEndDate:            data.DeliverToPlantBatchValidityEndDate,
			DeliverFromPlant:                              data.DeliverFromPlant,
			DeliverFromPlantTimeZone:                      data.DeliverFromPlantTimeZone,
			DeliverFromPlantStorageLocation:               data.DeliverFromPlantStorageLocation,
			ProductIsBatchManagedInDeliverFromPlant:       data.ProductIsBatchManagedInDeliverFromPlant,
			BatchMgmtPolicyInDeliverFromPlant:             data.BatchMgmtPolicyInDeliverFromPlant,
			DeliverFromPlantBatch:                         data.DeliverFromPlantBatch,
			DeliverFromPlantBatchValidityStartDate:        data.DeliverFromPlantBatchValidityStartDate,
			DeliverFromPlantBatchValidityEndDate:          data.DeliverFromPlantBatchValidityEndDate,
			DeliveryUnit:                                  data.DeliveryUnit,
			StockConfirmationBusinessPartner:              data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                        data.StockConfirmationPlant,
			StockConfirmationPlantTimeZone:                data.StockConfirmationPlantTimeZone,
			ProductIsBatchManagedInStockConfirmationPlant: data.ProductIsBatchManagedInStockConfirmationPlant,
			BatchMgmtPolicyInStockConfirmationPlant:       data.BatchMgmtPolicyInStockConfirmationPlant,
			StockConfirmationPlantBatch:                   data.StockConfirmationPlantBatch,
			StockConfirmationPlantBatchValidityStartDate:  data.StockConfirmationPlantBatchValidityStartDate,
			StockConfirmationPlantBatchValidityEndDate:    data.StockConfirmationPlantBatchValidityEndDate,
			OrderQuantityInBaseUnit:                       data.OrderQuantityInBaseUnit,
			OrderQuantityInDeliveryUnit:                   data.OrderQuantityInDeliveryUnit,
			StockConfirmationPolicy:                       data.StockConfirmationPolicy,
			StockConfirmationStatus:                       data.StockConfirmationStatus,
			ConfirmedOrderQuantityInBaseUnit:              data.ConfirmedOrderQuantityInBaseUnit,
			ItemWeightUnit:                                data.ItemWeightUnit,
			ProductGrossWeight:                            data.ProductGrossWeight,
			ItemGrossWeight:                               data.ItemGrossWeight,
			ProductNetWeight:                              data.ProductNetWeight,
			ItemNetWeight:                                 data.ItemNetWeight,
			NetAmount:                                     data.NetAmount,
			TaxAmount:                                     data.TaxAmount,
			GrossAmount:                                   data.GrossAmount,
			ProductionPlantBusinessPartner:                data.ProductionPlantBusinessPartner,
			ProductionPlant:                               data.ProductionPlant,
			ProductionPlantTimeZone:                       data.ProductionPlantTimeZone,
			ProductionPlantStorageLocation:                data.ProductionPlantStorageLocation,
			ProductIsBatchManagedInProductionPlant:        data.ProductIsBatchManagedInProductionPlant,
			BatchMgmtPolicyInProductionPlant:              data.BatchMgmtPolicyInProductionPlant,
			ProductionPlantBatch:                          data.ProductionPlantBatch,
			ProductionPlantBatchValidityStartDate:         data.ProductionPlantBatchValidityStartDate,
			ProductionPlantBatchValidityEndDate:           data.ProductionPlantBatchValidityEndDate,
			Incoterms:                                     data.Incoterms,
			TransactionTaxClassification:                  data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:         data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry:       data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:                      data.DefinedTaxClassification,
			AccountAssignmentGroup:                        data.AccountAssignmentGroup,
			ProductAccountAssignmentGroup:                 data.ProductAccountAssignmentGroup,
			PaymentTerms:                                  data.PaymentTerms,
			PaymentMethod:                                 data.PaymentMethod,
			Project:                                       data.Project,
			TaxCode:                                       data.TaxCode,
			TaxRate:                                       data.TaxRate,
			CountryOfOrigin:                               data.CountryOfOrigin,
			CountryOfOriginLanguage:                       data.CountryOfOriginLanguage,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_orders_item_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToCreationDateItem(systemDate *string) *CreationDate {
	pm := &requests.CreationDate{}

	pm.CreationDate = systemDate

	data := pm
	res := CreationDate{
		CreationDate: data.CreationDate,
	}

	return &res
}

func (psdc *SDC) ConvertToLastChangeDateItem(systemDate *string) *LastChangeDate {
	pm := &requests.LastChangeDate{}

	pm.LastChangeDate = systemDate

	data := pm
	res := LastChangeDate{
		LastChangeDate: data.LastChangeDate,
	}

	return &res
}

func (psdc *SDC) ConvertToCreationTimeItem(systemTime *string) *CreationTime {
	pm := &requests.CreationTime{}

	pm.CreationTime = systemTime

	data := pm
	res := CreationTime{
		CreationTime: data.CreationTime,
	}

	return &res
}

func (psdc *SDC) ConvertToLastChangeTimeItem(systemTime *string) *LastChangeTime {
	pm := &requests.LastChangeTime{}

	pm.LastChangeTime = systemTime

	data := pm
	res := LastChangeTime{
		LastChangeTime: data.LastChangeTime,
	}

	return &res
}
