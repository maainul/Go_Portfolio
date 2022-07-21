package storage

import (
	"database/sql"
	"time"
)

type Blog struct {
	ID        string         `db:"id"`
	Name      string         `db:"name"`
	CreatedAt time.Time      `db:"created_at,omitempty"`
	CreatedBy string         `db:"created_by"`
	UpdatedAt time.Time      `db:"updated_at,omitempty"`
	UpdatedBy string         `db:"updated_by,omitempty"`
	DeletedAt sql.NullTime   `db:"deleted_at,omitempty"`
	DeletedBy sql.NullString `db:"deleted_by,omitempty"`
}

type GetInTouch struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Message   string    `db:"message"`
	CreatedAt time.Time `db:"created_at,omitempty"`
	CreatedBy string    `db:"created_by"`
}
