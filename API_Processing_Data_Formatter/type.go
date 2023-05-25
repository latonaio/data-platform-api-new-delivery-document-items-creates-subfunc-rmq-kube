package api_processing_data_formatter

type SDC struct {
	MetaData                  *MetaData                    `json:"MetaData"`
	ProcessType               *ProcessType                 `json:"ProcessType"`
	OrderItem                 []*OrderItem                 `json:"OrderItem"`
	OrdersItemScheduleLine    []*OrdersItemScheduleLine    `json:"OrdersItemScheduleLine"`
	OrdersHeader              []*OrdersHeader              `json:"OrdersHeader"`
	CalculateDeliveryDocument []*CalculateDeliveryDocument `json:"CalculateDeliveryDocument"`
	DocumentDate              *DocumentDate                `json:"DocumentDate"`
	PaymentTerms              []*PaymentTerms              `json:"PaymentTerms"`
	InvoiceDocumentDate       *InvoiceDocumentDate         `json:"InvoiceDocumentDate"`
	HeaderGrossWeight         *HeaderGrossWeight           `json:"HeaderGrossWeight"`
	HeaderNetWeight           *HeaderNetWeight             `json:"HeaderNetWeight"`
	CreationDateHeader        *CreationDate                `json:"CreationDateHeader"`
	LastChangeDateHeader      *LastChangeDate              `json:"LastChangeDateHeader"`
	CreationTimeHeader        *CreationTime                `json:"CreationTimeHeader"`
	LastChangeTimeHeader      *LastChangeTime              `json:"LastChangeTimeHeader"`
	DeliveryDocumentItem      []*DeliveryDocumentItem      `json:"DeliveryDocumentItem"`
	OrdersItem                []*OrdersItem                `json:"OrdersItem"`
	Partner                   []*Partner                   `json:"Partner"`
	Address                   []*Address                   `json:"Address"`
	AddressMaster             []*AddressMaster             `json:"AddressMaster"`
	ItemIsBillingRelevant     *ItemIsBillingRelevant       `json:"ItemIsBillingRelevant"`
	CreationDateItem          *CreationDate                `json:"CreationDateItem"`
	LastChangeDateItem        *LastChangeDate              `json:"LastChangeDateItem"`
	CreationTimeItem          *CreationTime                `json:"CreationTimeItem"`
	LastChangeTimeItem        *LastChangeTime              `json:"LastChangeTimeItem"`
}

// Initializer
type MetaData struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
}

type ProcessType struct {
	BulkProcess       bool `json:"BulkProcess"`
	IndividualProcess bool `json:"IndividualProcess"`
	ArraySpec         bool `json:"ArraySpec"`
	RangeSpec         bool `json:"RangeSpec"`
}

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

type OrdersItemScheduleLineKey struct {
	OrderID                             []int     `json:"OrderID"`
	OrderItem                           []int     `json:"OrderItem"`
	ConfirmedDeliveryDate               []*string `json:"ConfirmedDeliveryDate"`
	ConfirmedDeliveryDateFrom           *string   `json:"ConfirmedDeliveryDateFrom"`
	ConfirmedDeliveryDateTo             *string   `json:"ConfirmedDeliveryDateTo"`
	ItemScheduleLineDeliveryBlockStatus bool      `json:"ItemScheduleLineDeliveryBlockStatus"`
	OpenConfirmedQuantityInBaseUnit     float32   `json:"OpenConfirmedQuantityInBaseUnit"`
}

type OrdersItemScheduleLine struct {
	OrderID                               int      `json:"OrderID"`
	OrderItem                             int      `json:"OrderItem"`
	ScheduleLine                          int      `json:"ScheduleLine"`
	RequestedDeliveryDate                 string   `json:"RequestedDeliveryDate"`
	ConfirmedDeliveryDate                 string   `json:"ConfirmedDeliveryDate"`
	OrderQuantityInBaseUnit               float32  `json:"OrderQuantityInBaseUnit"`
	ConfirmedOrderQuantityByPDTAvailCheck float32  `json:"ConfirmedOrderQuantityByPDTAvailCheck"`
	OpenConfirmedQuantityInBaseUnit       *float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	StockIsFullyConfirmed                 *bool    `json:"StockIsFullyConfirmed"`
	ItemScheduleLineDeliveryBlockStatus   *bool    `json:"ItemScheduleLineDeliveryBlockStatus"`
}

// Header
type OrdersHeader struct {
	OrderID                          int     `json:"OrderID"`
	OrderType                        string  `json:"OrderType"`
	SupplyChainRelationshipID        int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID *int    `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID *int    `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int     `json:"Buyer"`
	Seller                           int     `json:"Seller"`
	BillToParty                      *int    `json:"BillToParty"`
	BillFromParty                    *int    `json:"BillFromParty"`
	BillToCountry                    *string `json:"BillToCountry"`
	BillFromCountry                  *string `json:"BillFromCountry"`
	Payer                            *int    `json:"Payer"`
	Payee                            *int    `json:"Payee"`
	ContractType                     *string `json:"ContractType"`
	OrderValidityStartDate           *string `json:"OrderValidityStartDate"`
	OrderValidityEndDate             *string `json:"OrderValidityEndDate"`
	InvoicePeriodStartDate           *string `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate             *string `json:"InvoicePeriodEndDate"`
	TransactionCurrency              string  `json:"TransactionCurrency"`
	Incoterms                        *string `json:"Incoterms"`
	IsExportImport                   *bool   `json:"IsExportImport"`
	HeaderText                       *string `json:"HeaderText"`
}

type CalculateDeliveryDocument struct {
	DeliveryDocument int    `json:"DeliveryDocument"`
	OrderID          int    `json:"OrderID"`
	OrderItem        int    `json:"OrderItem"`
	DeliverFromPlant string `json:"DeliverFromPlant"`
	DeliverToPlant   string `json:"DeliverToPlant"`
}

type DeliverPlant struct {
	DeliverFromPlant string `json:"DeliverFromPlant"`
	DeliverToPlant   string `json:"DeliverToPlant"`
	OrderID          int    `json:"OrderID"`
	OrderItem        int    `json:"OrderItem"`
}

type DocumentDate struct {
	DocumentDate *string `json:"DocumentDate"`
}

type PaymentTerms struct {
	PaymentTerms                string `json:"PaymentTerms"`
	BaseDate                    int    `json:"BaseDate"`
	BaseDateCalcAddMonth        *int   `json:"BaseDateCalcAddMonth"`
	BaseDateCalcFixedDate       *int   `json:"BaseDateCalcFixedDate"`
	PaymentDueDateCalcAddMonth  *int   `json:"PaymentDueDateCalcAddMonth"`
	PaymentDueDateCalcFixedDate *int   `json:"PaymentDueDateCalcFixedDate"`
}

type InvoiceDocumentDate struct {
	PlannedGoodsIssueDate string  `json:"PlannedGoodsIssueDate"`
	InvoiceDocumentDate   *string `json:"InvoiceDocumentDate"`
}

type HeaderGrossWeight struct {
	HeaderGrossWeight *float32 `json:"HeaderGrossWeight"`
}

type HeaderNetWeight struct {
	HeaderNetWeight *float32 `json:"HeaderNetWeight"`
}

// Item
type DeliveryDocumentItem struct {
	DeliveryDocumentItemNumber int `json:"DeliveryDocumentItemNumber"`
}

type OrdersItem struct {
	OrderID                                       int      `json:"OrderID"`
	OrderItem                                     int      `json:"OrderItem"`
	OrderItemCategory                             string   `json:"OrderItemCategory"`
	SupplyChainRelationshipID                     int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID             *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID       *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID      *int     `json:"SupplyChainRelationshipProductionPlantID"`
	OrderItemText                                 string   `json:"OrderItemText"`
	OrderItemTextByBuyer                          string   `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller                         string   `json:"OrderItemTextBySeller"`
	Product                                       string   `json:"Product"`
	ProductStandardID                             string   `json:"ProductStandardID"`
	ProductGroup                                  *string  `json:"ProductGroup"`
	BaseUnit                                      string   `json:"BaseUnit"`
	DeliverToParty                                *int     `json:"DeliverToParty"`
	DeliverFromParty                              *int     `json:"DeliverFromParty"`
	DeliverToPlant                                *string  `json:"DeliverToPlant"`
	DeliverToPlantTimeZone                        *string  `json:"DeliverToPlantTimeZone"`
	DeliverToPlantStorageLocation                 *string  `json:"DeliverToPlantStorageLocation"`
	ProductIsBatchManagedInDeliverToPlant         *bool    `json:"ProductIsBatchManagedInDeliverToPlant"`
	BatchMgmtPolicyInDeliverToPlant               *string  `json:"BatchMgmtPolicyInDeliverToPlant"`
	DeliverToPlantBatch                           *string  `json:"DeliverToPlantBatch"`
	DeliverToPlantBatchValidityStartDate          *string  `json:"DeliverToPlantBatchValidityStartDate"`
	DeliverToPlantBatchValidityEndDate            *string  `json:"DeliverToPlantBatchValidityEndDate"`
	DeliverToPlantBatchValidityStartTime          *string  `json:"DeliverToPlantBatchValidityStartTime"`
	DeliverToPlantBatchValidityEndTime            *string  `json:"DeliverToPlantBatchValidityEndTime"`
	DeliverFromPlant                              *string  `json:"DeliverFromPlant"`
	DeliverFromPlantTimeZone                      *string  `json:"DeliverFromPlantTimeZone"`
	DeliverFromPlantStorageLocation               *string  `json:"DeliverFromPlantStorageLocation"`
	ProductIsBatchManagedInDeliverFromPlant       *bool    `json:"ProductIsBatchManagedInDeliverFromPlant"`
	BatchMgmtPolicyInDeliverFromPlant             *string  `json:"BatchMgmtPolicyInDeliverFromPlant"`
	DeliverFromPlantBatch                         *string  `json:"DeliverFromPlantBatch"`
	DeliverFromPlantBatchValidityStartDate        *string  `json:"DeliverFromPlantBatchValidityStartDate"`
	DeliverFromPlantBatchValidityEndDate          *string  `json:"DeliverFromPlantBatchValidityEndDate"`
	DeliverFromPlantBatchValidityStartTime        *string  `json:"DeliverFromPlantBatchValidityStartTime"`
	DeliverFromPlantBatchValidityEndTime          *string  `json:"DeliverFromPlantBatchValidityEndTime"`
	DeliveryUnit                                  string   `json:"DeliveryUnit"`
	StockConfirmationBusinessPartner              *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantTimeZone                *string  `json:"StockConfirmationPlantTimeZone"`
	ProductIsBatchManagedInStockConfirmationPlant *bool    `json:"ProductIsBatchManagedInStockConfirmationPlant"`
	BatchMgmtPolicyInStockConfirmationPlant       *string  `json:"BatchMgmtPolicyInStockConfirmationPlant"`
	StockConfirmationPlantBatch                   *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate  *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate    *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	StockConfirmationPlantBatchValidityStartTime  *string  `json:"StockConfirmationPlantBatchValidityStartTime"`
	StockConfirmationPlantBatchValidityEndTime    *string  `json:"StockConfirmationPlantBatchValidityEndTime"`
	OrderQuantityInBaseUnit                       float32  `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit                   float32  `json:"OrderQuantityInDeliveryUnit"`
	StockConfirmationPolicy                       *string  `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                       *string  `json:"StockConfirmationStatus"`
	ConfirmedOrderQuantityInBaseUnit              *float32 `json:"ConfirmedOrderQuantityInBaseUnit"`
	ItemWeightUnit                                *string  `json:"ItemWeightUnit"`
	ProductGrossWeight                            *float32 `json:"ProductGrossWeight"`
	ItemGrossWeight                               *float32 `json:"ItemGrossWeight"`
	ProductNetWeight                              *float32 `json:"ProductNetWeight"`
	ItemNetWeight                                 *float32 `json:"ItemNetWeight"`
	InternalCapacityQuantity                      *float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit                  *string  `json:"InternalCapacityQuantityUnit"`
	NetAmount                                     *float32 `json:"NetAmount"`
	TaxAmount                                     *float32 `json:"TaxAmount"`
	GrossAmount                                   *float32 `json:"GrossAmount"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string  `json:"ProductionPlant"`
	ProductionPlantTimeZone                       *string  `json:"ProductionPlantTimeZone"`
	ProductionPlantStorageLocation                *string  `json:"ProductionPlantStorageLocation"`
	ProductIsBatchManagedInProductionPlant        *bool    `json:"ProductIsBatchManagedInProductionPlant"`
	BatchMgmtPolicyInProductionPlant              *string  `json:"BatchMgmtPolicyInProductionPlant"`
	ProductionPlantBatch                          *string  `json:"ProductionPlantBatch"`
	ProductionPlantBatchValidityStartDate         *string  `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityEndDate           *string  `json:"ProductionPlantBatchValidityEndDate"`
	ProductionPlantBatchValidityStartTime         *string  `json:"ProductionPlantBatchValidityStartTime"`
	ProductionPlantBatchValidityEndTime           *string  `json:"ProductionPlantBatchValidityEndTime"`
	Incoterms                                     *string  `json:"Incoterms"`
	TransactionTaxClassification                  string   `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry         string   `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry       string   `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                      string   `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                        string   `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup                 string   `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                                  string   `json:"PaymentTerms"`
	PaymentMethod                                 string   `json:"PaymentMethod"`
	Project                                       *string  `json:"Project"`
	TaxCode                                       *string  `json:"TaxCode"`
	TaxRate                                       *float32 `json:"TaxRate"`
	CountryOfOrigin                               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                       *string  `json:"CountryOfOriginLanguage"`
}

type ItemIsBillingRelevant struct {
	ItemIsBillingRelevant bool `json:"ItemIsBillingRelevant"`
}

// Partner
type Partner struct {
	OrderID                 int     `json:"OrderID"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}

// Address
type Address struct {
	OrderID     int     `json:"OrderID"`
	AddressID   int     `json:"AddressID"`
	PostalCode  *string `json:"PostalCode"`
	LocalRegion *string `json:"LocalRegion"`
	Country     *string `json:"Country"`
	District    *string `json:"District"`
	StreetName  *string `json:"StreetName"`
	CityName    *string `json:"CityName"`
	Building    *string `json:"Building"`
	Floor       *int    `json:"Floor"`
	Room        *int    `json:"Room"`
}

type AddressMaster struct {
	AddressID         int     `json:"AddressID"`
	ValidityEndDate   string  `json:"ValidityEndDate"`
	ValidityStartDate string  `json:"ValidityStartDate"`
	PostalCode        string  `json:"PostalCode"`
	LocalRegion       string  `json:"LocalRegion"`
	Country           string  `json:"Country"`
	GlobalRegion      string  `json:"GlobalRegion"`
	TimeZone          string  `json:"TimeZone"`
	District          *string `json:"District"`
	StreetName        string  `json:"StreetName"`
	CityName          string  `json:"CityName"`
	Building          *string `json:"Building"`
	Floor             *int    `json:"Floor"`
	Room              *int    `json:"Room"`
}

type CalculateAddressIDKey struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
}

type CalculateAddressIDQueryGets struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
	LatestNumber             *int   `json:"LatestNumber"`
}

type CalculateAddressID struct {
	AddressIDLatestNumber *int `json:"AddressIDLatestNumber"`
	AddressID             int  `json:"AddressID"`
}

// 日付等の処理
type CreationDate struct {
	CreationDate *string `json:"CreationDate"`
}

type LastChangeDate struct {
	LastChangeDate *string `json:"LastChangeDate"`
}

type CreationTime struct {
	CreationTime *string `json:"CreationTime"`
}

type LastChangeTime struct {
	LastChangeTime *string `json:"LastChangeTime"`
}
