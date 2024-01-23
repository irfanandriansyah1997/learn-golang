package utils

import (
	"fmt"
	"learn-grpc/internal/entities"
	"net/url"
	"slices"
	"strings"
)

var eligibleExtensions = []string{"gif", "jpg", "jpeg", "png"}
var imageRatios = []entities.AssetRatio{entities.MobileRatio, entities.TabletRatio, entities.DesktopRatio}

func isImage(imageURL string) bool {
	splitURLPath := strings.Split(imageURL, ".")
	extension := splitURLPath[len(splitURLPath)-1]

	return slices.Contains(eligibleExtensions, extension)
}

func generateAssetByRatio(imageURL string, ratio entities.AssetRatio) entities.OptimizeAsset {
	splitURLPath := strings.Split(imageURL, ".")
	splitURLPath[len(splitURLPath)-2] = fmt.Sprintf("%s-%s", splitURLPath[len(splitURLPath)-2], ratio)

	return entities.OptimizeAsset{
		Size:    ratio,
		UrlPath: strings.Join(splitURLPath, "."),
	}
}

func GenerateOptimizeAsset(imageURL string) (*entities.Asset, error) {
	_, err := url.ParseRequestURI(imageURL)
	isImageUrl := isImage(imageURL)

	if err == nil && isImageUrl {
		// TODO: need to create func for check asset on S3 & GCS

		formattedOptimizeAsset := make([]entities.OptimizeAsset, 0)
		formattedOptimizeAssetHTML := make([]string, 0)

		for _, ratio := range imageRatios {
			result := generateAssetByRatio(imageURL, ratio)
			formattedOptimizeAsset = append(formattedOptimizeAsset, result)
			formattedOptimizeAssetHTML = append(formattedOptimizeAssetHTML, fmt.Sprintf("%s %s", result.UrlPath, ratio))
		}

		return &entities.Asset{
			UrlPath:                imageURL,
			OptimizeAsset:          formattedOptimizeAsset,
			FormattedOptimizeAsset: strings.Join(formattedOptimizeAssetHTML, ", "),
		}, nil
	}

	if err != nil {
		return nil, NewNotFoundError(fmt.Sprintf("this image %s is not available on S3 & GCS", imageURL))
	}

	return nil, fmt.Errorf("this url %s is not image", imageURL)

}
