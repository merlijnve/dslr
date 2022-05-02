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
		numerical := true
		for _, row := range dataset[1:] {
			if row[i] != "" {
				_, err := strconv.ParseFloat(row[i], 64)
				if err != nil {
					numerical = false
				}
			}
		}
		if numerical == true {
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
		handleError(err, "Error: could not read file \""+os.Args[1]+"\"")
		defer file.Close()
	} else {
		fmt.Println("Use ./histogram [dataset filename]")
		os.Exit(0)
	}

	csv := csv.NewReader(file)
	records, err := csv.ReadAll()

	return records
}
