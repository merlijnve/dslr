package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func findIndexOfFeature(dataset [][]string, feature string) int {
	for i := range dataset[0] {
		if dataset[0][i] == feature {
			return i
		}
	}
	return -1
}

func getDataPair(dataset [][]string, t Classifier) [][]float64 {
	dataPair := make([][]float64, 0)
	i0 := findIndexOfFeature(dataset, t.Feature0)
	i1 := findIndexOfFeature(dataset, t.Feature1)
	isClass := 0

	for i := range dataset {
		if i != 0 && dataset[i][i0] != "" && dataset[i][i1] != "" {
			val0, err := strconv.ParseFloat(dataset[i][i0], 64)
			handleError(err, "Error: could not parse "+dataset[i][i0])

			val1, err2 := strconv.ParseFloat(dataset[i][i1], 64)
			handleError(err2, "Error: could not parse "+dataset[i][i1])

			if dataset[i][1] == t.House {
				isClass = 1
			}

			dataPair = append(dataPair, []float64{val0, val1, float64(isClass)})
		}
	}
	return dataPair
}

func readDataset() [][]string {
	var file *os.File
	var err error

	if len(os.Args) > 1 {
		file, err = os.Open(os.Args[1])
		handleError(err, "Error: could not open file \""+os.Args[1]+"\"")
		defer file.Close()
	} else {
		fmt.Println("Use ./logreg_train [dataset filename]")
		os.Exit(0)
	}

	csv := csv.NewReader(file)
	records, err := csv.ReadAll()
	handleError(err, "Error: could not parse csv")
	return records
}
