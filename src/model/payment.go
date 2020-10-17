package model

// Payment data model
type Payment struct {
	Base
	ID         string  `json:"id"`
	UserID     string  `json:"-"`
	ItemAmount int     `json:"item_amount"`
	Total      float64 `json:"total"`

	Items []*PaymentItem `json:"items"`
}

// PaymentItem data model
type PaymentItem struct {
	PaymentID string `json:"-"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`

	Product *Product `json:"product,omitempty"`
}
