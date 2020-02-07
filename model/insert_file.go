package model

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func BulkInsert(fileName, extension string) bool {
	tableName := getTableName(fileName)
	filePath := "/Users/narayananv/Desktop/go_data_files/" + fileName + "." + extension

	fmt.Printf("table name: %s\nfile path:%s\n", tableName, filePath)
	// mysql.RegisterLocalFile(filePath)
	// result, err := connect.Exec("LOAD DATA LOCAL INFILE '" + filePath + "' INTO TABLE " + tableName)
	// if err != nil {
	// 	log.Fatal(err.Error)
	// }
	// fmt.Println(result)
	// mysql.RegisterLocalFile()
	// insertQ, err := connect.Query()
	return true
}

func getTableName(fileName string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(fileName, "")
	return strings.ToLower(processedString)
}
