package enytities

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"product_name"`
	Description  string `json:"description"`
	ProductPhoto Asset  `json:"photo"`
	Price        Price  `json:"price"`
}
