package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

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

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./compare results [training_file.csv] [predictions_file.csv]")
		os.Exit(1)
	}
	training_file := readCsvFile(os.Args[1])[1:]
	prediction_file := readCsvFile(os.Args[2])[1:]
	var correct = 0
	var incorrect = 0
	for i := range training_file {
		if training_file[i][1] == prediction_file[i][1] {
			correct++
		} else {
			incorrect++
		}
	}
	fmt.Println("Correct: ", correct)
	fmt.Println("Incorrect: ", incorrect)
	fmt.Println("Accuracy: ", float64(correct)/float64(correct+incorrect)*100, "%")
}
