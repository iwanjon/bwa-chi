package exception

import "log"

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

func PanicIfNotFound(err error, s string) {
	if err != nil {
		log.Println(s)
		log.Println(err)
		panic(NewNotFoundError(err.Error()))
	}
}
