package errors

import (
	"encoding/json"
	"log"
)

func New(text CustomError) error {
	return &errorString{text}
}

type errorString struct {
	s CustomError
}

func (e *errorString) Error() string {
	customMessage, err := json.Marshal(e.s)
	if err != nil {
		log.Fatal(err)
	}
	return string(customMessage)
}
