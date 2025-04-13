package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	filepath := "data.csv"
	file, err := os.Open(filepath)
	if err != nil {
		println("Error in opening file: ", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(records)
}
