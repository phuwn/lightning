package product

import (
	"github.com/labstack/echo"

	"github.com/phuwn/lightning/src/model"
)

// Store - product store interface
type Store interface {
	Get(c echo.Context, id int) (*model.Product, error)
	GetAll(c echo.Context) ([]*model.Product, error)
	Create(c echo.Context, product *model.Product) error
	Save(c echo.Context, product *model.Product) error
	Delete(c echo.Context, pid int) error
}
