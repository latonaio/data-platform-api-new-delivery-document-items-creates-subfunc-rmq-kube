package requests

type OrderItemKey struct {
	OrderID                       *int      `json:"OrderID"`
	OrderItem                     *int      `json:"OrderItem"`
	DeliverToParty                []*int    `json:"DeliverToParty"`
	DeliverToPartyFrom            *int      `json:"DeliverToPartyFrom"`
	DeliverToPartyTo              *int      `json:"DeliverToPartyTo"`
	DeliverFromParty              []*int    `json:"DeliverFromParty"`
	DeliverFromPartyFrom          *int      `json:"DeliverFromPartyFrom"`
	DeliverFromPartyTo            *int      `json:"DeliverFromPartyTo"`
	DeliverToPlant                []*string `json:"DeliverToPlant"`
	DeliverToPlantFrom            *string   `json:"DeliverToPlantFrom"`
	DeliverToPlantTo              *string   `json:"DeliverToPlantTo"`
	DeliverFromPlant              []*string `json:"DeliverFromPlant"`
	DeliverFromPlantFrom          *string   `json:"DeliverFromPlantFrom"`
	DeliverFromPlantTo            *string   `json:"DeliverFromPlantTo"`
	ItemCompleteDeliveryIsDefined bool      `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus            string    `json:"ItemDeliveryStatus"`
	ItemBlockStatus               bool      `json:"ItemBlockStatus"`
	ItemDeliveryBlockStatus       bool      `json:"ItemDeliveryBlockStatus"`
}
