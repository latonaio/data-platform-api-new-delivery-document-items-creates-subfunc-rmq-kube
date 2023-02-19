package subfunction

import (
	api_input_reader "data-platform-api-delivery-document-items-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-delivery-document-items-creates-subfunc/API_Processing_Data_Formatter"
	"strings"
)

func (f *SubFunction) DeliveryDocumentItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.DeliveryDocumentItem {
	data := psdc.ConvertToDeliveryDocumentItem(sdc)

	return data
}

func (f *SubFunction) OrdersItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrdersItem, error) {
	args := make([]interface{}, 0)

	orderItem := psdc.OrderItem

	repeat := strings.Repeat("(?,?),", len(orderItem)-1) + "(?,?)"
	for _, v := range orderItem {
		args = append(args, v.OrderID, v.OrderItem)
	}

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, OrderItemCategory, SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID,
		SupplyChainRelationshipDeliveryPlantID, SupplyChainRelationshipStockConfPlantID, SupplyChainRelationshipProductionPlantID,
		OrderItemText, OrderItemTextByBuyer, OrderItemTextBySeller, Product, ProductStandardID, ProductGroup, BaseUnit,
		DeliverToParty, DeliverFromParty, DeliverToPlant, DeliverToPlantTimeZone, DeliverToPlantStorageLocation,
		ProductIsBatchManagedInDeliverToPlant, BatchMgmtPolicyInDeliverToPlant, DeliverToPlantBatch, DeliverToPlantBatchValidityStartDate,
		DeliverToPlantBatchValidityEndDate, DeliverFromPlant, DeliverFromPlantTimeZone, DeliverFromPlantStorageLocation,
		ProductIsBatchManagedInDeliverFromPlant, BatchMgmtPolicyInDeliverFromPlant, DeliverFromPlantBatch,
		DeliverFromPlantBatchValidityStartDate, DeliverFromPlantBatchValidityEndDate, DeliveryUnit, StockConfirmationBusinessPartner,
		StockConfirmationPlant, StockConfirmationPlantTimeZone, ProductIsBatchManagedInStockConfirmationPlant,
		BatchMgmtPolicyInStockConfirmationPlant, StockConfirmationPlantBatch, StockConfirmationPlantBatchValidityStartDate,
		StockConfirmationPlantBatchValidityEndDate, OrderQuantityInBaseUnit, OrderQuantityInDeliveryUnit, StockConfirmationPolicy,
		StockConfirmationStatus, ConfirmedOrderQuantityInBaseUnit, ItemWeightUnit, ProductGrossWeight, ItemGrossWeight, ProductNetWeight,
		ItemNetWeight, InternalCapacityQuantity, InternalCapacityQuantityUnit, NetAmount, TaxAmount, GrossAmount,
		ProductionPlantBusinessPartner, ProductionPlant, ProductionPlantTimeZone, ProductionPlantStorageLocation,
		ProductIsBatchManagedInProductionPlant, BatchMgmtPolicyInProductionPlant, ProductionPlantBatch,
		ProductionPlantBatchValidityStartDate, ProductionPlantBatchValidityEndDate, Incoterms, TransactionTaxClassification,
		ProductTaxClassificationBillToCountry, ProductTaxClassificationBillFromCountry, DefinedTaxClassification, AccountAssignmentGroup,
		ProductAccountAssignmentGroup, PaymentTerms, PaymentMethod, Project, TaxCode, TaxRate, CountryOfOrigin, CountryOfOriginLanguage
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE (OrderID, OrderItem) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrdersItem(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) CreationDateItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.CreationDate {
	data := psdc.ConvertToCreationDateItem(getStringPtr(getSystemDate()))

	return data
}

func (f *SubFunction) LastChangeDateItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.LastChangeDate {
	data := psdc.ConvertToLastChangeDateItem(getStringPtr(getSystemDate()))

	return data
}

func (f *SubFunction) CreationTimeItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.CreationTime {
	data := psdc.ConvertToCreationTimeItem(getStringPtr(getSystemTime()))

	return data
}

func (f *SubFunction) LastChangeTimeItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.LastChangeTime {
	data := psdc.ConvertToLastChangeTimeItem(getStringPtr(getSystemTime()))

	return data
}
