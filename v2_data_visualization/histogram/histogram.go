package main

import (
	"fmt"
	"os"
)

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(0)
	}
}

func main() {
	dataset := readDataset()
	numericalFeatures := identifyNumericalFeatures(dataset)
	houses := make([]string, 0)

	for i := 1; i < len(dataset); i++ {
		houses = append(houses, dataset[i][1])
	}

	err := os.Mkdir("tmp", 0755)
	handleError(err, "Error: could not create folder tmp")
	for _, i := range numericalFeatures {
		createHistogram(dataset[0][i], i, dataset)
	}
	combineFeatureImages(dataset, numericalFeatures)
}
