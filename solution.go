package square

import "math"

type CustomType = int

const Pi = 3.14159
const SidesCircle = 0
const SidesTriangle = 3
const SidesSquare = 4

func CalcSquare(sideLen float64, sidesNum CustomType) float64 {
	switch sidesNum {
	case SidesSquare:
		return sideLen * sideLen
	case SidesTriangle:
		return (math.Sqrt(3) / 4) * sideLen * sideLen
	case SidesCircle:
		return Pi * sideLen * sideLen
	default:
		return 0
	}
}
