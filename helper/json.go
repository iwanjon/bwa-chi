package helper

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err, "error read from body")
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	log.Println(writer, "this is witer", response)
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)

	err := encoder.Encode(response)
	PanicIfError(err, "error write response")
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type response struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func APIResponse(message string, code int, status string, data interface{}) response {

	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	log.Println(meta, "metaresponse")
	jsonresponse := response{
		Meta: meta,
		Data: data,
	}
	log.Println(jsonresponse, "jsonresponse")
	return jsonresponse
}

func FormatValidationError(err error) []string {

	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
