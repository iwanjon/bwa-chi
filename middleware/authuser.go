package middleware

import (
	"bwastartupgochi/auth"
	"bwastartupgochi/helper"
	"bwastartupgochi/user"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type StructUser struct {
	CurrentUser user.User
}

type ContextKeyType string

var Contectkey ContextKeyType = "userKey"

func AuthChecker(jwtService auth.Service, userService user.Service, h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		authVal := r.Header.Get("Authorization")
		fmt.Println(authVal, "authVal")
		if !strings.Contains(authVal, "Bearer") {
			helper.PanicIfError(errors.New("error check bearer"), " error in check bearer")
		}

		arrayToken := strings.Split(authVal, " ")
		fmt.Println(arrayToken, "arrayToken")
		if len(arrayToken) != 2 {
			helper.PanicIfError(errors.New("error in spliting token "), "error split token")
		}

		jwtToken := arrayToken[1]
		tok, err := jwtService.ValidateToken(jwtToken)
		helper.PanicIfError(err, " error in validate token")

		claim, ok := tok.Claims.(jwt.MapClaims)
		fmt.Println(claim, "claim", tok, "took", ok)
		if !ok || !tok.Valid {
			helper.PanicIfError(errors.New("erro validate token"), "error claim token")
		}
		user_id := claim["jti"]
		stringuser, ok := user_id.(string)
		fmt.Println(stringuser, "floatuser", ok, user_id)
		if !ok {
			helper.PanicIfError(errors.New("error in conver to int auth checker"), "error conver user id to int auth checker")
		}
		intid, err := strconv.Atoi(stringuser)
		helper.PanicIfError(err, " error in id string to int")
		intuser := intid
		fmt.Println(intuser)

		LoginUser, err := userService.GetUserById(r.Context(), intuser)
		helper.PanicIfError(err, " error in getting loigin user middleware")
		context_value := StructUser{
			CurrentUser: LoginUser,
		}
		ctx := context.WithValue(r.Context(), Contectkey, context_value)
		h(w, r.WithContext(ctx))
	}

}

type AuthMidStruct struct {
	jwtService  auth.Service
	userService user.Service
}

func NewAuthMidStruct(jwtService auth.Service, userService user.Service) *AuthMidStruct {
	return &AuthMidStruct{jwtService, userService}
}
func (a *AuthMidStruct) AuthMid(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// defer func() {
		// 	rvr := recover()

		// 	if rvr != nil && rvr != http.ErrAbortHandler && notFoundError(w, r, rvr) {
		// 		return
		// 	}

		// 	if rvr != nil && rvr != http.ErrAbortHandler && validationErrors(w, r, rvr) {
		// 		return
		// 	}

		// 	if rvr != nil && rvr != http.ErrAbortHandler && notOwnerError(w, r, rvr) {
		// 		return
		// 	}

		// 	if rvr != nil && rvr != http.ErrAbortHandler {
		// 		internalServerError(w, r, rvr)
		// 		return

		// 	}

		// }()

		authVal := r.Header.Get("Authorization")
		fmt.Println(authVal, "authVal")
		if !strings.Contains(authVal, "Bearer") {
			helper.PanicIfError(errors.New("error check bearer"), " error in check bearer")
		}

		arrayToken := strings.Split(authVal, " ")
		fmt.Println(arrayToken, "arrayToken")
		if len(arrayToken) != 2 {
			helper.PanicIfError(errors.New("error in spliting token "), "error split token")
		}

		jwtToken := arrayToken[1]
		tok, err := a.jwtService.ValidateToken(jwtToken)
		helper.PanicIfError(err, " error in validate token")

		claim, ok := tok.Claims.(jwt.MapClaims)
		fmt.Println(claim, "claim", tok, "took", ok)
		if !ok || !tok.Valid {
			helper.PanicIfError(errors.New("erro validate token"), "error claim token")
		}
		user_id := claim["jti"]
		stringuser, ok := user_id.(string)
		fmt.Println(stringuser, "floatuser", ok, user_id)
		if !ok {
			helper.PanicIfError(errors.New("error in conver to int auth checker"), "error conver user id to int auth checker")
		}
		intid, err := strconv.Atoi(stringuser)
		helper.PanicIfError(err, " error in id string to int")
		intuser := intid
		fmt.Println(intuser)

		LoginUser, err := a.userService.GetUserById(r.Context(), intuser)
		helper.PanicIfError(err, " error in getting loigin user middleware")
		context_value := StructUser{
			CurrentUser: LoginUser,
		}
		ctx := context.WithValue(r.Context(), Contectkey, context_value)
		// h(w, r.WithContext(ctx))

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
