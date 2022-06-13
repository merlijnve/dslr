package main

import (
	"math"
)

// hypothesis function
// h = θ0 + θ1 * x1 + θ2 * x2
func h(c Classifier, x1 float64, x2 float64) float64 {
	return c.T0 + c.T1*x1 + c.T2*x2
}

// sigmoid function
func sigmoid(z float64) float64 {
	return 1 / (1 + math.Exp(-z))
}

// func normalizeDataset(dataset Dataset) Dataset {
// 	var normalizedDataset Dataset

// 	// create a copy of the dataset
// 	normalizedDataset.xName = dataset.xName
// 	normalizedDataset.yName = dataset.yName
// 	normalizedDataset.xMax = dataset.xMax
// 	normalizedDataset.xMin = dataset.xMin
// 	normalizedDataset.yMax = dataset.yMax
// 	normalizedDataset.yMin = dataset.yMin
// 	normalizedDataset.data = make([]Coordinate, len(dataset.data))
// 	copy(normalizedDataset.data, dataset.data)

// 	// normalize
// 	for i, c := range dataset.data {
// 		normalizedDataset.data[i].Y = (c.Y - dataset.yMin) / (dataset.yMax - dataset.yMin)
// 		normalizedDataset.data[i].X = (c.X - dataset.xMin) / (dataset.xMax - dataset.xMin)
// 	}

// 	return normalizedDataset
// }

// func denormalizeThetas(dataset Dataset, theta0 float64, theta1 float64) (float64, float64) {

// 	theta1 = (dataset.yMax - dataset.yMin) * theta1 / (dataset.xMax - dataset.xMin)
// 	theta0 = dataset.yMin + ((dataset.yMax - dataset.yMin) * theta0) + theta1*(1-dataset.xMin)

// 	return theta0, theta1
// }

func gradientDescent(dataset [][]float64) (float64, float64) {
	var theta0, theta1 float64 = 0, 0
	var delta0, delta1 float64 = 1, 1
	var prevCost0, prevCost1 float64
	m := float64(len(dataset))

	for delta0 > 0.000001 || delta1 > 0.000001 {
		cost0 := 0.0
		cost1 := 0.0
		for _, c := range dataset {
			// c[0] is c.x
			// c[1] is c.y
			cost0 += c[0]
		}
		delta0 = math.Abs(prevCost0 - cost0)
		delta1 = math.Abs(prevCost1 - cost1)
		theta0 -= 0.01 * (cost0 / m)
		theta1 -= 0.01 * (cost1 / m)
		prevCost0 = cost0
		prevCost1 = cost1
	}

	return theta0, theta1
}
