package entities

type AssetRatio string

// transaction status enum
const (
	MobileRatio  AssetRatio = "320w"
	TabletRatio  AssetRatio = "640w"
	DesktopRatio AssetRatio = "800w"
)

type OptimizeAsset struct {
	Size    AssetRatio `json:"sizes"`
	UrlPath string     `json:"url"`
}

type Asset struct {
	UrlPath                string          `json:"url"`
	OptimizeAsset          []OptimizeAsset `json:"optimize_assets"`
	FormattedOptimizeAsset string          `json:"formatted_optimize_assets"`
}
