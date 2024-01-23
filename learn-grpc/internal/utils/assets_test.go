package utils_test

import (
	"fmt"
	"learn-grpc/internal/entities"
	"learn-grpc/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleURL = "https://developer.mozilla.org/non-responsive-narrow.png"
var expectedMobileURL = "https://developer.mozilla.org/non-responsive-narrow-320w.png"
var expectedTabletURL = "https://developer.mozilla.org/non-responsive-narrow-640w.png"
var expectedDesktopURL = "https://developer.mozilla.org/non-responsive-narrow-800w.png"
var sampleWrongURL1 = "https://developer.mozilla.org/docs.pdf"
var sampleWrongURL2 = "random text example"

func TestGenerateOptimizeAsset(t *testing.T) {
	asset, err := utils.GenerateOptimizeAsset(sampleURL)

	assert.Nil(t, err)

	// check asset object
	assert.NotNil(t, asset)
	assert.Equal(t, sampleURL, asset.UrlPath)

	// check asset optimize asset array
	assert.Len(t, asset.OptimizeAsset, 3)

	mobileAsset := asset.OptimizeAsset[0]
	assert.Equal(t, expectedMobileURL, mobileAsset.UrlPath)
	assert.Equal(t, entities.MobileRatio, mobileAsset.Size)

	tabletAsset := asset.OptimizeAsset[1]
	assert.Equal(t, expectedTabletURL, tabletAsset.UrlPath)
	assert.Equal(t, entities.TabletRatio, tabletAsset.Size)

	desktopAsset := asset.OptimizeAsset[2]
	assert.Equal(t, expectedDesktopURL, desktopAsset.UrlPath)
	assert.Equal(t, entities.DesktopRatio, desktopAsset.Size)

	// check formatted optimize asset URL
	assert.Equal(
		t,
		fmt.Sprintf("%s 320w, %s 640w, %s 800w", expectedMobileURL, expectedTabletURL, expectedDesktopURL),
		asset.FormattedOptimizeAsset,
	)
}

func TestGenerateOptimizeAssetError(t *testing.T) {
	testCases := []entities.Fixture[string, error]{
		{
			Title:         "testing with url is not image URL",
			Input:         sampleWrongURL1,
			ExpectedValue: fmt.Errorf("this url %s is not image", sampleWrongURL1),
		},
		{
			Title:         "testing with incorrect format",
			Input:         sampleWrongURL2,
			ExpectedValue: utils.NewNotFoundError(fmt.Sprintf("this image %s is not available on S3 & GCS", sampleWrongURL2)),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Title, func(t *testing.T) {
			asset, err := utils.GenerateOptimizeAsset(testCase.Input)

			assert.NotNil(t, err)
			assert.Equal(t, testCase.ExpectedValue.Error(), err.Error())

			assert.Nil(t, asset)
		})
	}
}
