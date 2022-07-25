package main

import (
	"image/color"
	"os"
	"sort"
	"strconv"

	"github.com/vdobler/chart"
	"gonum.org/v1/plot/plotter"
)

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
	err := os.MkdirAll("tmp", 0755)
	handleError(err, "Error: could not create directory")

	for i := range numericalFeatures {
		for j := range numericalFeatures {
			if i != j {
				featureI := dataset[0][numericalFeatures[i]]
				featureJ := dataset[0][numericalFeatures[j]]

				s := chart.ScatterChart{Title: featureI + " - " + featureJ}
				s.XRange.Label = featureI
				s.YRange.Label = featureJ

				s = featuresToScatter(s, dataset, numericalFeatures[i], numericalFeatures[j])

				dumper := NewDumper("tmp/"+featureI+"-"+featureJ, 1, 1, 1000, 1000)
				dumper.Plot(&s)
				dumper.Close()
			}
		}
	}
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
