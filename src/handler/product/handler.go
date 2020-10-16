package product

import (
	"github.com/labstack/echo"
	"github.com/phuwn/lightning/src/model"
	"github.com/phuwn/lightning/src/server"
	"github.com/phuwn/tools/errors"
)

// GetAllProduct - take all product record from the db
func GetAllProduct(c echo.Context) ([]*model.Product, error) {
	cfg := server.GetServerCfg()
	res, err := cfg.Store().Product.GetAll(c)
	if err != nil {
		return nil, errors.Customize(err, 500, "failed to get all product")
	}
	return res, nil
}
