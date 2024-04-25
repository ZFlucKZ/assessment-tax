package middlewares

import "github.com/labstack/echo/v4"

func BasicAuthMiddleware(username, password string, c echo.Context) (bool, error) {
	if username == "adminTax" && password == "admin!" {
		return true, nil
	}
	return false, nil
}