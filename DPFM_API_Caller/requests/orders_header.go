package requests

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
