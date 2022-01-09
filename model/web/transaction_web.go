package web

type CreateTransactionRequest struct {
	ProductsID      []int  `validate:"required" json:"products_id,omitempty" form:"products_id"`
	SellerID        uint   `validate:"required" json:"seller_id,omitempty" form:"seller_id"`
	CustomerID      uint   `validate:"required" json:"customer_id,omitempty" form:"customer_id"`
	ShippingAddress string `validate:"required" json:"shipping_address,omitempty" form:"shipping_address"`
	PaymentMethode  string `validate:"required" json:"payment_methode,omitempty" form:"payment_methode"`
	Status          string `json:"status,omitempty"`
}

type UpdateStatusTransactionRequest struct {
	ID     uint   `json:"id,omitempty" form:"id" validate:"required"`
	Status string `json:"status,omitempty" form:"status" validate:"required"`
}

type TransactionResponse struct {
	ID              uint              `json:"id,omitempty"`
	Products        []ProductResponse `json:"products,omitempty"`
	SellerID        uint              `json:"seller_id,omitempty"`
	CustomerID      uint              `json:"customer_id,omitempty"`
	ShippingAddress string            `json:"shipping_address,omitempty"`
	PaymentMethode  string            `json:"payment_methode,omitempty"`
	SubTotal        int               `json:"sub_total,omitempty"`
	Qty             int               `json:"qty,omitempty"`
	Status          string            `json:"status,omitempty"`
}
