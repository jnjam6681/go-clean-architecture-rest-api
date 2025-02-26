package utils

import "github.com/labstack/echo/v4"

var (
	DebugMode bool
)

func JSONSuccessResponse(c echo.Context, status int, message string, data interface{}) error {
	response := map[string]interface{}{
		"success": true,
		"message": message,
		"data":    data,
	}
	return c.JSON(status, response)
}

func JSONErrorResponse(c echo.Context, status int, message string, err error) error {
	response := map[string]interface{}{
		"success": false,
		"message": message,
	}

	if DebugMode && err != nil {
		response["error"] = err.Error()
	}
	return c.JSON(status, response)
}
