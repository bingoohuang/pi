package main

import (
	"fmt"
	"math"
)

func main() {
	var iter float64 = 1000000
	fmt.Printf("Estimating π using %.0f iterations\n", iter)
	π := leibnizπ(iter)
	absoluteError, relativeError, percentError := errorCalculation(π)
	fmt.Println("π:              ", π)
	fmt.Println("math.Pi:        ", math.Pi)
	fmt.Println("Absolute Error: ", absoluteError)
	fmt.Println("Relative Error: ", relativeError)
	fmt.Println("Percent Error:  ", percentError, "%")
}

func errorCalculation(π float64) (float64, float64, float64) {
	absoluteError := math.Abs(math.Pi - π)
	relativeError := math.Abs(1 - π/math.Pi)
	percentError := relativeError * 100
	return absoluteError, relativeError, percentError
}

func leibnizπ(iter float64) float64 {
	var sum, i float64 = 0, 0
	for ; i < iter; i++ {
		sum += math.Pow(-1, i) / (2*i + 1)
	}
	return sum * 4
}
