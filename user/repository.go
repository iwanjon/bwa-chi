package user

import (
	"bwastartupgochi/exception"
	"bwastartupgochi/helper"
	"context"
	"database/sql"
)

type repository struct {
}

type Repository interface {
	Save(ctx context.Context, tx *sql.Tx, user User) (User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (User, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (User, error)
	Update(ctx context.Context, tx *sql.Tx, user User) (User, error)
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Save(ctx context.Context, tx *sql.Tx, user User) (User, error) {
	sql_state_1 := "insert into users(Name, Occupation, Email, Password_Hash, Avatar_File_Name, Role) values ($1,$2,$3,$4,$5,$6) returning id;"
	sate1, err := tx.PrepareContext(ctx, sql_state_1)
	helper.PanicIfError(err, " error in create staement 1 repo")
	defer sate1.Close()
	result := sate1.QueryRowContext(ctx, user.Name, user.Occupation, user.Email, user.PasswordHash, user.AvatarFileName, user.Role)
	helper.PanicIfError(err, " error in insert user repository save user")
	var id int
	err = result.Scan(&id)
	helper.PanicIfError(err, " erro in getting last insert id save user repo")
	user.ID = id

	// sql_state_2 := "insert into users(Name, Occupation, Email, Password_Hash, Avatar_File_Name, Role) values ($1,$2,$3,$4,$5,$6) returning id;"
	// sate2, err := tx.PrepareContext(ctx, sql_state_2)
	// helper.PanicIfError(err, " error in create staement 1 repo")
	// defer sate2.Close()
	// result2 := sate2.QueryRowContext(ctx, user.Name, user.Occupation, user.Email, user.PasswordHash, user.AvatarFileName, user.Role)
	// helper.PanicIfError(err, " error in insert user repository save user")
	// var idd int
	// err = result2.Scan(&idd)
	// helper.PanicIfError(err, " erro in getting last insert id save user repo")
	// user.ID = idd
	return user, nil
}

func (r *repository) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (User, error) {
	user := User{}
	sql_state_1 := " select users.id, users.name, users.occupation, users.email, users.avatar_file_name, users.role, users.token, users.password_hash from users where users.email = $1;"
	stetement, err := tx.PrepareContext(ctx, sql_state_1)
	helper.PanicIfError(err, "erro in cerate statement in find by email repository")
	defer stetement.Close()
	result := stetement.QueryRowContext(ctx, email)
	err = result.Scan(
		&user.ID,
		&user.Name,
		&user.Occupation,
		&user.Email,
		&user.AvatarFileName,
		&user.Role,
		&user.Token,
		&user.PasswordHash,
	)
	exception.PanicIfNotFound(err, " error in finding user by email reposiotry user ")

	return user, nil
}

func (r *repository) FindById(ctx context.Context, tx *sql.Tx, id int) (User, error) {
	user := User{}
	sql_state_1 := " select users.id, users.name, users.occupation, users.email, users.avatar_file_name, users.role, users.token from users where users.id = $1;"
	stetement, err := tx.PrepareContext(ctx, sql_state_1)
	helper.PanicIfError(err, "erro in cerate statement in find by id repository")
	defer stetement.Close()
	result := stetement.QueryRowContext(ctx, id)
	err = result.Scan(
		&user.ID,
		&user.Name,
		&user.Occupation,
		&user.Email,
		&user.AvatarFileName,
		&user.Role,
		&user.Token,
	)
	exception.PanicIfNotFound(err, " error in finding user by id reposiotry user ")

	return user, nil
}

func (r *repository) Update(ctx context.Context, tx *sql.Tx, user User) (User, error) {
	// user := User{}
	sql_state_1 := " update users set name = $1, occupation= $2, email = $3, avatar_file_name = $4, role =$5, token =$6 where id = $7 returning id;"
	stetement, err := tx.PrepareContext(ctx, sql_state_1)
	helper.PanicIfError(err, "erro in cerate statement in update user repository")
	defer stetement.Close()
	result := stetement.QueryRowContext(ctx, user.Name, user.Occupation, user.Email, user.AvatarFileName, user.Role, user.Token, user.ID)
	err = result.Scan(
		&user.ID,
	)
	exception.PanicIfNotFound(err, " error in update user reposiotry user ")

	return user, nil

}
