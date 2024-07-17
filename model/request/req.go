package request

type Register struct {
	Email    string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Login struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" validate:"required"`
}

type Cart struct {
	ProductId string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
}

type CartItem struct {
	CartItemId string `json:"cart_item_id" validate:"required"`
}

type Checkout struct {
	CartId string `json:"cart_id" validate:"required"`
}
