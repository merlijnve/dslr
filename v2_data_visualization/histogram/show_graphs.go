package main

import (
	"image/color"
	"image/png"
	"os"
	"sort"
	"strconv"

	gim "github.com/ozankasikci/go-image-merge"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

const HUFFLEPUFF = 0
const RAVENCLAW = 1
const GRYFFINDOR = 2
const SLYTHERIN = 3

var houseNames = [...]string{"Hufflepuff", "Ravenclaw", "Gryffindor", "Slytherin"}

func createHouseHist(house string, featureName string, values plotter.Values, featureMin float64, featureMax float64) {
	p := plot.New()
	p.Title.Text = featureName + ": " + house + " Distribution"

	p.X.Max = featureMax
	p.X.Min = featureMin

	p.X.Label.Text = "Value"
	p.Y.Label.Text = "Frequency"
	hist, err := plotter.NewHist(values, 20)
	handleError(err, "Error: could not make histogram")
	switch house {
	case "Hufflepuff":
		hist.FillColor = color.RGBA{R: 127, G: 177, B: 239, A: 255}
	case "Ravenclaw":
		hist.FillColor = color.RGBA{R: 241, G: 202, B: 252, A: 255}
	case "Gryffindor":
		hist.FillColor = color.RGBA{R: 193, G: 255, B: 162, A: 255}
	case "Slytherin":
		hist.FillColor = color.RGBA{R: 255, G: 255, B: 112, A: 255}
	}

	p.Add(hist)

	err = p.Save(600, 600, "tmp/"+featureName+"_"+house+"_distribution.png")
	handleError(err, "Error: could not make histogram")
}

func combineHouseImages(featureName string) {
	grids := []*gim.Grid{
		{ImageFilePath: "tmp/" + featureName + "_Hufflepuff_distribution.png"},
		{ImageFilePath: "tmp/" + featureName + "_Ravenclaw_distribution.png"},
		{ImageFilePath: "tmp/" + featureName + "_Gryffindor_distribution.png"},
		{ImageFilePath: "tmp/" + featureName + "_Slytherin_distribution.png"},
	}

	rgba, err := gim.New(grids, 1, 4).Merge()
	handleError(err, "Error: gim could not merge house images")

	file, err := os.Create("tmp/" + featureName + ".png")
	err = png.Encode(file, rgba)
	handleError(err, "Error: gim could not save merged images")

	os.Remove("tmp/" + featureName + "_Hufflepuff_distribution.png")
	os.Remove("tmp/" + featureName + "_Ravenclaw_distribution.png")
	os.Remove("tmp/" + featureName + "_Gryffindor_distribution.png")
	os.Remove("tmp/" + featureName + "_Slytherin_distribution.png")

}

func combineFeatureImages(dataset [][]string, numericalFeatures []int) {
	grids := make([]*gim.Grid, 0)
	for _, i := range numericalFeatures {
		g := gim.Grid{ImageFilePath: "tmp/" + dataset[0][i] + ".png"}
		grids = append(grids, &g)
	}

	rgba, err := gim.New(grids, len(grids), 1).Merge()
	handleError(err, "Error: gim could not merge feature images")

	file, err := os.Create("histogram.png")
	err = png.Encode(file, rgba)
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
		case "Ravenclaw":
			index = RAVENCLAW
		case "Slytherin":
			index = SLYTHERIN
		case "Hufflepuff":
			index = HUFFLEPUFF
		case "Gryffindor":
			index = GRYFFINDOR
		}
		if data[i][featureIndex] != "" {
			val, err := strconv.ParseFloat(data[i][featureIndex], 64)
			handleError(err, "Could not parse \""+data[i][featureIndex]+"\" to float")
			houseValues[index] = append(houseValues[index], val)
			allValues = append(allValues, val)
		}
	}

	sort.Float64s(allValues)
	for i := range houseValues {
		createHouseHist(houseNames[i], featureName, houseValues[i], allValues[0], allValues[len(allValues)-1])
	}
	combineHouseImages(featureName)
}
