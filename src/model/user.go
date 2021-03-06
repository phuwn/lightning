package model

import (
	"fmt"

	"github.com/labstack/echo"
)

const (
	MemberRole = iota + 1
	AdminRole
)

// User data model
type User struct {
	Base
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	Role   int    `json:"role"`

	AccessToken *string `json:"access_token,omitempty" sql:"-"`
}

const (
	uidKey string = "uid"
)

// SetUserIDToCtx - set a uid to echo context for later user's session usage
func SetUserIDToCtx(c echo.Context, uid string) {
	c.Set(uidKey, uid)
}

// GetUserIDFromCtx - get uid from echo context
func GetUserIDFromCtx(c echo.Context) string {
	return fmt.Sprintf("%v", c.Get(uidKey))
}
