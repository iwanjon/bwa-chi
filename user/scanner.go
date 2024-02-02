package user

import (
	"database/sql"
)

func UserScanner(r *sql.Row, u *User) error {
	err := r.Scan(
		&u.ID,
		&u.Name,
		&u.Occupation,
		&u.Email,
		&u.PasswordHash,
		&u.AvatarFileName,
		&u.Role,
		&u.Token,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	return err
}
func UserScanners(r *sql.Rows, u *User) error {
	err := r.Scan(
		&u.ID,
		&u.Name,
		&u.Occupation,
		&u.Email,
		&u.PasswordHash,
		&u.AvatarFileName,
		&u.Role,
		&u.Token,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	return err
}
