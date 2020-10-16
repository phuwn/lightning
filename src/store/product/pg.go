package product

import (
	"github.com/labstack/echo"

	"github.com/phuwn/tools/db"

	"github.com/phuwn/lightning/src/model"
)

type productPGStore struct{}

// NewStore - create new product store
func NewStore() Store {
	return &productPGStore{}
}

func (s productPGStore) Get(c echo.Context, id int) (*model.Product, error) {
	tx := db.GetTxFromCtx(c)
	var res model.Product
	return &res, tx.Where("id = ?", id).First(&res).Error
}

func (s productPGStore) GetAll(c echo.Context) ([]*model.Product, error) {
	tx := db.GetTxFromCtx(c)
	res := []*model.Product{}
	return res, tx.Find(&res).Error
}

func (s productPGStore) Create(c echo.Context, product *model.Product) error {
	tx := db.GetTxFromCtx(c)
	return tx.Create(product).Error
}

func (s productPGStore) Save(c echo.Context, product *model.Product) error {
	tx := db.GetTxFromCtx(c)
	return tx.Save(product).Error
}

func (s productPGStore) Delete(c echo.Context, pid int) error {
	tx := db.GetTxFromCtx(c)
	return tx.Delete(&model.Product{}, pid).Error
}
