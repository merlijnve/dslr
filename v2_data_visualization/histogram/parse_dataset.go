package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func parseFloatData(dataset [][]string, numericalFeatures []int) [][]float64 {
	floatData := make([][]float64, len(dataset[0]))

	for _, i := range numericalFeatures {
		var count float64 = 0
		for j := range dataset {
			if dataset[j][i] != "" {
				count++
				val, err := strconv.ParseFloat(dataset[j][i], 64)
				if err != nil {
					handleError(err, "Could not parse \""+dataset[j][i]+"\" to float")
				}
				floatData[i] = append(floatData[i], val)
			}
		}
		sort.Float64s(floatData[i])
	}
	return floatData
}

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
