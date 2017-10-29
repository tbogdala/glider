// Copyright 2016, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

/*

Package glider is a library that handles 3d collision testing.

*/
package glider

import (
	"math"

	mgl "github.com/go-gl/mathgl/mgl32"
)

const (
	// NoIntersect means there was no collision detect.
	NoIntersect = 0

	// Intersect means there was a collision
	Intersect = 1
)

// Collider is an interface for objects that con collide with other
// collision primitives.
type Collider interface {
	CollideVsSphere(sphere *Sphere) int
	CollideVsAABBox(box *AABBox) int
	CollideVsPlane(plane *Plane) int
	CollideVsRay(ray *CollisionRay) (int, float32)
	SetOffset(offset *mgl.Vec3)
	SetOffset3f(x, y, z float32)
	GetOffset() mgl.Vec3
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

	return NoIntersect
}

// CollisionRay represents a simple ray for casting in collision tests.
type CollisionRay struct {
	// Origin is the start of the ray
	Origin mgl.Vec3

	// direction is the unit vector representing the direction of the ray
	direction mgl.Vec3

	// a cached value used in raycasting
	directionFraction mgl.Vec3
}

// SetDirection sets the direction of the collision ray. Will be normalized
// and have some math cached as well.
func (cr *CollisionRay) SetDirection(d mgl.Vec3) {
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
func (cr *CollisionRay) GetDirection() mgl.Vec3 {
	return cr.direction
}

// IntersectTri will return true if the ray intersects a triangle
// defined by three points and the distance it intersects at.
func (cr *CollisionRay) IntersectTri(p0, p1, p2 mgl.Vec3) (intersected bool, dist float32) {
	u := p1.Sub(p0)
	v := p2.Sub(p0)

	var normal mgl.Vec3
	normal[0] = u[1]*v[2] - u[2]*v[1]
	normal[1] = u[2]*v[0] - u[0]*v[2]
	normal[2] = u[0]*v[1] - u[1]*v[0]

	denom := normal.Dot(cr.direction)
	if math.Abs(float64(denom)) > 0.0001 {
		t := p0.Sub(cr.Origin).Dot(normal) / denom
		if t >= 0 {
			return true, t
		}
	}

	return false, 0.0
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
