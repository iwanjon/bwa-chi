package exception

import "log"

type NotOwnerError struct {
	Error string
}

func NewNotOwnerError(error string) NotOwnerError {
	return NotOwnerError{Error: error}
}

func PanicIfNotOwner(err error, s string) {
	if err != nil {
		log.Println(s)
		log.Println(err)
		panic(NewNotOwnerError(err.Error()))
	}
}
