package api

import (
	"fmt"
	"net/http"
	"strconv"
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

func (srv *Server) GetCount(writer echo.Context) error {
	answer, err := srv.uc.FetchCount()
	if err != nil {
		return writer.String(500, err.Error())
	}

	return writer.String(200, "Значение счётчика: "+answer)
}

func (srv *Server) SetCount(writer echo.Context) error {
	input := struct {
		Msg *string `json:"msg"`
	}{}

	err := writer.Bind(&input)
	if err != nil {
		return writer.String(400, err.Error())
	}

	value, err := strconv.Atoi(*input.Msg)
	if err != nil {
		return writer.String(400, "Было введено не число или присутствуют пробелы в записи числа")
	}

	err = srv.uc.SetCount(value)
	if err != nil {
		return writer.String(500, err.Error())
	}

	if value > 0 {
		return writer.String(201, "Значение счётчика было изменено на +"+strconv.Itoa(value))
	} else {
		return writer.String(201, "Значение счётчика было изменено на "+strconv.Itoa(value))
	}
}
