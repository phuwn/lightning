package store

import (
	"github.com/phuwn/lightning/src/store/payment"
	"github.com/phuwn/lightning/src/store/product"
	"github.com/phuwn/lightning/src/store/user"
)

// Store - server store struct
type Store struct {
	User    user.Store
	Product product.Store
	Payment payment.Store
}

// New - create new store variable
func New() *Store {
	return &Store{
		User:    user.NewStore(),
		Product: product.NewStore(),
		Payment: payment.NewStore(),
	}
}
