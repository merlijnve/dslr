package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// opens a csv file, reads it, and returns a slice of strings
func readCsvFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return rawCSVdata
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println("Error: " + msg)
		os.Exit(1)
	}
}

func main() {
	// if the length of os.Args is not 3, then print an error and exit
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./compare results [training_file.csv] [predictions_file.csv]")
		os.Exit(1)
	}
	training_file := readCsvFile(os.Args[1])[1:]
	prediction_file := readCsvFile(os.Args[2])[1:]
	// loop through the csv file and print the results
	var correct = 0
	var incorrect = 0
	for i := range training_file {
		if training_file[i][1] == prediction_file[i][1] {
			correct++
		} else {
			incorrect++
		}
		// fmt.Println(training_file[i][1], prediction_file[i][1], training_file[i][1] == prediction_file[i][1])
	}
	fmt.Println("Correct: ", correct)
	fmt.Println("Incorrect: ", incorrect)
	// print the accuracy in percentage
	fmt.Println("Accuracy: ", float64(correct)/float64(correct+incorrect)*100, "%")
}
