package main

import (
	"fmt"
	"math"
)

func main() {
	k := 0.07
	tempRoom := 20.
	tempObject := 100.
	fcr := newCoolingRateDy(k, tempRoom)
	analytic := newTempFunc(k, tempRoom, tempObject)
	for _, deltaTime := range []float64{2, 5, 10} {
		fmt.Printf("Step size = %.1f\n", deltaTime)
		fmt.Println(" Time Euler's Analytic")
		temp := tempObject
		for time := 0.; time <= 100; time += deltaTime {
			fmt.Printf("%5.1f %7.3f %7.3f\n", time, temp, analytic(time))
			temp = eulerStep(fcr, time, temp, deltaTime)
		}
		fmt.Println()
	}

}

//fdy is a type for function f used in Euler's method
type fdy func(float64, float64) float64

//eulerStep computes a single new value using Euler's method
func eulerStep(f fdy, x, y, h float64) float64 {
	return y + h*f(x, y)
}

//newTempFunc returns a funtion that computes the analytical cooling rate
//for a given cooling rate constant k
func newCoolingRate(k float64) func(float64) float64 {
	return func(deltaTemp float64) float64 {
		return -k * deltaTemp
	}
}

//newTempFunc returns a function that computes the analytical solution of cooling rate integrated over time
func newTempFunc(k, ambientTemp, initialTemp float64) func(float64) float64 {
	return func(time float64) float64 {
		return ambientTemp + (initialTemp-ambientTemp)*math.Exp(-k*time)
	}
}

//newCoolingRateDy returns a function of the kind needed for Euler's method
//That is a function representing dy(x, y(x))

func newCoolingRateDy(k, ambientTemp float64) fdy {
	crf := newCoolingRate(k)
	return func(_, objectTemp float64) float64 {
		return crf(objectTemp - ambientTemp)
	}
}
