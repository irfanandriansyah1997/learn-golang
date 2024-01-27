package repositories

import (
	"context"
	"learn-grpc/internal/entities"
	"learn-grpc/internal/utils"
	"time"

	"github.com/google/uuid"
)

type TransactionRepo interface {
	Create(context.Context, entities.TransactionRequest) (*entities.Transaction, error)
	FindAll(context.Context) []entities.Transaction
	FindByID(context.Context, string) (*entities.Transaction, error)
}

type _TransactionProduct struct {
	entities.BaseTransactionProduct
	ProductID string
}

type _Transaction struct {
	entities.BaseTransaction
	UserID   string `json:"user_id"`
	Products []_TransactionProduct
}

type _TransactionRepoImpl struct {
	transaction []_Transaction
	productRepo ProductRepo
	userRepo    UserRepo
}

func NewTransactionRepo(productRepo ProductRepo, userRepo UserRepo) TransactionRepo {
	return &_TransactionRepoImpl{
		transaction: make([]_Transaction, 0),
		productRepo: productRepo,
		userRepo:    userRepo,
	}
}

func (t *_TransactionRepoImpl) _normalizeTransactionProduct(ctx context.Context, arg _TransactionProduct) (*entities.TransactionProduct, error) {
	var transactionProduct *entities.TransactionProduct
	formattedProduct, err := t.productRepo.FindByID(ctx, arg.ProductID)

	if formattedProduct != nil && err == nil {
		transactionProduct = &entities.TransactionProduct{}
		transactionProduct.Quantity = arg.Quantity
		transactionProduct.Product = *formattedProduct
	}

	return transactionProduct, err
}

type AsyncProduct struct {
	result *entities.TransactionProduct
	err    error
}

func (t *_TransactionRepoImpl) _normalizeTransaction(ctx context.Context, arg _Transaction) (*entities.Transaction, error) {
	fetchProductChannel := make(chan AsyncProduct)
	defer close(fetchProductChannel)

	formattedTransaction := entities.Transaction{}
	formattedTransaction.ID = arg.ID
	formattedTransaction.Status = arg.Status
	formattedTransaction.Date = arg.Date

	///////////////////////////////////////////////////////////
	// Mapping user operation
	///////////////////////////////////////////////////////////

	formattedUser, err := t.userRepo.FindByID(ctx, arg.UserID)

	if err != nil {
		return nil, err
	}

	if err == nil && formattedUser != nil {
		formattedTransaction.User = *formattedUser
	}

	///////////////////////////////////////////////////////////
	// Mapping transaction products operation
	///////////////////////////////////////////////////////////

	fetchProduct := func(channel chan<- AsyncProduct, transaction _TransactionProduct) {
		result, err := t._normalizeTransactionProduct(ctx, transaction)
		channel <- AsyncProduct{result, err}
	}

	formattedTransactionProduct := make([]entities.TransactionProduct, 0)
	for _, item := range arg.Products {
		go fetchProduct(fetchProductChannel, item)
	}

	occurenceProductChannel := 0
	for product := range fetchProductChannel {
		if product.result != nil && product.err == nil {
			formattedTransactionProduct = append(formattedTransactionProduct, *product.result)
		}

		occurenceProductChannel++

		if occurenceProductChannel == len(arg.Products) {
			break
		}
	}

	if len(formattedTransactionProduct) == len(arg.Products) {
		formattedTransaction.Products = formattedTransactionProduct
	} else {
		return nil, utils.NewNotFoundError("some product from transaction detail not found")
	}

	return &formattedTransaction, nil
}

func (t *_TransactionRepoImpl) Create(ctx context.Context, arg entities.TransactionRequest) (*entities.Transaction, error) {
	id := uuid.New().String()
	newTransaction := _Transaction{}

	newTransactionProducts := make([]_TransactionProduct, 0)
	for _, item := range arg.Products {
		transactionProductItem := _TransactionProduct{}
		transactionProductItem.ProductID = item.ProductID
		transactionProductItem.Quantity = item.Quantity
		newTransactionProducts = append(newTransactionProducts, transactionProductItem)
	}

	newTransaction.ID = id
	newTransaction.Status = arg.Status
	newTransaction.UserID = arg.UserID
	newTransaction.Date = time.Now()
	newTransaction.Products = newTransactionProducts

	result, err := t._normalizeTransaction(ctx, newTransaction)
	if err == nil && result != nil {
		return result, err
	}

	return nil, err
}

func (t *_TransactionRepoImpl) FindAll(_ context.Context) []entities.Transaction {
	panic("not implemented") // TODO: Implement
}

func (t *_TransactionRepoImpl) FindByID(_ context.Context, _ string) (*entities.Transaction, error) {
	panic("not implemented") // TODO: Implement
}
