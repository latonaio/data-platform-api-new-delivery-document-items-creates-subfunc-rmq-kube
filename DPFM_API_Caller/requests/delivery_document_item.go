package requests

type DeliveryDocumentItem struct {
	OrderID                    int `json:"OrderID"`
	OrderItem                  int `json:"OrderItem"`
	DeliveryDocumentItemNumber int `json:"DeliveryDocumentItemNumber"`
}
