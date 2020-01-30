package main

import (
	"fmt"
	"math"
)

//We want to calculate the shape of an unknown curve which starts at a given point with a given slope. This curve satisfies an ordinary differential equation (ODE):
//
//dy/dx = f(x, y(x)) with y(x0) = y0.
//
//The starting point A0 (x0, y0) is known as well as the slope to the curve at A0 and then the tangent line at A0 .
//
//Take a small step along that tangent line up to a point A1. Along this small step, the slope does not change too much, so A1 will be close to the curve. If we suppose that A1 is close enough to the curve, the same reasoning as for the point A1 above can be used for other points. After several steps, a polygonal curve A0, A1, .. , An is computed. The error between the 2 curves can be small if the step is small.
//
//We define points A0, A1, A2, ... , An whose x-coordinates are x0, x1 , ... , xn and y-coordinates are such that yk+1 = yk + f(xk, yk) * h where h is the common step. If T is the length xn - x 0 we have h = T/n.
//
//We will take as an example the differential equation:
//
//dy/dt = 2 - exp(-4t) - 2y with x0 = 0, y0 = 1, T = 1.
//
//We know that an exact solution is y = 1 + 0.5exp(-4t) - 0.5exp(-2t). For each xk we are able to calculate the yk as well as the values zk of the exact solution.
//
//Our task is, for a given number n of steps, to return the mean (truncated to 6 decimals) of the relative errors between the yk of the n+1 points Ak of our polygonal line and the n+1 zk of our exact solution. For that we can use:
//
//error in Ak = abs(yk - zk) / zk and then the mean is sum(errors in Ak) / (n + 1).
//
//#Examples:
//
//ex_euler(10) should return: 0.026314 (from 0.026314433214799246)
//with
//Y = [1.0,0.9..,0.85..,0.83..,0.83..,0.85..,0.86..,0.88..,0.90..,0.91..,0.93..]
//Z = [1.0,0.9..,0.88..,0.87..,0.87..,0.88..,0.89..,0.90..,0.91..,0.93..,0.94..]
//Relative errors = [0.0,0.02.,0.04..,0.04..,0.04..,0.03..,0.03..,0.02..,0.01..,0.01..,0.01..]
//ex_euler(17) should return:0.015193 (from 0.015193336263370796)`

func main() {
	fmt.Println("THIS IS THE EULER: ", ExEuler(10))
}

func ode(x float64, y float64) float64 {
	return (2 - math.Exp(-4 * x) - 2 * y)
}

func exact(x float64) float64 {
	return (1 * 0.5 * math.Exp(-4 * x) - 0.5 * math.Exp(-2 * x))
}

func ExEuler(nb int) float64 {
	yi := 1.; xi := 0.; h := 1./float64(nb)
	err := 0.
	for i := 1; i <= nb; i++ {
		yi += ode(xi, yi) * h
		xi += h
		var zi = exact(xi)
		err += math.Abs(yi - zi) /zi
	}
	res := err / (float64(nb) + 1. )
	return math.Floor(res * 1e6) / 1e6
}




