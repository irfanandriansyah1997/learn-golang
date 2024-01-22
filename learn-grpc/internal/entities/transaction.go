package entities

type TransactionStatus uint

// transaction status enum
const (
	Paid TransactionStatus = iota
	Unpaid
)

type TransactionProduct struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type Transaction struct {
	ID       uint                 `json:"transaction_id"`
	Date     Date                 `json:"transaction_date"`
	Status   TransactionStatus    `json:"status"`
	UserID   uint                 `json:"user_id"`
	Products []TransactionProduct `json:"transaction_prpoducts"`
}
