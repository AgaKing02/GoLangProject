package mysql

import (
	"awesomeProject/pkg/models"
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/pgxpool"
	"strconv"
	"time"
)

type SnippetModel struct {
	Pool *pgxpool.Pool
}

// This will insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	var id uint64
	interval, err := strconv.Atoi(expires)
	row := m.Pool.QueryRow(context.Background(), "INSERT INTO snippets (title,content,created,expires) VALUES ($1,$2,$3,$4) RETURNING id", title, content, time.Now(), time.Now().AddDate(0, 0, interval))
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	s := &models.Snippet{}
	err := m.Pool.QueryRow(context.Background(), "SELECT id, title, content, created, expires FROM snippets where id=$1 AND expires > now()", id).
		Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	rows, err := m.Pool.Query(context.Background(), "SELECT id, title, content, created, expires FROM snippets WHERE expires > now() ORDER BY created DESC LIMIT 10")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var snippets []*models.Snippet

	for rows.Next() {
		s := &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}
