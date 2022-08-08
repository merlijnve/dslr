package main

import (
	"fmt"
	"math"
)

// sigmoid function
func sigmoid(z float64) float64 {
	return 1 / (1 + math.Exp(-z))
}

// hypothesis function
// h = θ0 + θ1 * x1 + θ2 * x2
func h(c Classifier, x1 float64, x2 float64) float64 {
	return c.T0 + c.T1*x1 + c.T2*x2
}

func predict(c Classifier, i int) float64 {
	return sigmoid(h(c, c.data[i][0], c.data[i][1]))
}

func logloss(c Classifier) float64 {
	loss := 0.0

	for i := range c.data {
		realValue := c.data[i][2]
		prediction := sigmoid(h(c, c.data[i][0], c.data[i][1]))
		loss += -realValue*math.Log(prediction) - (1-realValue)*math.Log(1-prediction)
	}
	return loss
}

func accuracy(c Classifier) float64 {
	correct_pred := 0

	for i := range c.data {
		pred := predict(c, i)
		if math.Round(pred) == c.data[i][2] {
			correct_pred += 1
		}
	}
	accuracy := float64(correct_pred) / float64(len(c.data)) * 100.0
	fmt.Println(accuracy, "% accuracy\n")
	return accuracy
}

func calcThetas(c Classifier, learningRate float64) Classifier {
	m := float64(len(c.data))

	grad_t0 := 0.0
	grad_t1 := 0.0
	grad_t2 := 0.0

	for i := range c.data {
		prediction := predict(c, i)
		real := c.data[i][2]

		loss := prediction - real
		grad_t0 += 1 / m * loss
		grad_t1 += 1 / m * loss * c.data[i][0]
		grad_t2 += 1 / m * loss * c.data[i][1]
	}
	tmp0 := c.T0
	tmp1 := c.T1
	tmp2 := c.T2
	c.T0 = tmp0 - (learningRate * grad_t0)
	c.T1 = tmp1 - (learningRate * grad_t1)
	c.T2 = tmp2 - (learningRate * grad_t2)
	return c
}

func gradientDescent(c Classifier) Classifier {
	fmt.Println("Gradient descent for", c.House)

	for i := 0; i < 10000; i++ {
		c = calcThetas(c, 0.001)
	}
	fmt.Println("THETAS", c.T0, c.T1, c.T2)
	accuracy(c)
	return c
}
