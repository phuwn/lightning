package handler

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo"
	"github.com/phuwn/lightning/src/handler/payment"
	"github.com/phuwn/lightning/src/model"
	"github.com/phuwn/tools/errors"
)

func paymentRoutes(r *echo.Echo) {
	g := r.Group("/payment")
	{
		g.GET("", getMyPayments)
		g.POST("", newPayment)
		g.DELETE("/:id", removePayment)
	}
}

func getMyPayments(c echo.Context) error {
	uid := model.GetUserIDFromCtx(c)
	payments, err := payment.GetUserPayments(c, uid)
	if err != nil {
		return err
	}
	return JSON(c, 200, payments)
}

func newPayment(c echo.Context) error {
	uid := model.GetUserIDFromCtx(c)
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return errors.Customize(err, 400, "invalid request body")
	}
	p := &model.Payment{}
	err = json.Unmarshal(b, p)
	if err != nil {
		return errors.Customize(err, 400, "wrong request form")
	}
	if len(p.Items) == 0 {
		return errors.New("payment need some items", 400)
	}
	p.UserID = uid
	err = payment.CreateUserPayment(c, p)
	if err != nil {
		return err
	}
	return JSON(c, 201, p)
}

func removePayment(c echo.Context) error {
	uid := model.GetUserIDFromCtx(c)
	id := c.Param("id")
	err := payment.DeleteUserPayment(c, id, uid)
	if err != nil {
		return err
	}
	return JSON(c, 200, nil)
}
