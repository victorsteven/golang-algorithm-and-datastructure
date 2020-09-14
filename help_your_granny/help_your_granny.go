package help_your_granny

import (
	. "math"
)

//Your granny, who lives in town X0, has friends. These friends are given in an array, for example: array of friends is
//
//[ "A1", "A2", "A3", "A4", "A5" ].
//The order of friends is this array must not be changed since this order gives the order in which they will be visited.
//
//These friends inhabit towns and you have an array with friends and towns, for example:
//
//[ ["A1", "X1"], ["A2", "X2"], ["A3", "X3"], ["A4", "X4"] ]
//or
//[ ("A1", "X1"), ("A2", "X2"), ("A3", "X3"), ("A4", "X4") ]
//or
//(C)
//{"A1", "X1", "A2", "X2", "A3", "X3", "A4", "X4"}
//which means A1 is in town X1, A2 in town X2... It can happen that we don't know the town of one of the friends.
//
//Your granny wants to visit her friends and to know how many miles she will have to travel.
//
//You will make the circuit that permits her to visit her friends. For example here the circuit will contain:
//
//X0, X1, X2, X3, X4, X0
//and you must compute the total distance
//
//X0X1 + X1X2 + .. + X4X0.
//For the distance, fortunately, you have a map (and a hashmap) that gives each distance X0X1, X0X2 and so on. For example:
//
//[ ["X1", 100.0], ["X2", 200.0], ["X3", 250.0], ["X4", 300.0] ]
//or
//Map("X1" -> 100.0, "X2" -> 200.0, "X3" -> 250.0, "X4" -> 300.0)
//or (Coffeescript, Javascript)
//['X1',100.0, 'X2',200.0, 'X3',250.0, 'X4',300.0 ]
//or
//(C)
//{"X1", "100.0", "X2", "200.0", "X3", "250.0", "X4", "300.0"}
//which means that X1 is at 100.0 miles from X0, X2 at 200.0 miles from X0, etc...
//
//More fortunately (it's not real life, it's a story...), the towns X0, X1, ..Xn are placed in the following manner:
//
//X0X1X2 is a right triangle with the right angle in X1, X0X2X3 is a right triangle with the right angle in X2, etc...
//
//If a town Xi is not visited you will suppose that the triangle
//
//X0Xi-1Xi+1 is still a right triangle.
//
//(Ref: https://en.wikipedia.org/wiki/Pythagoras#Pythagorean_theorem)
//
//Task
//Can you help your granny and give her the distance to travel?
//
//Notes
//If you have some difficulty to see the tour I made a non terrific but maybe useful drawing:
//


func Tour(arrFriends []string, ftwns map[string]string, h map[string]float64) int {

	result := 0.0
	toFrom := h[ftwns[arrFriends[0]]]
	prev := 0.0

	for i := 1; i < len(arrFriends); i++ {
		if v, ok := h[ftwns[arrFriends[i]]]; ok {
			prev = h[ftwns[arrFriends[i - 1]]]

			result += Sqrt((Pow(v, 2)) - Pow(prev, 2))
			println(result)
		} else {
			break
		}
		prev = h[ftwns[arrFriends[i]]]
	}
	result += prev

	return int(Floor(result)) + int(Floor(toFrom))
}

//func Tour(arrFriends []string, ftwns map[string]string, h map[string]float64) int {
//	var t []string
//	for _, f := range arrFriends {
//		tt, ok := ftwns[f]
//		if ok {
//			t = append(t, tt)
//		}
//	}
//	var d []float64
//	for _, tt := range t {
//		dd, ok := h[tt]
//		if ok {
//			d = append(d, dd)
//		}
//	}
//	var s float64 = 0.0
//	for i := 0; i < len(d) - 1; i++ {
//		s += math.Sqrt(d[i + 1] * d[i + 1] - d[i] * d[i])
//	}
//	s += d[0] + d[len(d) - 1]
//	return int(s)
//}


//func Tour(arrFriends []string, ftwns map[string]string, h map[string]float64) int {
//	k := h[ftwns[arrFriends[0]]]
//	out := k
//	for _, v := range arrFriends[1:] {
//		if tmp := h[ftwns[v]]; tmp != 0 {
//			out += math.Sqrt(tmp*tmp - k*k)
//			k = tmp
//		}
//	}
//	out += k
//	return int(out)
//}