package entities

import "time"

type TransactionStatus uint

// transaction status enum
const (
	Paid TransactionStatus = iota
	Unpaid
)

///////////////////////////////////////////////////////////
// Base Transaction Model
///////////////////////////////////////////////////////////

type BaseTransaction struct {
	ID     string            `json:"transaction_id"`
	Status TransactionStatus `json:"status"`
}

type baseTransactionProduct struct {
	Quantity int `json:"quantity"`
}

///////////////////////////////////////////////////////////
// Transaction Model
// INFO: this model will be used for response API / GQL
///////////////////////////////////////////////////////////

type TransactionProduct struct {
	baseTransactionProduct
	Date    time.Time `json:"transaction_date"`
	Product Product   `json:"product"`
}

type Transaction struct {
	BaseTransaction
	User     User                 `json:"user"`
	Products []TransactionProduct `json:"transaction_products"`
}

///////////////////////////////////////////////////////////
// Transaction Request Model
// INFO: this model will be used for parameters API / GQL
///////////////////////////////////////////////////////////

type TransactionProductRequest struct {
	baseTransactionProduct
	ProductID string `json:"product_id"`
}

type TransactionRequest struct {
	BaseTransaction
	UserID   string                      `json:"user_id"`
	Products []TransactionProductRequest `json:"transaction_products"`
}
