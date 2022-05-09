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

	err := os.MkdirAll("tmp", 0755)
	handleError(err, "Error: could not create folder tmp")
	for _, i := range numericalFeatures {
		createHistogram(dataset[0][i], i, dataset)
	}
	combineFeatureImages(dataset, numericalFeatures)
}
