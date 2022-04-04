package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"text/tabwriter"
)

func formatCount(dataset [][]string, numericalFeatures []int) string {
	result := "Count\t"

	for _, i := range numericalFeatures {
		count := 0
		for j := range dataset {
			if dataset[j][i] != "" {
				count++
			}
		}
		result += strconv.Itoa(count) + "\t"
	}
	return result
}

func formatStd(dataset [][]string, numericalFeatures []int) string {
	result := "Std\t"

	for _, i := range numericalFeatures {
		var sum float64 = 0
		var count float64 = 0
		var sumSquares float64 = 0
		for j := range dataset {
			if dataset[j][i] != "" {
				count++
				val, err := strconv.ParseFloat(dataset[j][i], 64)
				if err != nil {
					handleError(err, "Could not parse \""+dataset[j][i]+"\" to float")
				}
				sum += val
			}
			mean := sum / count
			if dataset[j][i] != "" {
				val, err := strconv.ParseFloat(dataset[j][i], 64)
				if err != nil {
					handleError(err, "Could not parse \""+dataset[j][i]+"\" to float")
				}
				deviation := val - mean
				sumSquares += deviation * deviation
			}
		}
		variance := sumSquares / (count - 1)
		standardDeviation := math.Sqrt(variance)
		result += strconv.FormatFloat(standardDeviation, 'f', 6, 64) + "\t"
	}
	return result
}

func formatMean(dataset [][]string, numericalFeatures []int) string {
	result := "Mean\t"

	for _, i := range numericalFeatures {
		var sum float64 = 0
		var count float64 = 0
		for j := range dataset {
			if dataset[j][i] != "" {
				count++
				val, err := strconv.ParseFloat(dataset[j][i], 64)
				if err != nil {
					handleError(err, "Could not parse \""+dataset[j][i]+"\" to float")
				}
				sum += val
			}
		}
		result += strconv.FormatFloat(sum/count, 'f', 6, 64) + "\t"
	}
	return result
}

func formatMinMax(dataset [][]string, numericalFeatures []int, flag string) string {
	result := ""
	if flag == "min" {
		result += "Min\t"
	}
	if flag == "max" {
		result += "Max\t"
	}

	for _, i := range numericalFeatures {
		var endValue float64 = 0
		for j := range dataset {
			if dataset[j][i] != "" {
				val, err := strconv.ParseFloat(dataset[j][i], 64)
				if err != nil {
					handleError(err, "Could not parse \""+dataset[j][i]+"\" to float")
				}
				if flag == "min" && (j == 0 || val < endValue) {
					endValue = val
				}
				if flag == "max" && (j == 0 || val > endValue) {
					endValue = val
				}
			}
		}
		result += strconv.FormatFloat(endValue, 'f', 6, 64) + "\t"
	}
	return result
}

func formatFeatures(dataset [][]string, numericalFeatures []int) string {
	result := "\t"

	for _, i := range numericalFeatures {
		result += dataset[0][i] + "\t"
	}
	return result
}

func displayInformation(dataset [][]string, numericalFeatures []int) {

	w := tabwriter.NewWriter(os.Stdout, 16, 0, 2, ' ', 1)

	fmt.Fprintln(w, formatFeatures(dataset, numericalFeatures))
	fmt.Fprintln(w, formatCount(dataset[1:], numericalFeatures))
	fmt.Fprintln(w, formatMean(dataset[1:], numericalFeatures))
	fmt.Fprintln(w, formatStd(dataset[1:], numericalFeatures))
	fmt.Fprintln(w, formatMinMax(dataset[1:], numericalFeatures, "min"))
	// Best to do all these in one function
	// Format 25%
	// Format 50% (find the median)
	// Format 75%
	fmt.Fprintln(w, formatMinMax(dataset[1:], numericalFeatures, "max"))
	w.Flush()

}
