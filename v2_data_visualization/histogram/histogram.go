package main

import (
	"fmt"
	"os"
)

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(0)
	}
}

func main() {
	dataset := readDataset()
	numericalFeatures := identifyNumericalFeatures(dataset)
	floatDataset := parseFloatData(dataset[1:], numericalFeatures)
	houses := make([]string, 0)

	for i := 1; i < len(dataset); i++ {
		houses = append(houses, dataset[i][1])
	}
	
	os.Mkdir("Histograms",0755)
	os.Mkdir("tmp",0755)
	create_histogram(dataset[0][7], houses, floatDataset[7])
	// for _, i := range numericalFeatures {
	// 	create_histogram(dataset[0][i], houses, floatDataset[i])
	// }

}
