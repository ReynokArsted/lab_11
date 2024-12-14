package provider

import (
	"database/sql"
	"errors"
)

func (p *Provider) SelectUser(log string, pass string) error {
	var dbAnswer string

	row := p.conn.QueryRow("SELECT * FROM auth WHERE login = $1 AND password = $2", log, pass)
	err := row.Scan(&dbAnswer) // Проверка на то, есть ли искомые данные в БД
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}

	return nil
}

func (p *Provider) InsertUser(log string, pass string) error {
	_, err := p.conn.Exec("INSERT INTO auth (login, password) VALUES ($1, $2)", log, pass)
	if err != nil {
		return err
	}

	return nil
}
