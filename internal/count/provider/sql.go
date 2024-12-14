package provider

import (
	"database/sql"
	"errors"
)

func (p *Provider) SelectCount() (string, error) {
	var dbAnswer string

	row := p.conn.QueryRow("SELECT value FROM count")
	err := row.Scan(&dbAnswer) // Проверка на то, есть ли искомые данные в БД
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return dbAnswer, nil
}

func (p *Provider) UpdateCount(n int) error {
	_, err := p.conn.Exec("UPDATE count SET value = value + ($1)", n)
	if err != nil {
		return err
	}

	return nil
}
