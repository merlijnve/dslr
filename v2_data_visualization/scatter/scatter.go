package main

import (
	"fmt"
	"os"
	"strconv"
	"image/jpeg"

	gim "github.com/ozankasikci/go-image-merge"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(0)
	}
}

func indexesToXYs(dataset [][]string, i1 int, i2 int) plotter.XYs {
	pts := make(plotter.XYs, 0)

	for i := range dataset {
		if i != 0 && dataset[i][i1] != "" && dataset[i][i2] != "" {
			val1, err := strconv.ParseFloat(dataset[i][i1], 64)
			handleError(err, "Could not parse \""+dataset[i][i1]+"\" to float")
			val2, err := strconv.ParseFloat(dataset[i][i2], 64)
			handleError(err, "Could not parse \""+dataset[i][i2]+"\" to float")
			pts = append(pts, plotter.XY{X: val1, Y: val2})
		}
	}
	return pts
}

func plotScatter(dataset [][]string, numericalFeatures []int) {
	grids := make([]*gim.Grid, 0)
	err := os.MkdirAll("tmp", 0755)
	handleError(err, "Error: could not create directory")

	for i := 1; i < len(numericalFeatures)-1; i++ {
		for j := i + 1; j < len(numericalFeatures); j++ {
			featureI := dataset[0][numericalFeatures[i]]
			featureJ := dataset[0][numericalFeatures[j]]

			p := plot.New()
			p.Title.Text = featureI + " - " + featureJ + " Scatter"
			p.X.Label.Text = featureI
			p.Y.Label.Text = featureJ
			xys := indexesToXYs(dataset, numericalFeatures[i], numericalFeatures[j])
			s, err := plotter.NewScatter(xys)
			handleError(err, "Error: could not add points of data")
			p.Add(s)
			err = p.Save(600, 600, "tmp/"+featureI+"-"+featureJ+".jpg")
			handleError(err, "Error: could not save plot of linear equation")
			g := gim.Grid{ImageFilePath: "tmp/" + featureI + "-" + featureJ + ".jpg"}
			grids = append(grids, &g)
		}
	}
	// create merged image
	rgba, err := gim.New(grids, len(numericalFeatures) - 2, (len(numericalFeatures) - 1) / 2).Merge()
	handleError(err, "Error: gim could not merge feature images")

	file, err := os.Create("scatter.jpg")
	err = jpeg.Encode(file, rgba, nil)
	handleError(err, "Error: gim could save merged images")

	os.RemoveAll("tmp/")
}

func main() {
	dataset := readDataset()
	numericalFeatures := identifyNumericalFeatures(dataset)
	houses := make([]string, 0)

	for i := 1; i < len(dataset); i++ {
		houses = append(houses, dataset[i][1])
	}

	plotScatter(dataset, numericalFeatures)
}
