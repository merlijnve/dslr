package main

import (
	"fmt"
	"image/jpeg"
	"os"

	gim "github.com/ozankasikci/go-image-merge"
)

const HUFFLEPUFF = 0
const RAVENCLAW = 1
const GRYFFINDOR = 2
const SLYTHERIN = 3

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(0)
	}
}

func combineImages(dataset [][]string, numericalFeatures []int) {
	grids := make([]*gim.Grid, 0)

	// loop through all features and grid them
	for i := range numericalFeatures {
		for j := range numericalFeatures {
			if i == j {
				grids = append(grids, &gim.Grid{ImageFilePath: "tmp/" + dataset[0][numericalFeatures[i]] + ".jpeg"})
			} else {
				grids = append(grids, &gim.Grid{ImageFilePath: "tmp/" + dataset[0][numericalFeatures[i]] + "-" + dataset[0][numericalFeatures[j]] + ".jpeg"})
			}
		}
	}

	rgba, err := gim.New(grids, len(numericalFeatures), len(numericalFeatures)).Merge()
	handleError(err, "Error: gim could not merge feature images")

	file, err := os.Create("pairplot.jpeg")
	handleError(err, "Error: could not create pairplot")
	err = jpeg.Encode(file, rgba, nil)
	handleError(err, "Error: gim could save merged images")
}

func main() {
	dataset := readDataset()
	numericalFeatures := identifyNumericalFeatures(dataset)

	err := os.MkdirAll("tmp", 0755)
	handleError(err, "Error: could not create folder tmp")

	plotScatter(dataset, numericalFeatures)
	for _, i := range numericalFeatures {
		createHistogram(dataset[0][i], i, dataset)
	}
	combineImages(dataset, numericalFeatures)
	os.RemoveAll("tmp/")
}
