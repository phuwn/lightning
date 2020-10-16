package handler

import (
	"github.com/labstack/echo"
	"github.com/phuwn/lightning/src/handler/product"
)

func productRoutes(r *echo.Echo) {
	g := r.Group("/product")
	{
		g.GET("", getAllProduct)
	}
}

func getAllProduct(c echo.Context) error {
	products, err := product.GetAllProduct(c)
	if err != nil {
		return err
	}
	return JSON(c, 200, products)
}
