package requests

type InvoiceDocumentDate struct {
	PlannedGoodsIssueDate string  `json:"PlannedGoodsIssueDate"`
	InvoiceDocumentDate   *string `json:"InvoiceDocumentDate"`
}
