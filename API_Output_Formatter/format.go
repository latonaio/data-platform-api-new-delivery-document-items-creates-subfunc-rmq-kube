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
	res := make([]*Item, 0)

	ordersHeaderMap := StructArrayToMap(psdc.OrdersHeader, "OrderID")

	processType := psdc.ProcessType
	if processType.BulkProcess {
		for i := range psdc.DeliveryDocumentItem {
			item := &Item{}
			inputItem := sdc.Header.Item[0]

			orderID := psdc.OrdersItemScheduleLine[i].OrderID
			orderItem := psdc.OrdersItemScheduleLine[i].OrderItem
			idx := -1
			for j, v := range psdc.OrdersItem {
				if v.OrderID == orderID && v.OrderItem == orderItem {
					idx = j
					break
				}
			}
			if idx == -1 {
				continue
			}

			// 入力ファイル
			item, err = jsonTypeConversion(item, inputItem)
			if err != nil {
				return nil, err
			}

			// 1-2
			item, err = jsonTypeConversion(item, psdc.OrdersItem[i])
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

			item.DeliveryDocument = psdc.CalculateDeliveryDocument.DeliveryDocument
			item.DeliveryDocumentItem = psdc.DeliveryDocumentItem[i].DeliveryDocumentItemNumber
			item.DeliveryDocumentItemCategory = &psdc.OrdersItem[idx].OrderItemCategory

			item.DeliveryDocumentItemText = &psdc.OrdersItem[idx].OrderItemText
			item.DeliveryDocumentItemTextByBuyer = psdc.OrdersItem[idx].OrderItemTextByBuyer
			item.DeliveryDocumentItemTextBySeller = psdc.OrdersItem[idx].OrderItemTextBySeller

			item.OriginalQuantityInBaseUnit = &psdc.OrdersItem[idx].OrderQuantityInBaseUnit

			item.CreationDate = psdc.CreationDateItem.CreationDate
			item.CreationTime = psdc.CreationTimeItem.CreationTime
			item.LastChangeDate = psdc.LastChangeDateItem.LastChangeDate
			item.LastChangeTime = psdc.LastChangeTimeItem.LastChangeTime
			item.ItemBillingStatus = getStringPtr("NP")
			// item.SalesCostGLAccount =  //TBD
			// item.ReceivingGLAccount =  //TBD
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
			// item.ItemIsCancelled =  //仕様書になし
			// item.ItemIsDeleted =  //仕様書になし

			res = append(res, item)
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

			item.DeliveryDocument = psdc.CalculateDeliveryDocument.DeliveryDocument
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

			res = append(res, item)
		}
	}

	return res, nil
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
