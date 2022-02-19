package main

import "math"

Type myType int

func CalcSquare(sideLen float64, sidesNum myType) (res float64) {

	const SidesTriangle myType = 3
	const SidesSquare myType = 4
	const SidesCircle myType = 0
	if SidesTriangle == sidesNum {
		res = (math.Pow(sideLen, 2.0) * math.Sqrt(3.0))/4
	} else if SidesSquare == sidesNum {
		res = math.Pow(sideLen, 2.0)
	} else if SidesCircle == sidesNum {
		res = math.Pi * math.Pow(sideLen, 2.0)
	}
	return
