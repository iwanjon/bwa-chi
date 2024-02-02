package handler

import (
	"bwastartupgochi/campaign"
	"bwastartupgochi/helper"
	mid "bwastartupgochi/middleware"
	"fmt"
	"io"
	"os"

	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type campaignHandler struct {
	service campaign.Service
}

type CampaignHandler interface {
	CreateCampaign(writer http.ResponseWriter, request *http.Request)
	UpdateCampaign(writer http.ResponseWriter, request *http.Request)
	GetCampaigns(writer http.ResponseWriter, request *http.Request)
	GetCampaign(writer http.ResponseWriter, request *http.Request)
	UploadCampaignImage(writer http.ResponseWriter, request *http.Request)
}

func NewCampaignHandler(s campaign.Service) *campaignHandler {
	return &campaignHandler{s}
}

func (h *campaignHandler) CreateCampaign(writer http.ResponseWriter, request *http.Request) {
	struct_ctx_intf := request.Context().Value(mid.Contectkey)
	struct_context, ok := struct_ctx_intf.(mid.StructUser)
	if !ok {
		helper.PanicIfError(errors.New("eror conver to struct context"), "error convert context val to struct")
	}

	// user_id := struct_context.CurrentUser.ID

	var input campaign.CreateCampaignInput

	helper.ReadFromRequestBody(request, &input)
	input.User = struct_context.CurrentUser
	campaing, err := h.service.CreateCampaign(request.Context(), input)
	helper.PanicIfError(err, " error in create campaign handler")
	campaignFormater := campaign.FormatCampaign(campaing)
	res := helper.APIResponse("success", http.StatusOK, "success", campaignFormater)
	helper.WriteToResponseBody(writer, res)
}

func (h *campaignHandler) UpdateCampaign(writer http.ResponseWriter, request *http.Request) {
	var campaignId campaign.GetCampaignDetailInput
	var inputparam campaign.CreateCampaignInput

	struct_ctx_intf := request.Context().Value(mid.Contectkey)
	struct_context, ok := struct_ctx_intf.(mid.StructUser)
	if !ok {
		helper.PanicIfError(errors.New("eror conver to struct context"), "error convert context val to struct")

	}
	id := chi.URLParam(request, "campaignid")
	log.Println("this is id of the campaign", id)

	intid, err := strconv.Atoi(id)
	helper.PanicIfError(err, " error in convert campaign id string to int")

	helper.ReadFromRequestBody(request, &inputparam)

	inputparam.User = struct_context.CurrentUser
	campaignId.ID = intid
	updatedCampaign, err := h.service.UpdateCampaign(request.Context(), campaignId, inputparam)
	helper.PanicIfError(err, " error in update campaign handler")

	camoaignFormater := campaign.FormatCampaign(updatedCampaign)

	res := helper.APIResponse("success", http.StatusOK, "success", camoaignFormater)
	helper.WriteToResponseBody(writer, res)

}

func (h *campaignHandler) GetCampaigns(writer http.ResponseWriter, request *http.Request) {
	userid := request.URL.Query().Get("user_id")
	user_id := 0
	if userid != "" {
		userIdInt, err := strconv.Atoi(userid)
		helper.PanicIfError(err, "error in get user id Get Campaign handler")
		user_id = userIdInt
	}
	campaigns, err := h.service.FindCampaigns(request.Context(), user_id)
	helper.PanicIfError(err, " error in gettting campaigns handler")
	campaignsFormated := campaign.FormatCampaigns(campaigns)
	res := helper.APIResponse("success", http.StatusOK, "success", campaignsFormated)
	helper.WriteToResponseBody(writer, res)

}
func (h *campaignHandler) GetCampaign(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "campaignid")
	log.Println("this is id of the campaign", id)

	intid, err := strconv.Atoi(id)
	helper.PanicIfError(err, " error in convert campaign id string to int")
	input := campaign.GetCampaignDetailInput{
		ID: intid,
	}
	current, err := h.service.GetCampaignById(request.Context(), input)
	helper.PanicIfError(err, " error in get campaign by id")

	campignFormated := campaign.FormatCampaignDetail(current)
	res := helper.APIResponse("success", http.StatusOK, "success", campignFormated)

	helper.WriteToResponseBody(writer, res)

}
func (h *campaignHandler) UploadCampaignImage(writer http.ResponseWriter, request *http.Request) {
	struct_ctx_intf := request.Context().Value(mid.Contectkey)
	struct_context, ok := struct_ctx_intf.(mid.StructUser)
	if !ok {
		helper.PanicIfError(errors.New("eror conver to struct context"), "error convert context val to struct")

	}

	file, fileheader, err := request.FormFile("file")
	helper.PanicIfError(err, "error in upload campaign image handler")

	isPrimary := request.PostFormValue("is_primary")
	campaignId := request.PostFormValue("campaign_id")
	campaignIdInt, err := strconv.Atoi(campaignId)
	helper.PanicIfError(err, "error in convert campaign id upload campaign handler")

	isPrimaryBool, err := strconv.ParseBool(isPrimary)
	helper.PanicIfError(err, "error in convert primary boolean upload campaign handler")

	input := campaign.CreateCampaignImageInput{
		CampaignID: campaignIdInt,
		IsPrimary:  isPrimaryBool,
		User:       struct_context.CurrentUser,
	}

	path := fmt.Sprintf("images/camapign-%d-%s", input.CampaignID, fileheader.Filename)
	fileDestination, err := os.Create(path)
	helper.PanicIfError(err, "erro in create path for upload campaign image")

	_, err = io.Copy(fileDestination, file)
	helper.PanicIfError(err, " error in copy file to destination upload campaign Handler")

	_, err = h.service.SaveCampaignImage(request.Context(), input, path)
	helper.PanicIfError(err, " error in save campaign handler")

	m := make(map[string]bool)
	m["is_uploaded"] = true

	res := helper.APIResponse("success", http.StatusOK, "success", m)
	helper.WriteToResponseBody(writer, res)
}

// func (h *campaignHandler) CreateCampaign(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
// 	fmt.Println(request.Context(), "cocococooc", request.Context().Value(middleware.Contectkey))
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

// 	var input campaign.CreateCampaignInput

// 	helper.ReadFromRequestBody(request, &input)
// 	input.User = currentUser
// 	newcampaign, err := h.service.CreateCampaign(request.Context(), input)
// 	helper.PanicIfError(err, " error in create campaign handler")
// 	campaignFormater := campaign.FormatCampaign(newcampaign)
// 	response := helper.APIResponse("success", http.StatusOK, "success create campaign", campaignFormater)
// 	helper.WriteToResponseBody(writer, response)
// }

// func (h *campaignHandler) UpdateCampaign(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

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

// 	var inputid campaign.GetCampaignDetailInput

// 	id := params.ByName("campaignid")
// 	if id == "" {
// 		exception.PanicIfNotFound(errors.New("not found error"), "error in get dynamic url")
// 	}
// 	idint, err := strconv.Atoi(params.ByName("campaignid"))
// 	exception.PanicIfNotFound(err, "error in get dynamic url int")
// 	inputid.ID = idint

// 	var input campaign.CreateCampaignInput
// 	input.User = currentUser
// 	helper.ReadFromRequestBody(request, &input)

// 	updatedCampaign, err := h.service.UpdateCampaign(request.Context(), inputid, input)
// 	helper.PanicIfError(err, "error in update campaign")
// 	campaignFormater := campaign.FormatCampaign(updatedCampaign)
// 	response := helper.APIResponse("success", http.StatusOK, "success create campaign", campaignFormater)
// 	helper.WriteToResponseBody(writer, response)

// }
// func (h *campaignHandler) GetCampaigns(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
// 	userid := request.URL.Query().Get("user_id")
// 	user_id := 0
// 	if userid != "" {
// 		userIdInt, err := strconv.Atoi(userid)
// 		helper.PanicIfError(err, "error in get user id Get Campaign handler")
// 		user_id = userIdInt
// 	}

// 	campaigns, err := h.service.FindCampaigns(request.Context(), user_id)
// 	helper.PanicIfError(err, "error in getin campaigns")
// 	campaignFormater := campaign.FormatCampaigns(campaigns)
// 	response := helper.APIResponse("success", http.StatusOK, "success create campaign", campaignFormater)
// 	helper.WriteToResponseBody(writer, response)
// }
// func (h *campaignHandler) GetCampaign(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

// 	campaignId := params.ByName("campaignid")
// 	campaignIdInt, err := strconv.Atoi(campaignId)
// 	exception.PanicIfNotFound(err, "no campaign id found in get campaign by id handler")
// 	campdinIdStruct := campaign.GetCampaignDetailInput{
// 		ID: campaignIdInt,
// 	}
// 	campaignById, err := h.service.GetCampaignById(request.Context(), campdinIdStruct)
// 	exception.PanicIfNotFound(err, "error in finding campaign by is handler")
// 	campaignformater := campaign.FormatCampaign(campaignById)
// 	response := helper.APIResponse("succes", http.StatusOK, "success", campaignformater)

// 	helper.WriteToResponseBody(writer, response)

// }
// func (h *campaignHandler) UploadCampaignImage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

// 	struct_context_interface := request.Context().Value(middleware.Contectkey)
// 	struct_context, ok := struct_context_interface.(middleware.StructContextKey)
// 	if !ok {
// 		helper.PanicIfError(errors.New("error in struct context"), "error in conver struct context in upload avatar handler")
// 	}

// 	user_User := struct_context.CurrentUser

// 	var input campaign.CreateCampaignImageInput
// 	file, fileheader, err := request.FormFile("file")
// 	helper.PanicIfError(err, "error in upload campaign image handler")

// 	isPrimary := request.PostFormValue("is_primary")
// 	campaignId := request.PostFormValue("campaign_id")
// 	campaignIdInt, err := strconv.Atoi(campaignId)
// 	helper.PanicIfError(err, "error in convert campaign id upload campaign handler")

// 	isPrimaryBool, err := strconv.ParseBool(isPrimary)
// 	helper.PanicIfError(err, "error in convert primary boolean upload campaign handler")

// 	input.User = user_User
// 	input.CampaignID = campaignIdInt
// 	input.IsPrimary = isPrimaryBool

// 	path := fmt.Sprintf("images/camapign-%d-%s", input.CampaignID, fileheader.Filename)
// 	fileDestination, err := os.Create(path)
// 	helper.PanicIfError(err, "erro in create path for upload campaign image")

// 	_, err = io.Copy(fileDestination, file)
// 	helper.PanicIfError(err, " error in copy file to destination upload campaign Handler")
// 	_, err = h.service.SaveCampaignImage(request.Context(), input, fileDestination.Name())
// 	helper.PanicIfError(err, "error in save campaign image uploadcampaign handler")
// 	data := make(map[string]bool)
// 	data["is_uploaded"] = true

// 	respnse := helper.APIResponse("upload image success", http.StatusOK, "success", data)
// 	helper.WriteToResponseBody(writer, respnse)
// }
