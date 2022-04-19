package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func create_histogram(featureName string, houses []string, data []float64) {
	var values plotter.Values

	for i := range data {
		values = append(values, data[i])
	}
	p := plot.New()
	p.Title.Text = featureName + " Distribution"
	// p.X.Label.Text = dataset.xName
	// p.Y.Label.Text = dataset.yName
	hist, err := plotter.NewHist(values, 20)
	hist.FillColor = color.RGBA{R: 39,G:159, B: 245, A: 255}
    if err != nil {
        handleError(err, "Error: could not make histogram")
    }
    p.Add(hist)

	err = p.Save(600, 600, "tmp/"+featureName + "_distribution.png")
	handleError(err, "Error: could not make histogram")
}
