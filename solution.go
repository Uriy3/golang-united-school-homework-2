package square
	
import (
	"math"
)

type Sides int

const (
	SidesCircle   Sides = 0
	SidesTriangle Sides = 3
	SidesSquare   Sides = 4
)

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	if sidesNum == SidesCircle {
		return math.Pi * math.Pow(sideLen, 2)
	} else if sidesNum == SidesTriangle {
		return (math.Sqrt(3) * math.Pow(sideLen, 2)) / 4
	} else if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2)
	} else {
		return 0
	}
}
