package payment

import (
	"github.com/labstack/echo"

	"github.com/phuwn/tools/db"

	"github.com/phuwn/lightning/src/model"
)

type paymentPGStore struct{}

// NewStore - create new payment store
func NewStore() Store {
	return &paymentPGStore{}
}

func (s paymentPGStore) Get(c echo.Context, id string) (*model.Payment, error) {
	tx := db.GetTxFromCtx(c)
	var res model.Payment
	return &res, tx.Where("id = ?", id).
		Preload("Items").
		Preload("Items.Product").
		First(&res).Error
}

func (s paymentPGStore) GetUserPayments(c echo.Context, uid string) ([]*model.Payment, error) {
	tx := db.GetTxFromCtx(c)
	res := []*model.Payment{}
	return res, tx.Where("user_id = ?", uid).
		Preload("Items").
		Preload("Items.Product").
		Find(&res).Error
}

func (s paymentPGStore) Create(c echo.Context, payment *model.Payment) error {
	tx := db.GetTxFromCtx(c)
	return tx.Create(payment).Error
}

func (s paymentPGStore) Delete(c echo.Context, pid string, uid string) error {
	tx := db.GetTxFromCtx(c)
	return tx.Where("user_id = ? and id = ?", uid, pid).Delete(&model.Payment{}).Error
}
