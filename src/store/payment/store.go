package payment

import (
	"github.com/labstack/echo"

	"github.com/phuwn/lightning/src/model"
)

// Store - payment store interface
type Store interface {
	Get(c echo.Context, id string) (*model.Payment, error)
	GetUserPayments(c echo.Context, uid string) ([]*model.Payment, error)
	Create(c echo.Context, payment *model.Payment) error
	Delete(c echo.Context, pid string, uid string) error
}
