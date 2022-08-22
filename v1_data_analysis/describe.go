package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(0)
	}
}

func IdentifyNumericalFeatures(dataset [][]string) []int {
	numericalFeatures := make([]int, 0)

	for i := 0; i < len(dataset[0]); i++ {
		emptyFeatures := 0
		numerical := true
		for _, row := range dataset[1:] {
			if row[i] != "" {
				_, err := strconv.ParseFloat(row[i], 64)
				if err != nil {
					numerical = false
				}
			} else {
				emptyFeatures++
			}
		}
		if numerical && emptyFeatures < len(dataset)-1 {
			numericalFeatures = append(numericalFeatures, i)
		}
	}
	return numericalFeatures
}

func ReadDataset() [][]string {
	var file *os.File
	var err error

	if len(os.Args) > 1 {
		file, err = os.Open(os.Args[1])
		handleError(err, "Error: could not open file \""+os.Args[1]+"\"")
		defer file.Close()
	} else {
		fmt.Println("Use ./describe [dataset filename]")
		os.Exit(0)
	}

	csv := csv.NewReader(file)
	records, err := csv.ReadAll()
	handleError(err, "Error: could not read file \""+os.Args[1]+"\""+"\n(most likely not properly formatted as a csv)")
	return records
}

func Describe() {
	dataset := ReadDataset()
	if len(dataset) > 1 {
		numericalFeatures := IdentifyNumericalFeatures(dataset)
		displayInformation(dataset, numericalFeatures)
	} else {
		handleError(errors.New("empty dataset"), "Error: dataset is empty")
	}
}

func main() {
	Describe()
}
