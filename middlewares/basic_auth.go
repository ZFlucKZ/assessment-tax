package middlewares

import (
	"github.com/ZFlucKZ/assessment-tax/config"
	"github.com/labstack/echo/v4"
)

func BasicAuthMiddleware(username, password string, c echo.Context) (bool, error) {
	if username == config.Env.AdminUsername && password == config.Env.AdminPassword {
		return true, nil
	}
	return false, nil
}