package api

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

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
