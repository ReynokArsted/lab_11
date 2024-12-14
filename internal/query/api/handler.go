package api

import (
	"github.com/labstack/echo/v4"
)

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
