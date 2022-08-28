package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

type Classifier struct {
	House      string
	Feature0   string
	Feature1   string
	data       [][]float64
	T0         float64
	T1         float64
	T2         float64
	Prediction float64
	Mean0      float64
	Mean1      float64
	Std0       float64
	Std1       float64
}

// sigmoid function
func sigmoid(z float64) float64 {
	return 1 / (1 + math.Exp(-z))
}

// hypothesis function
// h = θ0 + θ1 * x1 + θ2 * x2
func h(c Classifier, x1 float64, x2 float64) float64 {
	return c.T0 + c.T1*x1 + c.T2*x2
}

func predict(c Classifier, f0 float64, f1 float64) float64 {
	f0 = (f0 - c.Mean0) / c.Std0
	f1 = (f1 - c.Mean1) / c.Std1
	return sigmoid(h(c, f0, f1))
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(0)
	}
}

func readThetaFile() []Classifier {
	classifiers := make([]Classifier, 0)
	var file *os.File
	var err error

	fmt.Println("Reading thetas file...")
	file, err = os.Open(os.Args[2])
	handleError(err, "Error: could not open file \""+os.Args[2]+"\"")
	byteValue, err := ioutil.ReadAll(file)
	handleError(err, "Error: could not read file \""+os.Args[2]+"\"")
	json.Unmarshal(byteValue, &classifiers)

	defer file.Close()

	return classifiers
}

func readDataset() [][]string {
	var file *os.File
	var err error
	var dataset [][]string

	fmt.Println("Reading dataset...")
	if len(os.Args) > 1 {
		file, err = os.Open(os.Args[1])
		handleError(err, "Error: could not open dataset file \""+os.Args[1]+"\"")
		csv := csv.NewReader(file)
		dataset, err = csv.ReadAll()
		handleError(err, "Error: could not parse csv")

		defer file.Close()
	} else {
		fmt.Println("Use ./logreg_predict [dataset filename] [theta filename]")
		os.Exit(0)
	}
	return dataset
}

func getFeatureIndexes(f1 string, f2 string, dataset [][]string) (int, int) {
	i1 := 0
	i2 := 0

	for i := range dataset[0] {
		if dataset[0][i] == f1 {
			i1 = i
		}
		if dataset[0][i] == f2 {
			i2 = i
		}
	}
	return i1, i2
}

func comparePredictionsAndReturnHighest(classifiers []Classifier) string {
	var highestClassifier Classifier
	highestClassifier = classifiers[0]
	for _, classifier := range classifiers {
		if classifier.Prediction > highestClassifier.Prediction {
			highestClassifier = classifier
		}
	}
	return highestClassifier.House
}

func writeLineToFile(file *os.File, line string) {
	_, err := file.WriteString(line + "\n")
	handleError(err, "Error: could not write to file")
}

func getDataPairs(dataset [][]string, c Classifier) [][]float64 {
	dataPairs := make([][]float64, 0)
	i0, i1 := getFeatureIndexes(c.Feature0, c.Feature1, dataset)
	var err error
	var err2 error

	dataPairs = append(dataPairs, []float64{0.0, 0.0})
	for i := range dataset {
		if i != 0 {
			val0 := c.Mean0
			if dataset[i][i0] != "" {
				val0, err = strconv.ParseFloat(dataset[i][i0], 64)
				handleError(err, "Error: could not parse "+dataset[i][i0])
			}
			val1 := c.Mean1
			if dataset[i][i1] != "" {
				val1, err2 = strconv.ParseFloat(dataset[i][i1], 64)
				handleError(err2, "Error: could not parse "+dataset[i][i1])
			}

			dataPairs = append(dataPairs, []float64{val0, val1})
		}
	}
	return dataPairs
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Use ./logreg_predict [dataset filename] [theta filename]")
		os.Exit(0)
	}
	dataset := readDataset()
	if len(dataset) == 0 {
		handleError(errors.New("dataset is empty"), "Error: dataset is empty")
	}
	classifiers := readThetaFile()

	for cI := range classifiers {
		classifiers[cI].data = getDataPairs(dataset, classifiers[cI])
	}

	prediction_file, err := os.Create("houses.csv")
	handleError(err, "Error: could not create prediction file (maybe ran out of memory?)")
	writeLineToFile(prediction_file, "Index,Hogwarts House")

	fmt.Println("Predicting...")
	for i := range dataset {
		if i != 0 {
			for cI := range classifiers {
				classifiers[cI].Prediction = predict(classifiers[cI], classifiers[cI].data[i][0], classifiers[cI].data[i][1])
			}
			result := comparePredictionsAndReturnHighest(classifiers)
			writeLineToFile(prediction_file, dataset[i][0]+","+result)
		}
	}
	fmt.Println("Done! Wrote predictions to houses.csv")
}
