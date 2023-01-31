package requests

type OrderItem struct {
	OrderID                       int     `json:"OrderID"`
	OrderItem                     int     `json:"OrderItem"`
	DeliverToParty                *int    `json:"DeliverToParty"`
	DeliverFromParty              *int    `json:"DeliverFromParty"`
	DeliverToPlant                *string `json:"DeliverToPlant"`
	DeliverFromPlant              *string `json:"DeliverFromPlant"`
	ItemCompleteDeliveryIsDefined *bool   `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus            *string `json:"ItemDeliveryStatus"`
	ItemBlockStatus               *bool   `json:"ItemBlockStatus"`
	ItemDeliveryBlockStatus       *bool   `json:"ItemDeliveryBlockStatus"`
}
