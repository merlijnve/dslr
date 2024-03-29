package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"
)

func formatCount(floatData [][]float64, numericalFeatures []int) string {
	result := "Count\t"

	for _, e := range numericalFeatures {
		result += strconv.Itoa(len(floatData[e])) + "\t"
	}
	return result
}

// square root of the sum of the squares of the differences between each value and the mean
func formatStd(floatData [][]float64, numericalFeatures []int) string {
	result := "Std\t"

	for _, i := range numericalFeatures {
		var sum float64 = 0
		var count float64 = float64(len(floatData[i]))
		var sumSquares float64 = 0
		for _, c := range floatData[i] {
			sum += c
		}
		mean := sum / count
		for _, c := range floatData[i] {
			deviation := c - mean
			sumSquares += deviation * deviation
		}
		variance := sumSquares / (count - 1)
		standardDeviation := math.Sqrt(variance)
		result += strconv.FormatFloat(standardDeviation, 'f', 6, 64) + "\t"
	}
	return result
}

// sum / count
func formatMean(floatData [][]float64, numericalFeatures []int) string {
	result := "Mean\t"

	for _, i := range numericalFeatures {
		var sum float64 = 0
		var count float64 = 0
		for j := range floatData[i] {
			count++
			sum += floatData[i][j]
		}
		result += strconv.FormatFloat(sum/count, 'f', 6, 64) + "\t"
	}
	return result
}

// formats the min and max values of the numerical features
// returns a formatted string as result
func formatMinMax(floatData [][]float64, numericalFeatures []int, flag string) string {
	result := ""
	switch flag {
	case "min":
		result = "Min\t"
	case "max":
		result = "Max\t"
	}

	for _, i := range numericalFeatures {
		switch flag {
		case "min":
			result += strconv.FormatFloat(floatData[i][0], 'f', 6, 64) + "\t"
		case "max":
			result += strconv.FormatFloat(floatData[i][len(floatData[i])-1], 'f', 6, 64) + "\t"
		}
	}
	return result
}

// formats the feature names into a string
func formatFeatures(dataset [][]string, numericalFeatures []int) string {
	result := "\t"

	for _, i := range numericalFeatures {
		result += dataset[0][i] + "\t"
	}
	return result
}

// calculates percentiles using interpolation
// returns a formatted string as result
func formatPercentiles(floatData [][]float64, numericalFeatures []int, flag string) string {
	result := ""

	switch flag {
	case "25":
		result = "25%\t"
	case "50":
		result = "50%\t"
	case "75":
		result = "75%\t"
	}
	for _, i := range numericalFeatures {
		rank := 0.0
		switch flag {
		case "25":
			rank = float64(0.25) * float64((len(floatData[i]) - 1))
		case "50":
			rank = float64(0.50) * float64((len(floatData[i]) - 1))
		case "75":
			rank = float64(0.75) * float64((len(floatData[i]) - 1))
		}
		// checks if value has fractional part or is integer
		if rank == float64(int(rank)) {
			result += strconv.FormatFloat(floatData[i][int(rank)], 'f', 6, 64) + "\t"
		} else {
			// interpolate in case of rank with fractional part
			fraction := math.Abs(rank - float64(int(rank)))
			rank = math.Floor(rank)
			val0 := floatData[i][int(rank)]
			val1 := val0
			if (rank + 1.0) < float64(len(floatData[i])) {
				val1 = floatData[i][int(rank)+1]
			}
			val := val0 + (val1-val0)*fraction
			result += strconv.FormatFloat(val, 'f', 6, 64) + "\t"
		}
	}
	return result
}

// parses the dataset to floats
func parseFloatData(dataset [][]string, numericalFeatures []int) [][]float64 {
	floatData := make([][]float64, len(dataset[0]))

	for _, i := range numericalFeatures {
		var count float64 = 0
		for j := range dataset {
			if dataset[j][i] != "" {
				count++
				val, err := strconv.ParseFloat(dataset[j][i], 64)
				if err != nil {
					handleError(err, "Could not parse \""+dataset[j][i]+"\" to float")
				}
				floatData[i] = append(floatData[i], val)
			}
		}
		sort.Float64s(floatData[i])
	}
	return floatData
}

func displayInformation(dataset [][]string, numericalFeatures []int) {
	floatData := parseFloatData(dataset[1:], numericalFeatures)
	w := tabwriter.NewWriter(os.Stdout, 16, 0, 2, ' ', 1)

	fmt.Fprintln(w, formatFeatures(dataset, numericalFeatures))
	fmt.Fprintln(w, formatCount(floatData, numericalFeatures))
	fmt.Fprintln(w, formatMean(floatData, numericalFeatures))
	fmt.Fprintln(w, formatStd(floatData, numericalFeatures))
	fmt.Fprintln(w, formatMinMax(floatData, numericalFeatures, "min"))

	fmt.Fprintln(w, formatPercentiles(floatData, numericalFeatures, "25"))
	fmt.Fprintln(w, formatPercentiles(floatData, numericalFeatures, "50"))
	fmt.Fprintln(w, formatPercentiles(floatData, numericalFeatures, "75"))

	fmt.Fprintln(w, formatMinMax(floatData, numericalFeatures, "max"))
	w.Flush()

}
