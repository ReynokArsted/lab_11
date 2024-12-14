package provider

import (
	"database/sql"
	"errors"
)

func (p *Provider) AddName(name string) error {
	_, err := p.conn.Exec("UPDATE query SET name = ($1)", name)
	if err != nil {
		return err
	}

	return nil
}

func (p *Provider) SelectName() (string, error) {
	var answer string

	row := p.conn.QueryRow("SELECT name FROM query LIMIT 1")
	err := row.Scan(&answer)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	return answer, nil
}
