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

func accuracy(c Classifier) float64 {
	correct_pred := 0

	for i := range c.data {
		pred := predict(c, i)
		if math.Round(pred) == c.data[i][2] {
			correct_pred += 1
		}
	}
	accuracy := float64(correct_pred) / float64(len(c.data)) * 100.0
	return accuracy
}

func calcThetas(c Classifier, learningRate float64) (Classifier, float64) {
	m := float64(len(c.data))

	grad_t0 := 0.0
	grad_t1 := 0.0
	grad_t2 := 0.0
	deltaGradient := 0.0

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
	deltaGradient = math.Abs(deltaGradient - (grad_t0 + grad_t1 + grad_t2))
	return c, deltaGradient
}

func gradientDescent(c Classifier) Classifier {
	fmt.Println("2. Starting Gradient Descent")

	grad := 1.0
	for i := 0; grad > 0.001; i++ {
		c, grad = calcThetas(c, 0.03)
	}
	c.Accuracy = accuracy(c)
	fmt.Println("3. Finished with training set prediction accuracy of:\n", c.Accuracy, "%")
	return c
}
