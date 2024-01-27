package repositories_test

import (
	"context"
	"encoding/json"
	"fmt"
	"learn-grpc/internal/entities"
	"learn-grpc/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ctx = context.TODO()
)

func TestTransactionRepo(t *testing.T) {
	var (
		userRepo                 = repositories.NewUserRepo()
		productRepo              = repositories.NewProductRepo()
		transactionRepo          = repositories.NewTransactionRepo(productRepo, userRepo)
		mockProductId   []string = make([]string, 0)
		mockUserId      string
		mockUserData    entities.User
	)

	user1 := userRepo.Create(ctx, mockUser)
	mockUserId = user1.ID
	mockUserData = user1

	product1 := productRepo.Create(ctx, mockProduct)
	mockProductId = append(mockProductId, product1.ID)
	fmt.Println(product1.ID)

	product2 := productRepo.Create(ctx, mockProduct2)
	mockProductId = append(mockProductId, product2.ID)
	fmt.Println(product2.ID)
	fmt.Println(mockProductId)

	t.Run("testing create operation", func(t *testing.T) {
		mockTransaction := entities.TransactionRequest{}
		mockTransactionProducts := make([]entities.TransactionProductRequest, 0)

		mockTransaction.UserID = mockUserId
		mockTransaction.Status = entities.Paid

		for _, productId := range mockProductId {
			temp := entities.TransactionProductRequest{}
			temp.Quantity = 10
			temp.ProductID = productId

			mockTransactionProducts = append(mockTransactionProducts, temp)
		}
		mockTransaction.Products = mockTransactionProducts
		result, err := transactionRepo.Create(ctx, mockTransaction)

		a, _ := json.Marshal(result)
		fmt.Print(string(a))

		assert.Nil(t, err)
		assert.NotNil(t, result)

		assert.Equal(t, (*result).Status, entities.Paid)
		assert.EqualValues(t, (*result).User, mockUserData)

		assert.Len(t, (*result).Products, 2)
	})
}
