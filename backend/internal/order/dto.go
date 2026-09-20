package order

type CreateOrderRequest struct {
	Items []CreateOrderItemRequest `json:"items"`
}

type CreateOrderItemRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
