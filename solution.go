package square

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type CustomType int

const (
	SidesCircle   = iota
	SidesTriangle = iota + 2
	SidesSquare   = iota + 2
)

func CalcSquare(sideLen float64, sidesNum CustomType) float64 {

	if sideLen == 0 {
		fmt.Println("You need to specify the length of the side of the figure or the length of the radius of the circle")
		return 0
	}

	switch sidesNum {
	case SidesSquare:
		return math.Pow(sideLen, 2)
	case SidesTriangle:
		return ((math.Pow(sideLen, 2)) * math.Sqrt(3)) / 4
	case SidesCircle:
		return math.Pow(sideLen, 2) * math.Pi
	default:
		fmt.Println("The sides are incorrect. Available options: 4, 3, 0")
		return 0
	}
}
