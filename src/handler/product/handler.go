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

// CreateProduct - create new product
func CreateProduct(c echo.Context, p *model.Product) error {
	cfg := server.GetServerCfg()
	err := cfg.Store().Product.Create(c, p)
	if err != nil {
		return errors.Customize(err, 500, "failed to create product")
	}
	return nil
}

// GetProductByID - get a product with the provided ID
func GetProductByID(c echo.Context, pid int) (*model.Product, error) {
	cfg := server.GetServerCfg()
	p, err := cfg.Store().Product.Get(c, pid)
	if err != nil {
		return nil, errors.Customize(err, 404, "product not found")
	}
	return p, nil
}

// UpdateProduct - update a product record
func UpdateProduct(c echo.Context, p *model.Product) error {
	cfg := server.GetServerCfg()
	err := cfg.Store().Product.Save(c, p)
	if err != nil {
		return errors.Customize(err, 500, "failed to update product")
	}
	return nil
}

// DeleteProduct - delete a product record
func DeleteProduct(c echo.Context, id int) error {
	cfg := server.GetServerCfg()
	err := cfg.Store().Product.Delete(c, id)
	if err != nil {
		return errors.Customize(err, 500, "failed to delete product")
	}
	return nil
}
