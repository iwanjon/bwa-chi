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
