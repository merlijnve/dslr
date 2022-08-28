package main

import (
	"errors"
	"fmt"
	"image/jpeg"
	"math"
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
	if len(huf) > 0 {
		s.AddData("Hufflepuff", huf, chart.PlotStylePoints, chart.Style{})
	}
	if len(gry) > 0 {
		s.AddData("Gryffindor", gry, chart.PlotStylePoints, chart.Style{})
	}
	if len(sly) > 0 {
		s.AddData("Slytherin", sly, chart.PlotStylePoints, chart.Style{})
	}
	if len(rav) > 0 {
		s.AddData("Ravenclaw", rav, chart.PlotStylePoints, chart.Style{})
	}
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

			if len(s.Data) == 0 {
				handleError(errors.New("no houses in dataset"), "Error: no houses in dataset")
			}
			fmt.Println("Plotting " + featureI + " - " + featureJ)
			dumper := NewDumper("tmp/"+featureI+"-"+featureJ, 1, 1, 1000, 1000)
			dumper.Plot(&s)
			dumper.Close()

			g := gim.Grid{ImageFilePath: "tmp/" + featureI + "-" + featureJ + ".jpeg"}
			grids = append(grids, &g)
		}
	}
	// create merged image
	fmt.Println("Creating merged image")
	dimension := math.Ceil(math.Sqrt(float64(len(grids))))
	rgba, err := gim.New(grids, int(dimension), int(dimension)).Merge()
	handleError(err, "Error: gim could not merge feature images (did you run out of space?)")

	file, err := os.Create("scatter.jpeg")
	handleError(err, "Error: could not create scatter")
	err = jpeg.Encode(file, rgba, nil)
	handleError(err, "Error: gim could save merged images")

	fmt.Println("Removing tmp/*")
	os.RemoveAll("tmp/")
}

func main() {
	dataset := readDataset()
	if len(dataset) <= 1 {
		handleError(errors.New("empty dataset"), "Error: dataset is empty")
	}

	numericalFeatures := identifyNumericalFeatures(dataset)
	if len(numericalFeatures) < 2 {
		handleError(errors.New("not enough features"), "Error: not enough (numerical) features in dataset")
	}
	plotScatter(dataset, numericalFeatures)
	fmt.Println("Done: created scatter.jpeg")
}
