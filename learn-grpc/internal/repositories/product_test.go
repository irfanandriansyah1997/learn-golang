package repositories_test

import (
	"context"
	"learn-grpc/internal/entities"
	"learn-grpc/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

var repo = repositories.NewProductRepo()

func TestProductRepo(t *testing.T) {
	mockProduct := entities.ProductRequest{
		BaseProduct: entities.BaseProduct{
			Name:        "Nescafe gold refill 170g",
			Description: "Tersedia juga nescafe gold altarica dan colombia di etalase kami. silahkan cek dan order..😇",
		},
		Photos: []string{"https://images.tokopedia.net/sample-image.jpg"},
		Price:  112500,
	}
	mockProduct2 := entities.ProductRequest{
		BaseProduct: entities.BaseProduct{
			Name:        "Nescafe latte caramel & mocha",
			Description: "Nescafe caramel 1 bks isi 20 sachets",
		},
		Photos: []string{"https://images.tokopedia.net/sample-image-2.jpg"},
		Price:  99900,
	}
	var selectedProductId string

	t.Run("testing create operation", func(t *testing.T) {
		// Check result crud operation
		resultCreate := repo.Create(context.TODO(), mockProduct)
		assert.Equal(t, resultCreate.Name, "Nescafe gold refill 170g")
		assert.Equal(t, resultCreate.Description, "Tersedia juga nescafe gold altarica dan colombia di etalase kami. silahkan cek dan order..😇")
		assert.EqualValues(t, resultCreate.Photos[0], entities.Asset{
			UrlPath: "https://images.tokopedia.net/sample-image.jpg",
			OptimizeAsset: []entities.OptimizeAsset{
				{Size: entities.MobileRatio, UrlPath: "https://images.tokopedia.net/sample-image-320w.jpg"},
				{Size: entities.TabletRatio, UrlPath: "https://images.tokopedia.net/sample-image-640w.jpg"},
				{Size: entities.DesktopRatio, UrlPath: "https://images.tokopedia.net/sample-image-800w.jpg"},
			},
			FormattedOptimizeAsset: "https://images.tokopedia.net/sample-image-320w.jpg 320w, https://images.tokopedia.net/sample-image-640w.jpg 640w, https://images.tokopedia.net/sample-image-800w.jpg 800w",
		})

		// Compare result create operation & get all data
		resultFindAll := repo.FindAll(context.TODO())
		assert.Len(t, resultFindAll, 1)
		assert.EqualValues(t, resultFindAll[0], resultCreate)

		selectedProductId = resultCreate.ID
	})

	t.Run("testing find by id operation", func(t *testing.T) {
		resultFindById, err := repo.FindByID(context.TODO(), selectedProductId)
		assert.Nil(t, err)
		assert.Equal(t, resultFindById.Name, "Nescafe gold refill 170g")
		assert.Equal(t, resultFindById.Description, "Tersedia juga nescafe gold altarica dan colombia di etalase kami. silahkan cek dan order..😇")
		assert.EqualValues(t, resultFindById.Photos[0], entities.Asset{
			UrlPath: "https://images.tokopedia.net/sample-image.jpg",
			OptimizeAsset: []entities.OptimizeAsset{
				{Size: entities.MobileRatio, UrlPath: "https://images.tokopedia.net/sample-image-320w.jpg"},
				{Size: entities.TabletRatio, UrlPath: "https://images.tokopedia.net/sample-image-640w.jpg"},
				{Size: entities.DesktopRatio, UrlPath: "https://images.tokopedia.net/sample-image-800w.jpg"},
			},
			FormattedOptimizeAsset: "https://images.tokopedia.net/sample-image-320w.jpg 320w, https://images.tokopedia.net/sample-image-640w.jpg 640w, https://images.tokopedia.net/sample-image-800w.jpg 800w",
		})

		resultFindById, err = repo.FindByID(context.TODO(), "random product id")
		assert.ErrorContains(t, err, "[Not Found]: product random product id is not found")
		assert.Nil(t, resultFindById)
	})

	t.Run("testing update operation", func(t *testing.T) {
		resultUpdate := repo.Update(context.TODO(), selectedProductId, mockProduct2)

		assert.Equal(t, resultUpdate.Name, "Nescafe latte caramel & mocha")
		assert.Equal(t, resultUpdate.Description, "Nescafe caramel 1 bks isi 20 sachets")
		assert.EqualValues(t, resultUpdate.Photos[0], entities.Asset{
			UrlPath: "https://images.tokopedia.net/sample-image-2.jpg",
			OptimizeAsset: []entities.OptimizeAsset{
				{Size: entities.MobileRatio, UrlPath: "https://images.tokopedia.net/sample-image-2-320w.jpg"},
				{Size: entities.TabletRatio, UrlPath: "https://images.tokopedia.net/sample-image-2-640w.jpg"},
				{Size: entities.DesktopRatio, UrlPath: "https://images.tokopedia.net/sample-image-2-800w.jpg"},
			},
			FormattedOptimizeAsset: "https://images.tokopedia.net/sample-image-2-320w.jpg 320w, https://images.tokopedia.net/sample-image-2-640w.jpg 640w, https://images.tokopedia.net/sample-image-2-800w.jpg 800w",
		})

		// Compare result update operation & get all data
		resultFindAll := repo.FindAll(context.TODO())
		assert.Len(t, resultFindAll, 1)
		assert.EqualValues(t, resultFindAll[0], *resultUpdate)
	})

	t.Run("testing update operation with product id not found", func(t *testing.T) {
		defer func() {
			err, ok := recover().(error)

			if ok {
				assert.ErrorContains(t, err, "[Not Found]: product random id is not found")
			} else {
				t.Errorf("This method should be throw error")
			}
		}()

		resultUpdate := repo.Update(context.TODO(), "random id", mockProduct2)
		assert.Nil(t, resultUpdate)
	})

	t.Run("testing delete operation with product id not found", func(t *testing.T) {
		defer func() {
			err, ok := recover().(error)

			if ok {
				assert.ErrorContains(t, err, "[Not Found]: product random id is not found")
			} else {
				t.Errorf("This method should be throw error")
			}
		}()

		resultUpdate := repo.Delete(context.TODO(), "random id")

		assert.Equal(t, resultUpdate, true)

		// Compare result delete operation & get all data
		resultFindAll := repo.FindAll(context.TODO())
		assert.Len(t, resultFindAll, 1)
	})

	t.Run("testing delete operation", func(t *testing.T) {
		resultUpdate := repo.Delete(context.TODO(), selectedProductId)

		assert.Equal(t, resultUpdate, true)

		// Compare result delete operation & get all data
		resultFindAll := repo.FindAll(context.TODO())
		assert.Len(t, resultFindAll, 0)
	})
}
