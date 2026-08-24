package product

type ProductResponse struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImageURL *string `json:"image_url"`
	Category string `json:"catgory"`
	SellerName string `json:"seller_name"`
}