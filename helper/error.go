package helper

import (
	"log"
)

func PanicIfError(err error, s string) {
	log.Println(err)
	if err != nil {
		log.Println(s)
		panic(err)
	}
}
