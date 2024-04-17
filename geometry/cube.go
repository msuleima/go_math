package geometry

import (
	"errors"
	"math"
)

func CubeVolume(n float64) (float64, error) {
	if n == 0 {
		return 0, errors.New("zero length edge is not allowed")
	}

	return math.Pow(n, 3), nil
}