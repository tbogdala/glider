// Copyright 2015, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

/*

Package glider is a library that handles 3d collision testing.

Currently only 3d AABB collisions are supported.

*/
package glider

import "math"

func max32(x, y float32) float32 {
	switch {
	case math.IsInf(float64(x), 1) || math.IsInf(float64(y), 1):
		return float32(math.Inf(1))
	case math.IsNaN(float64(x)) || math.IsNaN(float64(y)):
		return float32(math.NaN())
	case x == 0 && x == y:
		if math.Signbit(float64(x)) {
			return y
		}
		return x
	}
	if x > y {
		return x
	}
	return y
}

func min32(x, y float32) float32 {
	switch {
	case math.IsInf(float64(x), -1) || math.IsInf(float64(y), -1):
		return float32(math.Inf(-1))
	case math.IsNaN(float64(x)) || math.IsNaN(float64(y)):
		return float32(math.NaN())
	case x == 0 && x == y:
		if math.Signbit(float64(x)) {
			return x
		}
		return y
	}
	if x < y {
		return x
	}
	return y
}
