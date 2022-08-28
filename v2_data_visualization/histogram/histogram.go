package main

import (
	"errors"
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
	if len(dataset) <= 1 {
		handleError(errors.New("empty dataset"), "Error: dataset is empty")
	}

	numericalFeatures := identifyNumericalFeatures(dataset)
	err := os.MkdirAll("tmp", 0755)
	handleError(err, "Error: could not create folder tmp")
	for _, i := range numericalFeatures {
		createHistogram(dataset[0][i], i, dataset)
	}
	combineFeatureImages(dataset, numericalFeatures)
}
