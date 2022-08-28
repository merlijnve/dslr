package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func findIndexOfFeature(dataset [][]string, feature string) int {
	for i := range dataset[0] {
		if dataset[0][i] == feature {
			return i
		}
	}
	return -1
}

func findMean(c *Classifier) Classifier {
	for i := range c.data {
		c.Mean0 += c.data[i][0]
		c.Mean1 += c.data[i][1]
	}
	c.Mean0 /= float64(len(c.data))
	c.Mean1 /= float64(len(c.data))
	return *c
}

func findStd(c *Classifier) Classifier {
	for i := range c.data {
		c.Std0 += math.Pow(c.data[i][0]-c.Mean0, 2)
		c.Std1 += math.Pow(c.data[i][1]-c.Mean1, 2)
	}
	c.Std0 = math.Sqrt(c.Std0 / float64(len(c.data)))
	c.Std1 = math.Sqrt(c.Std1 / float64(len(c.data)))
	return *c
}

func standardization(c Classifier) Classifier {
	c = findMean(&c)
	c = findStd(&c)
	for i := range c.data {
		c.data[i][0] = (c.data[i][0] - c.Mean0) / c.Std0
		c.data[i][1] = (c.data[i][1] - c.Mean1) / c.Std1
	}
	return c
}

func getDataPairs(dataset [][]string, c Classifier) [][]float64 {
	dataPairs := make([][]float64, 0)
	i0 := findIndexOfFeature(dataset, c.Feature0)
	i1 := findIndexOfFeature(dataset, c.Feature1)
	isClass := 0.0

	if (i0 == -1) || (i1 == -1) {
		handleError(errors.New("feature not found"), "Error: feature not found in dataset")
	}
	for i := range dataset {
		if i != 0 && dataset[i][i0] != "" && dataset[i][i1] != "" {
			val0, err := strconv.ParseFloat(dataset[i][i0], 64)
			handleError(err, "Error: could not parse "+dataset[i][i0])

			val1, err2 := strconv.ParseFloat(dataset[i][i1], 64)
			handleError(err2, "Error: could not parse "+dataset[i][i1])

			if dataset[i][1] == c.House {
				isClass = 1.0
			} else {
				isClass = 0.0
			}
			dataPairs = append(dataPairs, []float64{val0, val1, isClass})
		}
	}
	return dataPairs
}

func readDataset() [][]string {
	var file *os.File
	var err error

	if len(os.Args) > 1 {
		file, err = os.Open(os.Args[1])
		handleError(err, "Error: could not open file \""+os.Args[1]+"\"")
		defer file.Close()
	} else {
		fmt.Println("Use ./logreg_train [dataset filename]")
		os.Exit(0)
	}

	csv := csv.NewReader(file)
	records, err := csv.ReadAll()
	handleError(err, "Error: could not parse csv")
	return records
}
