package main

import (
	"encoding/csv"
	"encoding/json"
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
	T0         float64
	T1         float64
	T2         float64
	Prediction float64
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
	return sigmoid(h(c, f0, f1))
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(0)
	}
}

func readThetaFile() ([][]string, []Classifier) {
	classifiers := make([]Classifier, 0)
	var file *os.File
	var err error
	var dataset [][]string

	if len(os.Args) > 2 {
		file, err = os.Open(os.Args[1])
		handleError(err, "Error: could not open file \""+os.Args[1]+"\"")
		csv := csv.NewReader(file)
		dataset, err = csv.ReadAll()
		handleError(err, "Error: could not parse csv")

		file, err = os.Open(os.Args[2])
		handleError(err, "Error: could not open file \""+os.Args[2]+"\"")
		byteValue, err := ioutil.ReadAll(file)
		handleError(err, "Error: could not read file \""+os.Args[2]+"\"")
		json.Unmarshal(byteValue, &classifiers)

		defer file.Close()
	} else {
		fmt.Println("Use ./logreg_predict [dataset filename] [thetas filename]")
		os.Exit(0)
	}
	return dataset, classifiers
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

func comparePredictionsAndReturnHighest(classifiers []Classifier) {
	// for i := range classifiers {
	// 	fmt.Println(classifiers[i].House, classifiers[i].Prediction)
	// }
	var highestClassifier Classifier
	highestClassifier = classifiers[0]
	for _, classifier := range classifiers {
		if classifier.Prediction > highestClassifier.Prediction {
			highestClassifier = classifier
		}
	}
	fmt.Println("Predicted:", highestClassifier.House)
}

func main() {
	dataset, classifiers := readThetaFile()
	// _, classifiers := readThetaFile()
	for i := range dataset {
		if i != 0 {
			for cI := range classifiers {
				// fmt.Println(classifiers[cI])
				i1, i2 := getFeatureIndexes(classifiers[cI].Feature0, classifiers[cI].Feature1, dataset)
				// fmt.Println(i, i1, i2)
				if dataset[i][i1] != "" && dataset[i][i2] != "" {
					val1, err := strconv.ParseFloat(dataset[i][i1], 64)
					handleError(err, "Error: could not parse i1 "+dataset[i][i1])
					val2, err := strconv.ParseFloat(dataset[i][i2], 64)
					handleError(err, "Error: could not parse i2"+dataset[i][i2])
					classifiers[cI].Prediction = predict(classifiers[cI], val1, val2)
				}
			}
		}
		comparePredictionsAndReturnHighest(classifiers)
	}
}
