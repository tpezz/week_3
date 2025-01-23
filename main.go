package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// Create a data structure that matches the incoming CSV
type data_for_json struct {
	Value    float64 `json:"value"`
	Income   float64 `json:"income"`
	Age      int     `json:"age"`
	Rooms    int     `json:"rooms"`
	Bedrooms int     `json:"bedrooms"`
	Pop      int     `json:"pop"`
	HH       int     `json:"hh"`
}

func main() {
	//Exit if the wrong number of inputs are given
	if len(os.Args) != 3 {
		fmt.Println("Incorrect input. Please use <code_file.go. <input.csv> <output.jsonl>")
		os.Exit(1)
	}
	//Pull files and set as vars
	InputFile := os.Args[1]
	OutputFile := os.Args[2]

	//Ensure that input file is a csv
	if filepath.Ext(InputFile) != ".csv" {
		fmt.Println("Incorrect input. Input file must have .csv extension")
		os.Exit(1)
	}

	//Open the input CSV file
	CsvFile, err := os.Open(InputFile)
	if err != nil {
		fmt.Printf("An error occured when opening the CSV file: %v\n", err)
		os.Exit(1)
	}
	//remember to close the file when the main function exits
	defer CsvFile.Close()

	//read csv and confirm no error
	reader := csv.NewReader(CsvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("An error occured when reading the CSV file: %v\n", err)
	}

	//create the output file
	JsonFile, err := os.Create(OutputFile)
	if err != nil {
		fmt.Printf("An error occured when creating the output file: %v\n", err)
	}

	//remember to close the file when the main function exits
	defer JsonFile.Close()

	//skip first row which is the headers and iteract through each row
	for _, record := range records[1:] {
		//select the right datatype and store as var
		// use _ as we dont want to track errors
		value, _ := strconv.ParseFloat(record[0], 64)
		income, _ := strconv.ParseFloat(record[1], 64)
		age, _ := strconv.Atoi(record[2])
		rooms, _ := strconv.Atoi(record[3])
		bedrooms, _ := strconv.Atoi(record[4])
		pop, _ := strconv.Atoi(record[5])
		hh, _ := strconv.Atoi(record[6])

		//create data object using data_for_json
		FinalData := data_for_json{
			Value:    value,
			Income:   income,
			Age:      age,
			Rooms:    rooms,
			Bedrooms: bedrooms,
			Pop:      pop,
			HH:       hh,
		}

		//Convert to Json
		JsonData, err := json.Marshal((FinalData))
		if err != nil {
			fmt.Printf("Error marshalling %v:", err)
			continue
		}
		//Write data to Json file and only keep error. Do not need to  declare it
		//add in new lines to JSON file
		_, err = JsonFile.Write(append(JsonData, '\n'))
		if err != nil {
			fmt.Printf("Error writing to output file %v\n", err)
			os.Exit(1)
		}

	}

	fmt.Println("File has been created")
	fmt.Println(OutputFile)

}
