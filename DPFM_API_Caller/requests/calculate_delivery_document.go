package requests

type CalculateDeliveryDocument struct {
	DeliveryDocument int    `json:"DeliveryDocument"`
	OrderID          int    `json:"OrderID"`
	OrderItem        int    `json:"OrderItem"`
	DeliverFromPlant string `json:"DeliverFromPlant"`
	DeliverToPlant   string `json:"DeliverToPlant"`
}
