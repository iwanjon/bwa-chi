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

//go install -tags 'postgres,mysql,mongodb' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
// migrate  create -ext sql -dir db/migrations create_table_users
// migrate -database "postgresql://postgres:a1@localhost:5432/bwastartup" -dir db/migrations up
// migrate -database "postgresql://postgres:a1@localhost:5432/bwastartup" -dir db/migrations down
