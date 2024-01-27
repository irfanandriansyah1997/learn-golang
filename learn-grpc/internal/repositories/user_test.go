package repositories_test

import (
	"context"
	"learn-grpc/internal/entities"
	"learn-grpc/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockUser = entities.UserRequest{
		BaseUser: entities.BaseUser{
			Name: "John Doe",
		},
		Avatar: "https://images.tokopedia.net/sample-image.jpg",
	}
	mockUser2 = entities.UserRequest{
		BaseUser: entities.BaseUser{
			Name: "Jim Morrison",
		},
		Avatar: "https://images.tokopedia.net/sample-image-2.jpg",
	}
)

func TestUserRepo(t *testing.T) {
	var userRepo = repositories.NewUserRepo()
	var selectedUserId string

	t.Run("testing create operation", func(t *testing.T) {
		// Check result create operation
		resultCreate := userRepo.Create(context.TODO(), mockUser)

		assert.Equal(t, resultCreate.Name, "John Doe")
		assert.EqualValues(t, resultCreate.Avatar, entities.Asset{
			UrlPath: "https://images.tokopedia.net/sample-image.jpg",
			OptimizeAsset: []entities.OptimizeAsset{
				{Size: entities.MobileRatio, UrlPath: "https://images.tokopedia.net/sample-image-320w.jpg"},
				{Size: entities.TabletRatio, UrlPath: "https://images.tokopedia.net/sample-image-640w.jpg"},
				{Size: entities.DesktopRatio, UrlPath: "https://images.tokopedia.net/sample-image-800w.jpg"},
			},
			FormattedOptimizeAsset: "https://images.tokopedia.net/sample-image-320w.jpg 320w, https://images.tokopedia.net/sample-image-640w.jpg 640w, https://images.tokopedia.net/sample-image-800w.jpg 800w",
		})

		// Compare result create operation & get all data
		resultFindAll := userRepo.FindAll(context.TODO())
		assert.Len(t, resultFindAll, 1)
		assert.EqualValues(t, resultFindAll[0], resultCreate)

		selectedUserId = resultCreate.ID
	})

	t.Run("testing find by id operation", func(t *testing.T) {
		resultFindById, err := userRepo.FindByID(context.TODO(), selectedUserId)

		assert.Nil(t, err)
		assert.Equal(t, resultFindById.Name, "John Doe")
		assert.EqualValues(t, resultFindById.Avatar, entities.Asset{
			UrlPath: "https://images.tokopedia.net/sample-image.jpg",
			OptimizeAsset: []entities.OptimizeAsset{
				{Size: entities.MobileRatio, UrlPath: "https://images.tokopedia.net/sample-image-320w.jpg"},
				{Size: entities.TabletRatio, UrlPath: "https://images.tokopedia.net/sample-image-640w.jpg"},
				{Size: entities.DesktopRatio, UrlPath: "https://images.tokopedia.net/sample-image-800w.jpg"},
			},
			FormattedOptimizeAsset: "https://images.tokopedia.net/sample-image-320w.jpg 320w, https://images.tokopedia.net/sample-image-640w.jpg 640w, https://images.tokopedia.net/sample-image-800w.jpg 800w",
		})

		resultFindById, err = userRepo.FindByID(context.TODO(), "random user id")
		assert.ErrorContains(t, err, "[Not Found]: user random user id is not found")
		assert.Nil(t, resultFindById)
	})

	t.Run("testing do activity operation should be working properly", func(t *testing.T) {
		defer func() {
			_, ok := recover().(error)

			if ok {
				t.Errorf("This method should be not throwing error")
			}
		}()

		userRepo.DoActivity(context.TODO(), selectedUserId)
	})

	t.Run("testing update operation", func(t *testing.T) {
		resultUpdate := userRepo.Update(context.TODO(), selectedUserId, mockUser2)

		assert.Equal(t, resultUpdate.Name, "Jim Morrison")
		assert.EqualValues(t, resultUpdate.Avatar, entities.Asset{
			UrlPath: "https://images.tokopedia.net/sample-image-2.jpg",
			OptimizeAsset: []entities.OptimizeAsset{
				{Size: entities.MobileRatio, UrlPath: "https://images.tokopedia.net/sample-image-2-320w.jpg"},
				{Size: entities.TabletRatio, UrlPath: "https://images.tokopedia.net/sample-image-2-640w.jpg"},
				{Size: entities.DesktopRatio, UrlPath: "https://images.tokopedia.net/sample-image-2-800w.jpg"},
			},
			FormattedOptimizeAsset: "https://images.tokopedia.net/sample-image-2-320w.jpg 320w, https://images.tokopedia.net/sample-image-2-640w.jpg 640w, https://images.tokopedia.net/sample-image-2-800w.jpg 800w",
		})

		// Compare result update operation & get all data
		resultFindAll := userRepo.FindAll(context.TODO())
		assert.Len(t, resultFindAll, 1)
		assert.EqualValues(t, resultFindAll[0], *resultUpdate)
	})

	t.Run("testing update operation with user id not found", func(t *testing.T) {
		defer func() {
			err, ok := recover().(error)

			if ok {
				assert.ErrorContains(t, err, "[Not Found]: user random user id is not found")
			} else {
				t.Errorf("This method should be throw error")
			}
		}()

		resultUpdate := userRepo.Update(context.TODO(), "random user id", mockUser2)
		assert.Nil(t, resultUpdate)
	})

	t.Run("testing delete operation with user not found", func(t *testing.T) {
		defer func() {
			err, ok := recover().(error)

			if ok {
				assert.ErrorContains(t, err, "[Not Found]: user random user id is not found")
			} else {
				t.Errorf("This method should be throw error")
			}
		}()

		resultDelete := userRepo.Delete(context.TODO(), "random user id")
		assert.Equal(t, resultDelete, false)

		// Compare result delete operation & get all data
		resultFindAll := userRepo.FindAll(context.TODO())
		assert.Len(t, resultFindAll, 1)
	})

	t.Run("testing delete operation", func(t *testing.T) {
		resultDelete := userRepo.Delete(context.TODO(), selectedUserId)

		assert.Equal(t, resultDelete, true)

		// Compare result delete operation & get all data
		resultFindAll := userRepo.FindAll(context.TODO())
		assert.Len(t, resultFindAll, 0)
	})
}
