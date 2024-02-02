package main

import (
	"bwastartupgochi/app"
	"bwastartupgochi/auth"
	"bwastartupgochi/campaign"
	"bwastartupgochi/exception"
	"bwastartupgochi/handler"
	"bwastartupgochi/helper"
	"bwastartupgochi/payment"
	"bwastartupgochi/transaction"
	"bwastartupgochi/user"
	"fmt"
	"net/http"

	mid "bwastartupgochi/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	defer db.Close()
	fmt.Println(db)
	validate := validator.New()

	userAuth := auth.Newjwtservice()
	// middleware := middleware.AuthChecker
	paymentService := payment.NewPaymentService()

	userrepository := user.NewRepository()
	userService := user.NewService(db, userrepository, validate)
	serHandler := handler.NewUserHandler(userService, userAuth)

	campaignRepository := campaign.NewRepository()
	campaignService := campaign.NewService(campaignRepository, db, validate)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	transactionRepo := transaction.NewRepository()
	transactionService := transaction.NewServiceTransaction(transactionRepo, campaignRepository, paymentService, db)
	transactionHAndler := handler.NewTransactionHandler(transactionService)

	tx, err := db.Begin()
	helper.PanicIfError(err, "erro in main tx")
	fmt.Println(tx)
	defer helper.CommitOrRollback(tx)
	// ctx := context.Background()

	// FindAll(ctx context.Context, tx *sql.DB) ([]Campaign, error)
	// FindByUserId(ctx context.Context, tx *sql.DB, userId int) ([]Campaign, error)
	// FindById(ctx context.Context, tx *sql.DB, campaignId int) (Campaign, error)
	// SaveCampaign(ctx context.Context, tx *sql.Tx, campaign Campaign) (Campaign, error)
	// UpdateCampaign(ctx context.Context, tx *sql.Tx, campaign Campaign) (Campaign, error)
	// SaveImage(ctx context.Context, tx *sql.Tx, campaignImage CampaignImage) (CampaignImage, error)
	// MarkAllImagesAsNonPrimary(ctx context.Context, tx *sql.Tx, campaignId int) (bool, error)

	// a, err := campaignRepository.FindAll(ctx, db)
	// campen := campaign.Campaign{
	// 	ID:               2,
	// 	UserID:           19,
	// 	Name:             "wessssssssssrrrrr",
	// 	ShortDescription: "rr",
	// 	Description:      "sssssss",
	// 	Perks:            "sssssssssss",
	// 	BackerCount:      123333330,
	// 	GoalAmount:       1110,
	// 	CurrentAmount:    10,
	// 	Slug:             "eeeeeeeeeeeeeeee",
	// }
	// cempenimage := campaign.CampaignImage{

	// 	CampaignID: 3,
	// 	FileName:   "qqqqqqqqqqqq",
	// 	IsPrimary:  0,
	// }
	// a, err := campaignRepository.SaveImage(ctx, tx, cempenimage)
	// fmt.Println(a, err, "alalal")
	// helper.CommitOrRollback(tx)
	// tesuse := user.User{
	// 	ID:             5,
	// 	Name:           "eeeeeemadang",
	// 	Occupation:     "madangeeee",
	// 	Email:          "emaile.email@gmail, .com",
	// 	PasswordHash:   "madange",
	// 	AvatarFileName: "madaneg.jpg",
	// 	Role:           "madang",
	// }
	// uu, err := userrepository.Save(ctx, tx, tesuse)
	// uu, err := userrepository.FindByEmail(ctx, tx, "email.email@gmail.com")FInd
	// uu, err := userrepository.FindById(ctx, tx, 9)
	// uu, err := userrepository.Update(ctx, tx, tesuse)
	// uu, err := userrepository.FindById(ctx, tx, 9)
	// fmt.Println(uu, "madang", err)

	// var input user.CheckEmailInput
	// input.Email = "email.email@gmail.com"
	// uu, err := userService.CheckEmailAvailable(ctx, input)
	// uu, err := userService.GetUserById(ctx, 3)
	// uu, err := userService.SaveAvatar(ctx, 13, "makan")
	// var input user.RegisterUser
	// input.Email = " beruang@gmail.com"
	// input.Name = "beruang"
	// input.Occupation = "beruang"
	// input.Password = "beruangberuangberuang"
	// uu, err := userService.RegisterUser(ctx, input)
	// fmt.Println(uu, "madang", err)
	// var inputt user.LoginInput
	// inputt.Email = "beruang@gmail.com"
	// inputt.Password = "beruangberuangberuang"
	// uuu, err := userService.LoginUser(ctx, inputt)
	// fmt.Println("madang", uuu, err)
	// tok, err := userAuth.GenerateJWTToken(3)
	// fmt.Println(tok, "rrrrrrrr", err)
	// // time.Sleep(time.Second * 10)
	// tt, err := userAuth.ValidateToken(tok)
	// fmt.Println(tt, "tttrrrrrrrr", err)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(exception.Recoverer)
	ruter := app.NewUserHandler(r, serHandler, userAuth, userService, mid.AuthChecker)
	campaignruter := app.NewCampaignHandler(ruter, campaignHandler, userAuth, userService, mid.AuthChecker)
	transactionRouter := app.NewTransactionHandler(campaignruter, transactionHAndler, userAuth, userService, mid.AuthChecker)

	http.ListenAndServe(":3000", transactionRouter)

}
