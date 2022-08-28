package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func identifyNumericalFeatures(dataset [][]string) []int {
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

func readDataset() [][]string {
	var file *os.File
	var err error

	if len(os.Args) > 1 {
		file, err = os.Open(os.Args[1])
		handleError(err, "Error: could not open file \""+os.Args[1]+"\"")
		defer file.Close()
	} else {
		fmt.Println("Use ./pair [dataset filename]")
		os.Exit(0)
	}

	csv := csv.NewReader(file)
	records, err := csv.ReadAll()
	handleError(err, "Error: invalid csv")
	return records
}
