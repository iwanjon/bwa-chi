package user

import (
	"bwastartupgochi/exception"
	"bwastartupgochi/helper"
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(ctx context.Context, input RegisterUser) (User, error)
	LoginUser(ctx context.Context, input LoginInput) (User, error)
	CheckEmailAvailable(ctx context.Context, input CheckEmailInput) (bool, error)
	SaveAvatar(ctx context.Context, id int, filelocation string) (User, error)
	GetUserById(ctx context.Context, id int) (User, error)
}

type service struct {
	db       *sql.DB
	repo     Repository
	Validate *validator.Validate
}

func NewService(db *sql.DB, r Repository, v *validator.Validate) Service {
	return &service{db: db, repo: r, Validate: v}
}

func (s *service) RegisterUser(ctx context.Context, input RegisterUser) (User, error) {
	var userregister User

	tx, err := s.db.Begin()

	helper.PanicIfError(err, " erro in create tx service user register")

	// _, err = s.repo.FindByEmail(ctx, tx, input.Email)
	// if err == nil {
	// 	err = errors.New("email already registerd")
	// 	helper.PanicIfError(err, "error in find user by email register user service")
	// }

	userregister.Name = strings.TrimSpace(input.Name)
	userregister.Role = "user"
	userregister.Occupation = strings.TrimSpace(input.Occupation)
	userregister.Email = strings.TrimSpace(input.Email)
	bytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	helper.PanicIfError(err, "error in create hash password register user service")
	userregister.PasswordHash = string(bytes)
	fmt.Println(string(bytes))
	defer helper.CommitOrRollback(tx)
	user, err := s.repo.Save(ctx, tx, userregister)
	helper.PanicIfError(err, " error in cave user in register user service")

	return user, nil
}

func (s *service) LoginUser(ctx context.Context, input LoginInput) (User, error) {
	// var userLogin User
	tx, err := s.db.Begin()
	helper.PanicIfError(err, " erro in create tx service user")
	defer helper.CommitOrRollback(tx)
	userLogin, err := s.repo.FindByEmail(ctx, tx, input.Email)
	helper.PanicIfError(err, "erro in find by email user service")
	err = bcrypt.CompareHashAndPassword([]byte(userLogin.PasswordHash), []byte(input.Password))
	fmt.Println(err, input.Password, "madang", userLogin.PasswordHash, "madang")
	helper.PanicIfError(err, " error in validate password")
	return userLogin, nil
}

func (s *service) CheckEmailAvailable(ctx context.Context, input CheckEmailInput) (bool, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err, " erro in create tx service user")
	defer helper.CommitOrRollback(tx)
	_, err = s.repo.FindByEmail(ctx, tx, input.Email)
	// helper.PanicIfError(err, "erro in find by email user service")
	exception.PanicIfNotFound(err, " error in finding user by email")
	return false, nil
}

func (s *service) SaveAvatar(ctx context.Context, id int, filelocation string) (User, error) {
	tx, err := s.db.Begin()
	helper.PanicIfError(err, " erro in create tx service user")
	defer helper.CommitOrRollback(tx)
	user, err := s.repo.FindById(ctx, tx, id)
	helper.PanicIfError(err, " error in finding user by id suer respoitory save avatar")
	user.AvatarFileName = filelocation
	updatedUser, err := s.repo.Update(ctx, tx, user)
	helper.PanicIfError(err, " error in update suer respoitory save avatar")
	return updatedUser, nil
}

func (s *service) GetUserById(ctx context.Context, id int) (User, error) {

	tx, err := s.db.Begin()
	helper.PanicIfError(err, " erro in create tx service user")
	defer helper.CommitOrRollback(tx)
	user, err := s.repo.FindById(ctx, tx, id)
	helper.PanicIfError(err, " error in finding user by id suer respoitory save avatar")
	return user, nil
}
