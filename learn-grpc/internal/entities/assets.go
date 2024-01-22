package entities

type OptimizeAsset struct {
	Size    string `json:"sizes"`
	UrlPath string `json:"url"`
}

type Asset struct {
	UrlPath                string          `json:"url"`
	OptimizeAsset          []OptimizeAsset `json:"optimize_assets"`
	FormattedOptimizeAsset string          `json:"formatted_optimize_assets"`
}
