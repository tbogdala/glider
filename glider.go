// Copyright 2016, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

/*

Package glider is a library that handles 3d collision testing.

Currently only 3d AABB collisions are supported.

*/
package glider

import "math"

// Vec2 is a 2 dimenional vector
type Vec2 [2]float32

// Vec3 is a 3 dimenional vector
type Vec3 [3]float32

const (
	// Inside means the collision was considered to be inside the object
	Inside = 1

	// Outside means the collision was considered to be outside the object
	Outside = 2

	// Intersect means there was a collision
	Intersect = 4
)

// Collider is an interface for objects that con collide with other
// collision primitives.
type Collider interface {
	CollideVsSphere(sphere *Sphere) int
	CollideVsAABBox(box *AABBox) int
	CollideVsPlane(plane *Plane) int
	CollideVsRay(ray *CollisionRay) (int, float32)
	SetOffset(offset *Vec3)
	SetOffset3f(x,y,z float32)
}

// Collide tests two objects that are Colliders and returns the collision test result.
// NOTE: currently this supports cubes and spheres.
// FIXME: planes and rays are not tested here
func Collide(c1 Collider, c2 Collider) int {
	targetBox, okay := c2.(*AABBox)
	if okay {
		return c1.CollideVsAABBox(targetBox)
	}

	targetSphere, okay := c2.(*Sphere)
	if okay {
		return c1.CollideVsSphere(targetSphere)
	}

	return Outside
}

// Dot calculates the dot product between two vectors and returns the scalar result.
func (v1 *Vec3) Dot(v2 *Vec3) float32 {
	return (v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2])
}

// SubInto performs v2-v3 and puts the result into v1, the calling object.
func (v1 *Vec3) SubInto(v2 *Vec3, v3 *Vec3) {
	v1[0] = v2[0] - v3[0]
	v1[1] = v2[1] - v3[1]
	v1[2] = v2[2] - v3[2]
}

// AddInto performs v2+v3 and puts the result into v1, the calling object.
func (v1 *Vec3) AddInto(v2 *Vec3, v3 *Vec3) {
	v1[0] = v2[0] + v3[0]
	v1[1] = v2[1] + v3[1]
	v1[2] = v2[2] + v3[2]
}

// MulWith multiplies a vector with a scalar value
func (v1 *Vec3) MulWith(f float32) {
	v1[0] = v1[0] * f
	v1[1] = v1[1] * f
	v1[2] = v1[2] * f
}


// CollisionRay represents a simple ray for casting in collision tests.
type CollisionRay struct {
	// Origin is the start of the ray
	Origin Vec3

	// direction is the unit vector representing the direction of the ray
	direction Vec3

	// a cached value used in raycasting
	directionFraction Vec3
}

// SetDirection sets the direction of the collision ray. Will be normalized
// and have some math cached as well.
func (cr *CollisionRay) SetDirection(d Vec3) {
	// normalize the direction vector
	dLen := float32(math.Sqrt(float64(d[0]*d[0] + d[1]*d[1] + d[2]*d[2])))
	l := 1.0 / dLen
	cr.direction[0] = d[0] * l
	cr.direction[1] = d[1] * l
	cr.direction[2] = d[2] * l

	// cache some math calculations
	cr.directionFraction[0] = 1.0 / cr.direction[0]
	cr.directionFraction[1] = 1.0 / cr.direction[1]
	cr.directionFraction[2] = 1.0 / cr.direction[2]
}

// GetDirection gets the direction of the collision ray.
func (cr *CollisionRay) GetDirection() Vec3 {
	return cr.direction
}

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
