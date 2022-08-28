package main

import (
	"image/color"
	"image/jpeg"
	"os"

	gim "github.com/ozankasikci/go-image-merge"
	"github.com/vdobler/chart"
)

func featuresToScatter(s chart.ScatterChart, c Classifier) chart.ScatterChart {
	f0 := make([]chart.EPoint, 0)
	f1 := make([]chart.EPoint, 0)

	for i := range c.data {
		if c.data[i][2] == 1.0 {
			f0 = append(f0, chart.EPoint{Y: c.data[i][1], X: c.data[i][0]})
		} else {
			f1 = append(f1, chart.EPoint{Y: c.data[i][1], X: c.data[i][0]})
		}
	}
	if len(f0) > 0 {
		s.AddData(c.House, f0, chart.PlotStylePoints, chart.Style{LineColor: color.NRGBA{0x00, 0xff, 0x00, 0xff}})
	}
	if len(f1) > 0 {
		s.AddData("others", f1, chart.PlotStylePoints, chart.Style{LineColor: color.NRGBA{0xff, 0x00, 0x00, 0xff}})
	}
	return s
}

// Creates scatter.jpeg which contains scatter plots with decision boundary for every classifier (one per house)
func plotScatter(classifiers []Classifier) {
	grids := make([]*gim.Grid, 0)
	err := os.MkdirAll("tmp", 0755)
	handleError(err, "Error: could not create directory")

	// For every classifier, create a scatter plot with decision boundary .jpeg
	for _, c := range classifiers {
		s := chart.ScatterChart{Title: c.Feature0 + " - " + c.Feature1}
		s.XRange.Label = c.Feature0
		s.YRange.Label = c.Feature1

		s = featuresToScatter(s, c)
		s.AddFunc("Decision boundary", func(x float64) float64 {
			return -(c.T1/c.T2)*x - (c.T0 / c.T2)
		}, chart.PlotStyleLines, chart.Style{})

		dumper := NewDumper("tmp/"+c.House, 1, 1, 1000, 1000)
		dumper.Plot(&s)
		dumper.Close()

		g := gim.Grid{ImageFilePath: "tmp/" + c.House + ".jpeg"}
		grids = append(grids, &g)
	}
	// Create the merged image
	rgba, err := gim.New(grids, 4, 1).Merge()
	handleError(err, "Error: gim could not merge feature images (did you run out of space again?)")

	file, err := os.Create("scatter.jpeg")
	handleError(err, "Error: could not create scatter")
	err = jpeg.Encode(file, rgba, nil)
	handleError(err, "Error: gim could save merged images")

	os.RemoveAll("tmp/")
}
