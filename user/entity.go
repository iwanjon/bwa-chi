package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	Token          sql.NullString
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// migrate -database "postgresql://postgres:a@localhost:5432/bwastartup" -dir db/migrations up
// migrate -database "postgresql://postgres:a@localhost:5432/bwastartup" -dir db/migrations down
// migrate  create -ext sql -dir db/migrations create_table_users
