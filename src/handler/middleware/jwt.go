package middleware

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/phuwn/tools/errors"

	"github.com/phuwn/lightning/src/model"
)

var authPath = map[string][]string{"/product": []string{"POST", "PUT", "DELETE"}}

func authenticate(c echo.Context) error {
	auth := c.Request().Header.Get("Authorization")
	if !strings.Contains(auth, "Bearer ") {
		return errors.New("invalid auth method", 401)
	}
	token := auth[7:]
	if token == "" {
		return errors.New("missing access_token", 401)
	}
	uid, err := model.VerifyUserSession(token)
	if err != nil {
		return err
	}
	model.SetUserIDToCtx(c, uid)
	return nil
}

// WithAuth - authentication middleware
func WithAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if methods, ok := authPath[c.Request().RequestURI]; ok {
			for _, v := range methods {
				if v == c.Request().Method {
					err := authenticate(c)
					if err != nil {
						return errors.Customize(err, 401, "invalid token")
					}
					break
				}
			}
		}

		return next(c)
	}
}
