package middleware

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/phuwn/tools/errors"

	"github.com/phuwn/lightning/src/model"
)

var authPath = map[string]int{
	"GET-user":       1,
	"PUT-user":       1,
	"POST-product":   2,
	"PUT-product":    2,
	"DELETE-product": 2,
	"GET-payment":    1,
	"POST-payment":   1,
	"DELETE-payment": 1,
}

func authenticate(c echo.Context, authorizeCode int) error {
	auth := c.Request().Header.Get("Authorization")
	if !strings.Contains(auth, "Bearer ") {
		return errors.New("invalid auth method", 401)
	}
	token := auth[7:]
	if token == "" {
		return errors.New("missing access_token", 401)
	}
	tokenInfo, err := model.VerifyUserSession(token)
	if err != nil {
		return err
	}
	if tokenInfo.Role < authorizeCode {
		return errors.New("permission denied", 403)
	}
	model.SetUserIDToCtx(c, tokenInfo.UserID)
	return nil
}

// WithAuth - authentication middleware
func WithAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		URIs := strings.Split(c.Request().RequestURI, "/")
		if len(URIs) < 2 {
			return errors.New("invalid path", 404)
		}
		if role, ok := authPath[c.Request().Method+"-"+URIs[1]]; ok {
			err := authenticate(c, role)
			if err != nil {
				return err
			}
		}

		return next(c)
	}
}
