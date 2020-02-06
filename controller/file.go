package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"trip-planer/model"
)

func PostFileName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Println("Encountered POST method of PostTripDetails")
			// r.ParseForm()
			// fileName := r.PostForm.Get("file_name")
			fileName := FileName{}
			json.NewDecoder(r.Body).Decode(&fileName)
			// fmt.Printf("file name: %s", fileName.FileName)
			model.BulkInsert(fileName.FileName, fileName.Extension)

			json.NewEncoder(w).Encode(struct {
				Status string `json:"status"`
			}{"file uploaded to database"})
		}
	}
}

type FileName struct {
	FileName  string `json:"file_name"`
	Extension string `json:"extension"`
}
