package api

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TokenData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func (srv *Server) login(writer echo.Context) error {
	input := struct {
		Login    *string `json:"login"`
		Password *string `json:"password"`
	}{}

	err := writer.Bind(&input)
	if err != nil {
		return writer.String(400, err.Error())
	}

	err = srv.uc.FindUser(*input.Login, *input.Password)
	if err != nil {
		return echo.ErrUnauthorized
	}

	err = srv.uc.AddUser(*input.Login, *input.Password)
	if err != nil {
		return writer.String(500, err.Error())
	}

	claims := &TokenData{
		*input.Login,
		*input.Password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 60)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return writer.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
