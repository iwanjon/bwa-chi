package handler

import (
	"bwastartupgochi/helper"
	mid "bwastartupgochi/middleware"
	"bwastartupgochi/transaction"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type transactionHandler struct {
	Service transaction.Service
}

type TransactionHandler interface {
	CreateTransaction(writer http.ResponseWriter, request *http.Request)
	GetUserTransactions(writer http.ResponseWriter, request *http.Request)
	GetCampaignTransactions(writer http.ResponseWriter, request *http.Request)
	GetNotif(writer http.ResponseWriter, request *http.Request)
}

func NewTransactionHandler(s transaction.Service) *transactionHandler {
	return &transactionHandler{
		Service: s,
	}
}

// func (h *transactionHandler) CreateTransaction(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

// 	struct_context_interface := request.Context().Value(middleware.Contectkey)
// 	struct_context, ok := struct_context_interface.(middleware.StructContextKey)
// 	if !ok {
// 		helper.PanicIfError(errors.New("error in struct context"), "error in conver struct context in create campaign handler")
// 	}

// 	currentUser := struct_context.CurrentUser

// 	var input transaction.CreateTransactionInput
// 	helper.ReadFromRequestBody(request, &input)
// 	input.User = currentUser
// 	newTrans, err := h.Service.CreateTransaction(request.Context(), input)
// 	helper.PanicIfError(err, "error in creaye handler transaction")
// 	transFormater := transaction.FormatTransaction(newTrans)
// 	response := helper.APIResponse("success", http.StatusOK, "success create campaign", transFormater)
// 	helper.WriteToResponseBody(writer, response)
// }

func (h *transactionHandler) CreateTransaction(writer http.ResponseWriter, request *http.Request) {

	struct_ctx_intf := request.Context().Value(mid.Contectkey)
	struct_context, ok := struct_ctx_intf.(mid.StructUser)
	if !ok {
		helper.PanicIfError(errors.New("eror conver to struct context"), "error convert context val to struct")

	}

	currentUser := struct_context.CurrentUser

	var input transaction.CreateTransactionInput
	helper.ReadFromRequestBody(request, &input)
	input.User = currentUser
	newTrans, err := h.Service.CreateTransaction(request.Context(), input)
	helper.PanicIfError(err, "error in creaye handler transaction")
	transFormater := transaction.FormatTransaction(newTrans)
	response := helper.APIResponse("success", http.StatusOK, "success create campaign", transFormater)
	helper.WriteToResponseBody(writer, response)
}

// func (h *transactionHandler) GetUserTransactions(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
// 	struct_context_interface := request.Context().Value(middleware.Contectkey)
// 	struct_context, ok := struct_context_interface.(middleware.StructContextKey)
// 	if !ok {
// 		helper.PanicIfError(errors.New("error in struct context"), "error in conver struct context in create campaign handler")
// 	}

// 	currentUser := struct_context.CurrentUser
// 	user_id := currentUser.ID
// 	transactions, err := h.Service.GetTransactionByUserID(request.Context(), user_id)
// 	helper.PanicIfError(err, "error in get transaction by user id handler ")
// 	transFormater := transaction.FormatUserTransactions(transactions)
// 	fmt.Println(transactions[0], "fffffffffffffffffff")
// 	fmt.Println(transFormater[0], "fffffffffffff")
// 	response := helper.APIResponse("success", http.StatusOK, "success get by user id campaign", transFormater)
// 	helper.WriteToResponseBody(writer, response)
// }

func (h *transactionHandler) GetUserTransactions(writer http.ResponseWriter, request *http.Request) {
	struct_ctx_intf := request.Context().Value(mid.Contectkey)
	struct_context, ok := struct_ctx_intf.(mid.StructUser)
	if !ok {
		helper.PanicIfError(errors.New("eror conver to struct context"), "error convert context val to struct")

	}
	currentUser := struct_context.CurrentUser
	user_id := currentUser.ID
	transactions, err := h.Service.GetTransactionByUserID(request.Context(), user_id)
	helper.PanicIfError(err, "error in get transaction by user id handler ")
	transFormater := transaction.FormatUserTransactions(transactions)
	fmt.Println(transactions[0], "fffffffffffffffffff")
	fmt.Println(transFormater[0], "fffffffffffff")
	response := helper.APIResponse("success", http.StatusOK, "success get by user id campaign", transFormater)
	helper.WriteToResponseBody(writer, response)
}

// func (h *transactionHandler) GetCampaignTransactions(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

// 	var input transaction.GetCampaignTransactionsInput

// 	campaign_id := params.ByName("campaignid")
// 	int_campaign_id, err := strconv.Atoi(campaign_id)
// 	helper.PanicIfError(err, "error convert to int campaign id handler transaction")
// 	struct_context_interface := request.Context().Value(middleware.Contectkey)
// 	struct_context, ok := struct_context_interface.(middleware.StructContextKey)
// 	if !ok {
// 		helper.PanicIfError(errors.New("error in struct context"), "error in conver struct context in create campaign handler")
// 	}

// 	// userid := r.Context().Value(middleware.Contectkey)
// 	// fmt.Println(userid, "user id")
// 	// user_id, ok := userid.(int)

// 	// if !ok {
// 	// 	helper.PanicIfError(errors.New("error in convert user id"), "error in conver user id in upload avatar handler")
// 	// }

// 	currentUser := struct_context.CurrentUser
// 	input.User = currentUser
// 	input.ID = int_campaign_id

// 	trans, err := h.Service.GetTransactionByCampaignID(request.Context(), input)
// 	helper.PanicIfError(err, " errror in get trans by campaign id handler transaction")
// 	transFormater := transaction.FormatCampaignTransactions(trans)
// 	fmt.Println(trans, "dddddd", transFormater)
// 	response := helper.APIResponse("success", http.StatusOK, "success get campaigns", transFormater)
// 	helper.WriteToResponseBody(writer, response)

// }

func (h *transactionHandler) GetCampaignTransactions(writer http.ResponseWriter, request *http.Request) {

	var input transaction.GetCampaignTransactionsInput
	campaign_id := chi.URLParam(request, "campaignid")
	log.Println("this is id of the campaign", campaign_id)
	// campaign_id := params.ByName("campaignid")

	int_campaign_id, err := strconv.Atoi(campaign_id)
	helper.PanicIfError(err, "error convert to int campaign id handler transaction")

	struct_ctx_intf := request.Context().Value(mid.Contectkey)
	struct_context, ok := struct_ctx_intf.(mid.StructUser)
	if !ok {
		helper.PanicIfError(errors.New("eror conver to struct context"), "error convert context val to struct")

	}

	// userid := r.Context().Value(middleware.Contectkey)
	// fmt.Println(userid, "user id")
	// user_id, ok := userid.(int)

	// if !ok {
	// 	helper.PanicIfError(errors.New("error in convert user id"), "error in conver user id in upload avatar handler")
	// }

	currentUser := struct_context.CurrentUser
	input.User = currentUser
	input.ID = int_campaign_id

	trans, err := h.Service.GetTransactionByCampaignID(request.Context(), input)
	helper.PanicIfError(err, " errror in get trans by campaign id handler transaction")
	transFormater := transaction.FormatCampaignTransactions(trans)
	fmt.Println(trans, "dddddd", transFormater)
	response := helper.APIResponse("success", http.StatusOK, "success get campaigns", transFormater)
	helper.WriteToResponseBody(writer, response)

}

// // func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
// // 	var input transaction.GetCampaignTransactionsInput

// // 	err := c.ShouldBindUri(&input)
// // 	if err != nil {
// // 		errors := helper.FormatValidationError(err)
// // 		errormessage := gin.H{"errors": errors}
// // 		responseuser := helper.APIResponse("error get campaign transaction", http.StatusBadRequest, "error", errormessage)
// // 		c.JSON(http.StatusBadRequest, responseuser)
// // 		return
// // 	}
// // 	currentUser := c.MustGet("curresntuser").(user.User)
// // 	input.User = currentUser

// // 	transact, err := h.service.GetTransactionByCampaignID(input)
// // 	if err != nil {

// // 		responseuser := helper.APIResponse("error get campaign transaction", http.StatusBadRequest, "error", nil)
// // 		c.JSON(http.StatusBadRequest, responseuser)
// // 		return
// // 	}
// // 	responseuser := helper.APIResponse("OK", http.StatusOK, "OK", transaction.FormatCampaignTransactions(transact))
// // 	c.JSON(http.StatusOK, responseuser)

// // }

// func (h *transactionHandler) GetNotif(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
// 	var input transaction.TransactionNotificationInput
// 	helper.ReadFromRequestBody(request, &input)
// 	err := h.Service.ProcessPayment(request.Context(), input)
// 	helper.PanicIfError(err, "error in process payment, handler transaction")
// 	response := helper.APIResponse("success", http.StatusOK, "success create campaign", input)
// 	helper.WriteToResponseBody(writer, response)
// }

func (h *transactionHandler) GetNotif(writer http.ResponseWriter, request *http.Request) {
	var input transaction.TransactionNotificationInput
	helper.ReadFromRequestBody(request, &input)
	err := h.Service.ProcessPayment(request.Context(), input)
	helper.PanicIfError(err, "error in process payment, handler transaction")
	response := helper.APIResponse("success", http.StatusOK, "success create campaign", input)
	helper.WriteToResponseBody(writer, response)
}

// // func (h *transactionHandler) GetNotif(c *gin.Context) {
// // 	var input transaction.TransactionNotificationInput

// // 	err := c.ShouldBindJSON(&input)
// // 	if err != nil {
// // 		errors := helper.FormatValidationError(err)
// // 		errormessage := gin.H{"errors": errors}
// // 		responseuser := helper.APIResponse("error procedd transaction notification", http.StatusBadRequest, "error", errormessage)
// // 		c.JSON(http.StatusBadRequest, responseuser)
// // 		return
// // 	}

// // 	err = h.service.ProcessPayment(input)

// // 	if err != nil {

// // 		responseuser := helper.APIResponse("error procedd transaction notification", http.StatusBadRequest, "error", nil)
// // 		c.JSON(http.StatusBadRequest, responseuser)
// // 		return
// // 	}

// // 	responseuser := helper.APIResponse("OK", http.StatusOK, "OK", input)
// // 	c.JSON(http.StatusOK, responseuser)

// // }
