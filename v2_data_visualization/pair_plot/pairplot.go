package main

import (
	"errors"
	"fmt"
	"image/jpeg"
	"math"
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

	dimension := math.Ceil(math.Sqrt(float64(len(grids))))
	rgba, err := gim.New(grids, int(dimension), int(dimension)).Merge()
	handleError(err, "Error: gim could not merge feature images")

	file, err := os.Create("pairplot.jpeg")
	handleError(err, "Error: could not create pairplot")
	err = jpeg.Encode(file, rgba, nil)
	handleError(err, "Error: gim could not save merged images")
}

func main() {
	dataset := readDataset()
	if len(dataset) <= 1 {
		handleError(errors.New("empty dataset"), "Error: dataset is empty")
	}

	numericalFeatures := identifyNumericalFeatures(dataset)

	err := os.MkdirAll("tmp", 0755)
	handleError(err, "Error: could not create folder tmp")
	
	fmt.Println("Creating feature images...")
	plotScatter(dataset, numericalFeatures)
	for _, i := range numericalFeatures {
		createHistogram(dataset[0][i], i, dataset)
	}
	fmt.Println("Creating pairplot...")
	combineImages(dataset, numericalFeatures)
	fmt.Println("Done: created pairplot.jpeg")
	os.RemoveAll("tmp/")
}
