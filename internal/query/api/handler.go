package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var mySigningKey = []byte("secret") // подпись для проверки токена

type Claims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Необходимая подпись не обнаружена")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("Токен невалиден")
	}
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Отсутствует необходимый для авторизации заголовок")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Неверный формат авторизационного заголовка")
		}
		tokenString := parts[1]

		_, err := ValidateToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Cрок действия токена истёк"))
		}

		return next(c)
	}
}

func (srv *Server) PostName(writer echo.Context) error {
	name := writer.QueryParam("name")

	if name == "" {
		return writer.String(400, "Попробуй ввести своё имя через query-параметр 'name'")
	}

	err := srv.uc.SetName(name)
	if err != nil {
		return writer.String(500, err.Error())
	}

	return writer.String(201, "Имя было изменено на "+name)
}

func (srv *Server) GetName(writer echo.Context) error {
	answer, err := srv.uc.FetchName()

	if err != nil {
		return writer.String(500, err.Error())
	}

	return writer.String(200, "Привет, "+answer+"!")
}
