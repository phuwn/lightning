package payment

import (
	"github.com/labstack/echo"
	"github.com/phuwn/lightning/src/model"
	"github.com/phuwn/lightning/src/server"
	"github.com/phuwn/tools/errors"
)

// GetUserPayments - take all user's payments
func GetUserPayments(c echo.Context, uid string) ([]*model.Payment, error) {
	cfg := server.GetServerCfg()
	res, err := cfg.Store().Payment.GetUserPayments(c, uid)
	if err != nil {
		return nil, errors.Customize(err, 500, "failed to get user's payments")
	}
	return res, nil
}

// CreateUserPayment - create new user's payment
func CreateUserPayment(c echo.Context, payment *model.Payment) error {
	cfg := server.GetServerCfg()
	err := cfg.Store().Payment.Create(c, payment)
	if err != nil {
		return errors.Customize(err, 500, "failed to create user's payment")
	}
	return nil
}

// DeleteUserPayment - delete user's payment by id
func DeleteUserPayment(c echo.Context, pid string, uid string) error {
	cfg := server.GetServerCfg()
	err := cfg.Store().Payment.Delete(c, pid, uid)
	if err != nil {
		return errors.Customize(err, 500, "failed to delete user's payment")
	}
	return nil
}
