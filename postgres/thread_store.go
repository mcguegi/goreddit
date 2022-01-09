package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mcguegi/goreddit"
)

func NewThreadStore(db *sqlx.DB) *ThreadStore {
	return &ThreadStore{
		DB: db,
	}
}

type ThreadStore struct {
	*sqlx.DB
}

func (s *ThreadStore) Thread(id uuid.UUID) (goreddit.Thread, error) {
	var t goreddit.Thread
	if err := s.Get(&t, `SELECT * FROM threads WHERE id = $1`, id); err != nil {
		return goreddit.Thread{}, fmt.Errorf("error getting thread: %w", err)
	}
	return t, nil
}

func (s *ThreadStore) Threads() ([]goreddit.Thread, error) {
	var ts []goreddit.Thread
	if err := s.Select(&ts, `SELECT * FROM threads`); err != nil {
		return []goreddit.Thread{}, fmt.Errorf("error getting threads: %w", err)
	}
	return ts, nil
}

func (s *ThreadStore) CreateThread(t *goreddit.Thread) error {
	panic("not implemented") // TODO: Implement
}

func (s *ThreadStore) UpdateThread(t *goreddit.Thread) error {
	panic("not implemented") // TODO: Implement
}

func (s *ThreadStore) DeleteThread(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}
