package app

import (
	"bwastartupgochi/auth"
	"bwastartupgochi/handler"
	"bwastartupgochi/middleware"
	"bwastartupgochi/user"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

// func NewUserRouter(router *httprouter.Router, userHandler handler.UserHandler, directory http.Dir, jwtservice auth.Service, userservice user.Service, authmiddleware func(jwtservice auth.Service, userservice user.Service, h httprouter.Handle) httprouter.Handle) *httprouter.Router {
// 	// router := httprouter.New()
// 	router.ServeFiles("/files/*filepath", directory)
// 	router.POST("/api/v1/users", userHandler.RegisterUser)
// 	router.POST("/api/v1/sessions", userHandler.LoginUser)
// 	// router.POST("/api/v1/checkemail", middleware.EmailChecker(userHandler.CheckEmail))
// 	router.POST("/api/v1/email_checkers", userHandler.CheckEmail)
// 	router.POST("/api/v1/avatars", authmiddleware(jwtservice, userservice, userHandler.UploadAvatar))
// 	router.POST("/api/v1/users/fetch", authmiddleware(jwtservice, userservice, userHandler.FetchUser))
// 	router.PanicHandler = exception.ErrorHandler

// 	return router
// }

// func NewCampaignHandler(router *httprouter.Router, campaignHandler handler.CampaignHandler, jwtservice auth.Service, userservice user.Service, authmiddleware func(jwtservice auth.Service, userservice user.Service, h httprouter.Handle) httprouter.Handle) *httprouter.Router {

// 	router.GET("/api/v1/campaigns", campaignHandler.GetCampaigns)
// 	router.POST("/api/v1/campaigns", authmiddleware(jwtservice, userservice, campaignHandler.CreateCampaign))
// 	router.PUT("/api/v1/campaigns/:campaignid", authmiddleware(jwtservice, userservice, campaignHandler.UpdateCampaign))
// 	router.GET("/api/v1/campaigns/:campaignid", campaignHandler.GetCampaign)
// 	router.POST("/api/v1/campaign-images", authmiddleware(jwtservice, userservice, campaignHandler.UploadCampaignImage))

// 	return router
// }

// func NewTransactionHandler(router *httprouter.Router, transactionHandler handler.TransactionHandler, jwtservice auth.Service, userservice user.Service, authmiddleware func(jwtservice auth.Service, userservice user.Service, h httprouter.Handle) httprouter.Handle) *httprouter.Router {

// 	router.GET("/api/v1/campaigns/:campaignid/transactions", authmiddleware(jwtservice, userservice, transactionHandler.GetCampaignTransactions))
// 	router.GET("/api/v1/transactions", authmiddleware(jwtservice, userservice, transactionHandler.GetUserTransactions))
// 	router.POST("/api/v1/transactions", authmiddleware(jwtservice, userservice, transactionHandler.CreateTransaction))
// 	router.POST("/api/v1/transactions/notification", transactionHandler.GetNotif)
// 	return router
// }
// director:= http.Dir("./images")

// fileServer := http.fileServer(directory)

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}
	fmt.Println("kakakasssssssk")

	if path != "/" && path[len(path)-1] != '/' {
		fmt.Println("pattt", path, len(path), path[len(path)-1])
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		fmt.Println("pattt", path, len(path), path[len(path)-1])
		path += "/"
		fmt.Println("pattt", path, len(path), path[len(path)-1])

	}
	path += "*"
	fmt.Println("kakakak")
	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fmt.Println(rctx, "ggg", pathPrefix, "ggg", fs, "lll", rctx.RoutePattern())
		fs.ServeHTTP(w, r)
		fmt.Println("after fs run ", fs)
		return
	})
}

func NewUserHandler(r *chi.Mux, h handler.UserHandler, jwtService auth.Service, userService user.Service, authModdleware func(jwtService auth.Service, userService user.Service, h http.HandlerFunc) http.HandlerFunc) *chi.Mux {
	// r.Handle()
	dire := http.Dir("./images")
	FileServer(r, "/static", dire)

	r.Post("/api/v1/users", h.RegisterUser)
	r.Post("/api/v1/sessions", h.LoginUser)
	r.Post("/api/v1/email_checkers", h.CheckEmail)
	r.Post("/api/v1/avatars", authModdleware(jwtService, userService, h.UploadAvatar))
	// r.Post("/api/v1/avatares", h.UploadAvatar)
	r.Post("/api/v1/user/fetch", h.FetchUser)
	return r
}

func NewCampaignHandler(r *chi.Mux, h handler.CampaignHandler, jwtService auth.Service, userService user.Service, authModdleware func(jwtService auth.Service, userService user.Service, h http.HandlerFunc) http.HandlerFunc) *chi.Mux {
	// r.Handle()
	// dire := http.Dir("./images")
	// FileServer(r, "/static", dire)

	// r.Post("/api/v1/users", h.RegisterUser)
	// r.Post("/api/v1/sessions", h.LoginUser)
	// r.Post("/api/v1/email_checkers", h.CheckEmail)
	// r.Post("/api/v1/avatars", authModdleware(jwtService, userService, h.UploadAvatar))
	// // r.Post("/api/v1/avatares", h.UploadAvatar)
	// r.Post("/api/v1/user/fetch", h.FetchUser)

	//// set middleware in chi way
	authmidchi := middleware.NewAuthMidStruct(jwtService, userService)

	r.Get("/api/v1/campaigns", h.GetCampaigns)
	r.Get("/api/v1/campaigns/{campaignid}", h.GetCampaign)

	//// standar way using middleware
	// r.Post("/api/v1/campaigns", authModdleware(jwtService, userService, h.CreateCampaign))

	//// Another Chi way using middleware
	r.With(authmidchi.AuthMid).Post("/api/v1/campaigns", h.CreateCampaign)

	//// standar way using middleware
	// r.Put("/api/v1/campaigns/{campaignid}", authModdleware(jwtService, userService, h.UpdateCampaign))
	// r.Post("/api/v1/campaign-images", authModdleware(jwtService, userService, h.UploadCampaignImage))

	//// Another Chi way using middleware
	r.Group(func(r chi.Router) {
		r.Use(authmidchi.AuthMid)
		r.Put("/api/v1/campaigns/{campaignid}", h.UpdateCampaign)
		r.Post("/api/v1/campaign-images", h.UploadCampaignImage)
	})

	return r
}

func NewTransactionHandler(r *chi.Mux, transactionHandler handler.TransactionHandler, jwtservice auth.Service, userservice user.Service, authmiddleware func(jwtService auth.Service, userService user.Service, h http.HandlerFunc) http.HandlerFunc) *chi.Mux {

	// func NewTransactionHandler(router *httprouter.Router, transactionHandler handler.TransactionHandler, jwtservice auth.Service, userservice user.Service, authmiddleware func(jwtservice auth.Service, userservice user.Service, h httprouter.Handle) httprouter.Handle) *httprouter.Router {

	r.Get("/api/v1/campaigns/{campaignid}/transactions", authmiddleware(jwtservice, userservice, transactionHandler.GetCampaignTransactions))
	r.Get("/api/v1/transactions", authmiddleware(jwtservice, userservice, transactionHandler.GetUserTransactions))
	r.Post("/api/v1/transactions", authmiddleware(jwtservice, userservice, transactionHandler.CreateTransaction))
	r.Post("/api/v1/transactions/notification", transactionHandler.GetNotif)
	return r
}
