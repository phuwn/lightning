package user

import (
	"github.com/labstack/echo"

	"github.com/phuwn/lightning/src/model"
)

// Store - user store interface
type Store interface {
	Get(c echo.Context, id string) (*model.User, error)
	Create(c echo.Context, user *model.User) error
	GetByEmail(c echo.Context, email string) (*model.User, error)
	Save(c echo.Context, user *model.User) error
}
