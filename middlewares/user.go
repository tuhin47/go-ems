package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tuhin47/go-ems/types"
)

func CurrentUserFromCtx(c echo.Context) (*types.CurrentUser, error) {
	user, ok := c.Get(ContextKeyCurrentUser).(types.CurrentUser)
	if !ok {
		return nil, fmt.Errorf("user not found in request")
	}
	return &user, nil
}
