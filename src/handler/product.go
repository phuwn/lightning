package handler

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/labstack/echo"
	"github.com/phuwn/lightning/src/handler/product"
	"github.com/phuwn/lightning/src/model"
	"github.com/phuwn/tools/errors"
)

func productRoutes(r *echo.Echo) {
	g := r.Group("/product")
	{
		g.GET("", getAllProduct)
		g.POST("", newProduct)
		g.PUT("/:id", updateProduct)
		g.DELETE("/:id", removeProduct)
	}
}

func getAllProduct(c echo.Context) error {
	products, err := product.GetAllProduct(c)
	if err != nil {
		return err
	}
	return JSON(c, 200, products)
}

// ProductRequest - product handling form
type ProductRequest struct {
	Name  *string  `json:"name"`
	Price *float64 `json:"price"`
	Photo *string  `json:"photo"`
}

func newProduct(c echo.Context) error {
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return errors.Customize(err, 400, "invalid body")
	}
	form := &ProductRequest{}
	err = json.Unmarshal(b, form)
	if err != nil {
		return errors.Customize(err, 400, "wrong json form to create product")
	}
	if form.Name == nil {
		return errors.New("product name is required", 400)
	}
	if form.Price == nil {
		return errors.New("product price is required", 400)
	}
	p := &model.Product{Name: *form.Name, Price: *form.Price}
	if form.Photo != nil {
		p.Photo = *form.Photo
	}
	err = product.CreateProduct(c, p)
	if err != nil {
		return err
	}
	return JSON(c, 201, p)
}

func updateProduct(c echo.Context) error {
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return errors.Customize(err, 400, "invalid body")
	}
	form := &ProductRequest{}
	err = json.Unmarshal(b, form)
	if err != nil {
		return errors.Customize(err, 400, "wrong json form to create product")
	}
	if form.Name == nil && form.Price == nil && form.Photo == nil {
		return errors.New("no change is listed", 400)
	}
	pid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.Customize(err, 400, "product id have to be a number")
	}
	p, err := product.GetProductByID(c, pid)
	if err != nil {
		return err
	}

	if form.Name != nil {
		p.Name = *form.Name
	}
	if form.Photo != nil {
		p.Photo = *form.Photo
	}
	if form.Price != nil {
		p.Price = *form.Price
	}
	err = product.UpdateProduct(c, p)
	if err != nil {
		return err
	}
	return JSON(c, 200, p)
}

func removeProduct(c echo.Context) error {
	pid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.Customize(err, 400, "product id have to be a number")
	}
	err = product.DeleteProduct(c, pid)
	if err != nil {
		return err
	}
	return JSON(c, 200, nil)
}
