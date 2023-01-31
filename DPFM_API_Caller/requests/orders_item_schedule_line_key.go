package requests

type OrdersItemScheduleLineKey struct {
	OrderID                             []int     `json:"OrderID"`
	OrderItem                           []int     `json:"OrderItem"`
	ConfirmedDeliveryDate               []*string `json:"ConfirmedDeliveryDate"`
	ConfirmedDeliveryDateFrom           *string   `json:"ConfirmedDeliveryDateFrom"`
	ConfirmedDeliveryDateTo             *string   `json:"ConfirmedDeliveryDateTo"`
	ItemScheduleLineDeliveryBlockStatus bool      `json:"ItemScheduleLineDeliveryBlockStatus"`
	OpenConfirmedQuantityInBaseUnit     float32   `json:"OpenConfirmedQuantityInBaseUnit"`
}
