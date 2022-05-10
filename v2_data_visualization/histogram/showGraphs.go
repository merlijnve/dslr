package main

import (
	"image/color"
	"image/jpeg"
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

	rgba, err := gim.New(grids, len(grids)/3, 4).Merge()
	handleError(err, "Error: gim could not merge feature images")

	file, err := os.Create("stackedHistogram.jpeg")
	handleError(err, "Error: could not create stackedHistogram.jpeg")
	err = jpeg.Encode(file, rgba, nil)
	handleError(err, "Error: gim could save merged images")

	os.RemoveAll("tmp/")
}

func createHistogram(featureName string, featureIndex int, dataset [][]string) {
	allValues := make([]float64, 0)
	houseValues := make([]plotter.Values, 4)
	index := 0

	data := dataset[1:]
	for i := range data {
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
		if data[i][featureIndex] != "" {
			val, err := strconv.ParseFloat(data[i][featureIndex], 64)
			handleError(err, "Could not parse \""+data[i][featureIndex]+"\" to float")
			houseValues[index] = append(houseValues[index], val)
			allValues = append(allValues, val)
		}
	}

	sort.Float64s(allValues)

	hist := chart.HistChart{Title: featureName, Stacked: true, Counts: false}
	hist.XRange.Label = "Sample Value"
	hist.YRange.Label = "Rel. Frequency [%]"

	points := houseValues[0]
	hist.AddData("HUFFLEPUFF", points,
		chart.Style{LineColor: color.NRGBA{0xff, 0x00, 0x00, 0xff}, LineWidth: 1, FillColor: color.NRGBA{0xff, 0x80, 0x80, 0xff}})

	points2 := houseValues[1]
	hist.AddData("RAVENCLAW", points2,
		chart.Style{LineColor: color.NRGBA{0x00, 0xff, 0x00, 0xff}, LineWidth: 1, FillColor: color.NRGBA{0x80, 0xff, 0x80, 0xff}})

	points3 := houseValues[2]
	hist.AddData("GRYFFINDOR", points3,
		chart.Style{LineColor: color.NRGBA{0x00, 0x00, 0xff, 0xff}, LineWidth: 1, FillColor: color.NRGBA{0x80, 0x80, 0xff, 0xff}})

	points4 := houseValues[3]
	hist.AddData("SLYTHERIN", points4,
		chart.Style{LineColor: color.NRGBA{0x00, 0xff, 0xff, 0x00}, LineWidth: 1, FillColor: color.NRGBA{0x80, 0xff, 0xff, 0x80}})

	dumper := NewDumper("tmp/"+featureName, 1, 1, 1000, 1000)
	defer dumper.Close()

	dumper.Plot(&hist)
	hist.Reset()
}
