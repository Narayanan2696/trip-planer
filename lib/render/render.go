package render

import (
	"encoding/json"
	"net/http"
	"strings"
	"trip-planer/lib/errors"
)

func JSON(w http.ResponseWriter, err error, output interface{}) {
	if err != nil {
		var errorResponse errors.CustomError
		json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		json.NewEncoder(w).Encode(output)
	}
}
