package main

import (
	"errors"
	"fmt"
	"image/color"
	"image/jpeg"
	"math"
	"os"
	"sort"
	"strconv"

	gim "github.com/ozankasikci/go-image-merge"
	"github.com/vdobler/chart"
	"gonum.org/v1/plot/plotter"
)

const HUFFLEPUFF = 0
const RAVENCLAW = 1
const GRYFFINDOR = 2
const SLYTHERIN = 3

var houseNames = [...]string{"Hufflepuff", "Ravenclaw", "Gryffindor", "Slytherin"}

func combineFeatureImages(dataset [][]string, numericalFeatures []int) {
	grids := make([]*gim.Grid, 0)
	for _, i := range numericalFeatures {
		g := gim.Grid{ImageFilePath: "tmp/" + dataset[0][i] + ".jpeg"}
		grids = append(grids, &g)
	}

	dimension := math.Ceil(math.Sqrt(float64(len(grids))))
	rgba, err := gim.New(grids, int(dimension), int(dimension)).Merge()
	handleError(err, "Error: gim could not merge feature images")

	file, err := os.Create("stackedHistogram.jpeg")
	handleError(err, "Error: could not create stackedHistogram.jpeg")
	err = jpeg.Encode(file, rgba, nil)
	handleError(err, "Error: gim could save merged images")
	fmt.Println("stackedHistogram.jpeg created")

	os.RemoveAll("tmp/")
}

func createHistogram(featureName string, featureIndex int, dataset [][]string) {
	allValues := make([]float64, 0)
	houseValues := make([]plotter.Values, 4)
	index := -1

	if len(dataset[0]) <= 1 {
		handleError(errors.New("dataset is invalid"), "Error: dataset is invalid")
	}
	data := dataset[1:]
	for i := range data {
		if len(data[i]) <= 1 {
			handleError(errors.New("row is invalid"), "Error: row is invalid")
		}
		switch data[i][1] {
		case "Hufflepuff":
			index = HUFFLEPUFF
		case "Ravenclaw":
			index = RAVENCLAW
		case "Gryffindor":
			index = GRYFFINDOR
		case "Slytherin":
			index = SLYTHERIN
		}
		if index == -1 {
			handleError(errors.New("house not found"), "Error: house not found")
		}
		if data[i][featureIndex] != "" {
			val, err := strconv.ParseFloat(data[i][featureIndex], 64)
			handleError(err, "Could not parse \""+data[i][featureIndex]+"\" to float")
			houseValues[index] = append(houseValues[index], val)
			allValues = append(allValues, val)
		}
	}

	sort.Float64s(allValues)

	hist := chart.HistChart{Title: featureName, Stacked: false, Counts: false}
	hist.XRange.Label = "Sample Value"
	hist.YRange.Label = "Rel. Frequency [%]"

	points := houseValues[0]
	if len(points) > 0 {
		hist.AddData("HUFFLEPUFF", points,
			chart.Style{LineColor: color.NRGBA{0xff, 0x00, 0x00, 0xff}, LineWidth: 1, FillColor: color.NRGBA{0xff, 0x80, 0x80, 0xff}})
	}

	points2 := houseValues[1]
	if len(points2) > 0 {
		hist.AddData("RAVENCLAW", points2,
			chart.Style{LineColor: color.NRGBA{0x00, 0xff, 0x00, 0xff}, LineWidth: 1, FillColor: color.NRGBA{0x80, 0xff, 0x80, 0xff}})
	}

	points3 := houseValues[2]
	if len(points3) > 0 {
		hist.AddData("GRYFFINDOR", points3,
			chart.Style{LineColor: color.NRGBA{0x00, 0x00, 0xff, 0xff}, LineWidth: 1, FillColor: color.NRGBA{0x80, 0x80, 0xff, 0xff}})
	}

	points4 := houseValues[3]
	if len(points4) > 0 {
		hist.AddData("SLYTHERIN", points4,
			chart.Style{LineColor: color.NRGBA{0x00, 0xff, 0xff, 0x00}, LineWidth: 1, FillColor: color.NRGBA{0x80, 0xff, 0xff, 0x80}})
	}

	dumper := NewDumper("tmp/"+featureName, 1, 1, 1000, 1000)
	defer dumper.Close()

	dumper.Plot(&hist)
	hist.Reset()
}
