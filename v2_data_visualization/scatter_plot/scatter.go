package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"strconv"

	gim "github.com/ozankasikci/go-image-merge"
	"github.com/vdobler/chart"
)

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(0)
	}
}

func featuresToScatter(s chart.ScatterChart, dataset [][]string, i1 int, i2 int) chart.ScatterChart {
	huf := make([]chart.EPoint, 0)
	gry := make([]chart.EPoint, 0)
	sly := make([]chart.EPoint, 0)
	rav := make([]chart.EPoint, 0)

	for i := range dataset {
		if i != 0 && dataset[i][i1] != "" && dataset[i][i2] != "" {
			val1, err := strconv.ParseFloat(dataset[i][i1], 64)
			handleError(err, "Could not parse \""+dataset[i][i1]+"\" to float")
			val2, err := strconv.ParseFloat(dataset[i][i2], 64)
			handleError(err, "Could not parse \""+dataset[i][i2]+"\" to float")
			switch dataset[i][1] {
			case "Hufflepuff":
				huf = append(huf, chart.EPoint{X: val1, Y: val2})
			case "Ravenclaw":
				rav = append(rav, chart.EPoint{X: val1, Y: val2})
			case "Gryffindor":
				gry = append(gry, chart.EPoint{X: val1, Y: val2})
			case "Slytherin":
				sly = append(sly, chart.EPoint{X: val1, Y: val2})
			}
		}
	}
	s.AddData("Hufflepuff", huf, chart.PlotStylePoints, chart.Style{})
	s.AddData("Gryffindor", gry, chart.PlotStylePoints, chart.Style{})
	s.AddData("Slytherin", sly, chart.PlotStylePoints, chart.Style{})
	s.AddData("Ravenclaw", rav, chart.PlotStylePoints, chart.Style{})
	return s
}

func plotScatter(dataset [][]string, numericalFeatures []int) {
	grids := make([]*gim.Grid, 0)
	err := os.MkdirAll("tmp", 0755)
	handleError(err, "Error: could not create directory")

	for i := 0; i < len(numericalFeatures)-1; i++ {
		for j := i + 1; j < len(numericalFeatures); j++ {
			featureI := dataset[0][numericalFeatures[i]]
			featureJ := dataset[0][numericalFeatures[j]]

			s := chart.ScatterChart{Title: featureI + " - " + featureJ}
			s.XRange.Label = featureI
			s.YRange.Label = featureJ

			s = featuresToScatter(s, dataset, numericalFeatures[i], numericalFeatures[j])

			dumper := NewDumper("tmp/"+featureI+"-"+featureJ, 1, 1, 1000, 1000)
			dumper.Plot(&s)
			dumper.Close()

			g := gim.Grid{ImageFilePath: "tmp/" + featureI + "-" + featureJ + ".jpeg"}
			grids = append(grids, &g)
		}
	}
	// create merged image
	rgba, err := gim.New(grids, len(numericalFeatures)-1, (len(numericalFeatures))/2).Merge()
	handleError(err, "Error: gim could not merge feature images (did you run out of space?)")

	file, err := os.Create("scatter.jpeg")
	handleError(err, "Error: could not create scatter")
	err = jpeg.Encode(file, rgba, nil)
	handleError(err, "Error: gim could save merged images")

	os.RemoveAll("tmp/")
}

func main() {
	dataset := readDataset()
	numericalFeatures := identifyNumericalFeatures(dataset)
	plotScatter(dataset, numericalFeatures)
}
