package entities

type BaseProduct struct {
	ID          string `json:"id"`
	Name        string `json:"product_name"`
	Description string `json:"description"`
}

///////////////////////////////////////////////////////////
// Product Model
// INFO: this model will be used for response API / GQL
///////////////////////////////////////////////////////////

type Product struct {
	BaseProduct
	Photos []Asset `json:"photos"`
	Price  Price   `json:"price"`
}

///////////////////////////////////////////////////////////
// Product Request Model
// INFO: this model will be used for parameters API / GQL
///////////////////////////////////////////////////////////

type ProductRequest struct {
	BaseProduct
	Photos []string `json:"photos"`
	Price  int32    `json:"price"`
}
