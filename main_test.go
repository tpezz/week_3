package main

import (
	"os"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// Create a temporary test input file
	Input := "test_input.csv"
	Output := "test_output.jsonl"

	//sample data
	first_line_data := "value,income,age,rooms,bedrooms,pop,hh\n452600,8.3252,41,880,129,322,126\n"
	err := os.WriteFile(Input, []byte(first_line_data), 0644)

	//make sure to close files when done
	defer os.Remove(Input)
	defer os.Remove(Output)

	//run program
	os.Args = []string{"cmd", Input, Output}
	main()

	//get the output data
	OutputData, err := os.ReadFile(Output)
	if err != nil {
		t.Fatalf("output error: %v", err)
	}

	Expected := "{\"value\":452600,\"income\":8.3252,\"age\":41,\"rooms\":880,\"bedrooms\":129,\"pop\":322,\"hh\":126}\n"

	if string(OutputData) != Expected {
		t.Errorf("Output content does not match expected content")
	}

}
