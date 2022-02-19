package square
	
	import "math"
	
	type CustomType = int
	
	const Pi = 3.14159
	const SidesCircle = 0
	const SidesTriangle = 3
	const SidesSquare = 4
	
	// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
	
	// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
	// it's like:
	// CalcSquare(10.0, SidesTriangle)
	// CalcSquare(10.0, SidesSquare)
	// CalcSquare(10.0, SidesCircle)
	//sideLen float64, sidesNum CustomType
	
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
