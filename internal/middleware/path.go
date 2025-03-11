package middleware

import (
	"net/http"

	"github.com/jnjam6681/go-clean-architecture-rest-api/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (mw *MiddlewareManager) NotFoundMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			if he, ok := err.(*echo.HTTPError); ok && he.Code == http.StatusNotFound {
				return utils.JSONErrorResponse(c, http.StatusNotFound, "Path not found", errors.New("The requested resource was not found on this server"))
			}
		}
		return err
	}
}
