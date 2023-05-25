package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-delivery-document-items-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-delivery-document-items-creates-subfunc/API_Processing_Data_Formatter"
	"encoding/json"
	"reflect"

	"golang.org/x/xerrors"
)

func ConvertToItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*Item, error) {
	var err error
	items := make([]*Item, 0)

	ordersHeaderMap := StructArrayToMap(psdc.OrdersHeader, "OrderID")

	processType := psdc.ProcessType
	if processType.BulkProcess {
		for i := range psdc.DeliveryDocumentItem {
			item := &Item{}
			inputItem := sdc.Header.Item[0]

			orderID := psdc.OrdersItemScheduleLine[i].OrderID
			orderItem := psdc.OrdersItemScheduleLine[i].OrderItem
			var deliverFromPlant, deliverToPlant string

			ordersItemIdx := -1
			for j, ordersItem := range psdc.OrdersItem {
				if ordersItem.OrderID == orderID && ordersItem.OrderItem == orderItem {
					if ordersItem.DeliverFromPlant == nil || ordersItem.DeliverToPlant == nil {
						continue
					}
					deliverFromPlant = *ordersItem.DeliverFromPlant
					deliverToPlant = *ordersItem.DeliverToPlant

					ordersItemIdx = j
					break
				}
			}
			if ordersItemIdx == -1 {
				continue
			}

			// 入力ファイル
			item, err = jsonTypeConversion(item, inputItem)
			if err != nil {
				return nil, err
			}

			// 1-2
			item, err = jsonTypeConversion(item, psdc.OrdersItem[ordersItemIdx])
			if err != nil {
				return nil, xerrors.Errorf("request create error: %w", err)
			}

			// 1-1
			if _, ok := ordersHeaderMap[orderID]; !ok {
				continue
			}
			item, err = jsonTypeConversion(item, ordersHeaderMap[orderID])
			if err != nil {
				return nil, xerrors.Errorf("request create error: %w", err)
			}

			deliveryDocumentIdx := -1
			for j, deliveryDocument := range psdc.CalculateDeliveryDocument {
				if deliveryDocument.DeliverFromPlant == deliverFromPlant && deliveryDocument.DeliverToPlant == deliverToPlant {
					deliveryDocumentIdx = j
					break
				}
			}
			if deliveryDocumentIdx == -1 {
				continue
			}
			item.DeliveryDocument = psdc.CalculateDeliveryDocument[deliveryDocumentIdx].DeliveryDocument
			item.DeliveryDocumentItem = psdc.DeliveryDocumentItem[i].DeliveryDocumentItemNumber
			item.DeliveryDocumentItemCategory = &psdc.OrdersItem[ordersItemIdx].OrderItemCategory

			item.DeliveryDocumentItemText = &psdc.OrdersItem[ordersItemIdx].OrderItemText
			item.DeliveryDocumentItemTextByBuyer = psdc.OrdersItem[ordersItemIdx].OrderItemTextByBuyer
			item.DeliveryDocumentItemTextBySeller = psdc.OrdersItem[ordersItemIdx].OrderItemTextBySeller

			item.OriginalQuantityInBaseUnit = &psdc.OrdersItem[ordersItemIdx].OrderQuantityInBaseUnit

			item.CreationDate = psdc.CreationDateItem.CreationDate
			item.CreationTime = psdc.CreationTimeItem.CreationTime
			item.LastChangeDate = psdc.LastChangeDateItem.LastChangeDate
			item.LastChangeTime = psdc.LastChangeTimeItem.LastChangeTime
			item.ItemBillingStatus = getStringPtr("NP")
			// item.SalesCostGLAccount =  //TBD
			// item.ReceivingGLAccount =  //TBD
			// item.IssuingGoodsMovementType = //TBD
			// item.ReceivingGoodsMovementType = //TBD
			item.ItemCompleteDeliveryIsDefined = getBoolPtr(false)
			// item.ItemIsBillingRelevant =  //TBD
			// item.DueCalculationBaseDate =  //仕様書になし
			// item.PaymentDueDate =  //仕様書になし
			// item.NetPaymentDays =  //仕様書になし

			item.ConfirmedDeliveryDate = &psdc.OrdersItemScheduleLine[i].ConfirmedDeliveryDate

			item.ItemDeliveryBlockStatus = getBoolPtr(false)
			item.ItemIssuingBlockStatus = getBoolPtr(false)
			item.ItemReceivingBlockStatus = getBoolPtr(false)
			item.ItemBillingBlockStatus = getBoolPtr(false)
			item.IsCancelled = getBoolPtr(false)
			item.IsMarkedForDeletion = getBoolPtr(false)

			items = append(items, item)
		}
	} else if processType.IndividualProcess {
		for i := range sdc.Header.Item {
			item := &Item{}
			inputItem := sdc.Header.Item[i]

			// 入力ファイル
			item, err = jsonTypeConversion(item, inputItem)
			if err != nil {
				return nil, err
			}

			// 1-1
			item, err = jsonTypeConversion(item, psdc.OrdersHeader[0])
			if err != nil {
				return nil, xerrors.Errorf("request create error: %w", err)
			}

			// 1-2
			item, err = jsonTypeConversion(item, psdc.OrdersItem[0])
			if err != nil {
				return nil, xerrors.Errorf("request create error: %w", err)
			}

			item.DeliveryDocument = psdc.CalculateDeliveryDocument[0].DeliveryDocument
			item.DeliveryDocumentItem = psdc.DeliveryDocumentItem[i].DeliveryDocumentItemNumber
			item.DeliveryDocumentItemCategory = &psdc.OrdersItem[0].OrderItemCategory

			item.DeliveryDocumentItemText = &psdc.OrdersItem[0].OrderItemText
			item.DeliveryDocumentItemTextByBuyer = psdc.OrdersItem[0].OrderItemTextByBuyer
			item.DeliveryDocumentItemTextBySeller = psdc.OrdersItem[0].OrderItemTextBySeller

			item.OriginalQuantityInBaseUnit = &psdc.OrdersItem[0].OrderQuantityInBaseUnit

			item.CreationDate = psdc.CreationDateItem.CreationDate
			item.CreationTime = psdc.CreationTimeItem.CreationTime
			item.LastChangeDate = psdc.LastChangeDateItem.LastChangeDate
			item.LastChangeTime = psdc.LastChangeTimeItem.LastChangeTime
			item.ItemBillingStatus = getStringPtr("NP")
			// item.SalesCostGLAccount =  //TBD
			// item.ReceivingGLAccount =  //TBD
			// item.IssuingGoodsMovementType = //TBD
			// item.ReceivingGoodsMovementType = //TBD
			item.ItemCompleteDeliveryIsDefined = getBoolPtr(false)
			// item.ItemIsBillingRelevant =  //TBD

			// item.DueCalculationBaseDate =  //仕様書になし
			// item.PaymentDueDate =  //仕様書になし
			// item.NetPaymentDays =  //仕様書になし

			item.ConfirmedDeliveryDate = &psdc.OrdersItemScheduleLine[0].ConfirmedDeliveryDate

			item.ItemDeliveryBlockStatus = getBoolPtr(false)
			item.ItemIssuingBlockStatus = getBoolPtr(false)
			item.ItemReceivingBlockStatus = getBoolPtr(false)
			item.ItemBillingBlockStatus = getBoolPtr(false)
			// item.ItemIsCancelled =  //仕様書になし
			// item.ItemIsDeleted =  //仕様書になし

			items = append(items, item)
		}
	}

	return items, nil
}

func ConvertToPartner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*Partner, error) {
	var err error

	partners := make([]*Partner, 0)
	for _, deliveryDocument := range psdc.CalculateDeliveryDocument {
		for _, deliveryDocumentPartner := range psdc.Partner {
			partner := &Partner{}
			inputPartner := sdc.Header.Partner[0]
			partnerFunction := deliveryDocumentPartner.PartnerFunction
			businessPartner := deliveryDocumentPartner.BusinessPartner

			// 入力ファイル
			partner, err = jsonTypeConversion(partner, inputPartner)
			if err != nil {
				return nil, err
			}

			if partnerContain(partners, partnerFunction, businessPartner) {
				continue
			}

			partner.DeliveryDocument = deliveryDocument.DeliveryDocument
			partner.PartnerFunction = deliveryDocumentPartner.PartnerFunction
			partner.BusinessPartner = deliveryDocumentPartner.BusinessPartner
			partner.BusinessPartnerFullName = deliveryDocumentPartner.BusinessPartnerFullName
			partner.BusinessPartnerName = deliveryDocumentPartner.BusinessPartnerName
			partner.Country = deliveryDocumentPartner.Country
			partner.Language = deliveryDocumentPartner.Language
			partner.Currency = deliveryDocumentPartner.Currency
			partner.ExternalDocumentID = deliveryDocumentPartner.ExternalDocumentID
			partner.AddressID = deliveryDocumentPartner.AddressID

			partners = append(partners, partner)
		}
	}

	return partners, nil
}

func ConvertToAddress(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*Address, error) {
	var err error

	addresses := make([]*Address, 0)
	for _, deliveryDocument := range psdc.CalculateDeliveryDocument {
		for _, deliveryDocumentAddress := range psdc.Address {
			address := &Address{}
			inputAddress := sdc.Header.Address[0]
			addressID := deliveryDocumentAddress.AddressID

			// 入力ファイル
			address, err = jsonTypeConversion(address, inputAddress)
			if err != nil {
				return nil, err
			}

			if addressContain(addresses, addressID) {
				continue
			}

			address.DeliveryDocument = deliveryDocument.DeliveryDocument
			address.AddressID = deliveryDocumentAddress.AddressID
			address.PostalCode = deliveryDocumentAddress.PostalCode
			address.LocalRegion = deliveryDocumentAddress.LocalRegion
			address.Country = deliveryDocumentAddress.Country
			address.District = deliveryDocumentAddress.District
			address.StreetName = deliveryDocumentAddress.StreetName
			address.CityName = deliveryDocumentAddress.CityName
			address.Building = deliveryDocumentAddress.Building
			address.Floor = deliveryDocumentAddress.Floor
			address.Room = deliveryDocumentAddress.Room

			addresses = append(addresses, address)
		}
	}

	return addresses, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}

func getStringPtr(s string) *string {
	return &s
}

func StructArrayToMap[T any](data []T, key string) map[any]T {
	res := make(map[any]T, len(data))

	for _, value := range data {
		m := StructToMap[T](&value, key)
		for k, v := range m {
			res[k] = v
		}
	}

	return res
}

func StructToMap[T any](data interface{}, key string) map[any]T {
	res := make(map[any]T)
	elem := reflect.Indirect(reflect.ValueOf(data).Elem())
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		if field == key {
			rv := reflect.ValueOf(elem.Field(i).Interface())
			if rv.Kind() == reflect.Ptr {
				if rv.IsNil() {
					return nil
				}
			}
			value := reflect.Indirect(elem.Field(i)).Interface()
			var dist T
			res[value], _ = jsonTypeConversion(dist, elem.Interface())
			break
		}
	}

	return res
}

func jsonTypeConversion[T any](dist T, data interface{}) (T, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}

func partnerContain(partners []*Partner, partnerFunction string, businessPartner int) bool {
	for _, partner := range partners {
		if partnerFunction == partner.PartnerFunction && businessPartner == partner.BusinessPartner {
			return true
		}
	}
	return false
}

func addressContain(addresses []*Address, addressID int) bool {
	for _, address := range addresses {
		if addressID == address.AddressID {
			return true
		}
	}
	return false
}
