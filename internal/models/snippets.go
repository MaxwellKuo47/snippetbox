package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := "INSERT INTO snippets (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))"
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), err
}

func (m *SnippetModel) Get(id int) (data *Snippet, err error) {
	stmt := "SELECT * FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?"
	row := m.DB.QueryRow(stmt, id)
	data = &Snippet{}
	err = row.Scan(&data.ID, &data.Title, &data.Content, &data.Created, &data.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return
		}
	}
	return
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
