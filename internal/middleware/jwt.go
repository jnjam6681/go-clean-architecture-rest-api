package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jnjam6681/go-clean-architecture-rest-api/internal/model"
	"github.com/labstack/echo/v4"
)

func (mw *MiddlewareManager) JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("x-test-auth-x")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
		}

		token, err := jwt.ParseWithClaims(tokenString, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(mw.cfg.Server.HostKey), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		if claims, ok := token.Claims.(*model.JWTClaims); ok && token.Valid {
			expireToken := "30m"
			if mw.cfg.Server.ExpireToken != "" {
				expireToken = mw.cfg.Server.ExpireToken
			}
			duration, err := time.ParseDuration(expireToken)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Error parsing duration")
			}

			expirationTime := claims.IssuedAt.Time.Unix() + int64(duration.Seconds())
			if time.Now().Unix() > expirationTime {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token is expired")
			}

			hostname, err := os.Hostname()
			if err != nil {
				mw.logger.Errorf("Error getting hostname:", err)
			}

			if hostname != claims.Hostname {
				return echo.NewHTTPError(http.StatusForbidden, "Access denied")
			}
		}
		return next(c)
	}
}
