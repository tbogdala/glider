// Copyright 2015, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package glider

import (
	"math"
)

// Vec3 is a 3 dimenional vector
type Vec3 [3]float32

// Vec2 is a 2 dimenional vector
type Vec2 [2]float32

// AABSquare is a axis aligned sqare shape defined by a minimum and maximum corner.
type AABSquare struct {
	// Min is the corner of the box opposite of Max. (e.g. lower-left corner)
	Min Vec2

	// Max is the corner of the box opposite of Min. (e.g. top-right corner)
	Max Vec2

	// Offset is the world-space location of the that can be considered an offset to both Min and Max
	Offset Vec2

	// Tags provides a way to label an AABB geometry in a custom application
	// (e.g. labelling a collision as "wall" or "floor").
	Tags []string
}

// NewAABSquare creates a new AABSquare object
func NewAABSquare() *AABSquare {
	aabs := new(AABSquare)
	return aabs
}

// IntersectPoint tests to see if the point is intersects the AABSquare.
func (aabs *AABSquare) IntersectPoint(v *Vec2) bool {
	if v[0] < aabs.Offset[0]+aabs.Min[0] || v[0] > aabs.Offset[0]+aabs.Max[0] {
		return false
	}
	if v[1] < aabs.Offset[1]+aabs.Min[1] || v[1] > aabs.Offset[1]+aabs.Max[1] {
		return false
	}
	return true
}

// AABBox is a axis aligned cube shape defined by a minimum and maximum corner.
type AABBox struct {
	// Min is the corner of the box opposite of Max. (e.g. lower-back-left corner)
	Min Vec3

	// Max is the corner of the box opposite of Min. (e.g. top-front-right corner)
	Max Vec3

	// Offset is the world-space location of the that can be considered an offset to both Min and Max
	Offset Vec3

	// Tags provides a way to label an AABB geometry in a custom application
	// (e.g. labelling a collision as "wall" or "floor").
	Tags []string
}

// NewAABBox creates a new AABBox object
func NewAABBox() *AABBox {
	aabb := new(AABBox)
	return aabb
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

// IntersectPoint tests to see if the point is intersects the AABBox.
func (aabb *AABBox) IntersectPoint(v *Vec3) bool {
	aMinX := aabb.Min[0] + aabb.Offset[0]
	aMinY := aabb.Min[1] + aabb.Offset[1]
	aMinZ := aabb.Min[2] + aabb.Offset[2]
	aMaxX := aabb.Max[0] + aabb.Offset[0]
	aMaxY := aabb.Max[1] + aabb.Offset[1]
	aMaxZ := aabb.Max[2] + aabb.Offset[2]

	if v[0] < aMinX || v[0] > aMaxX {
		return false
	}
	if v[1] < aMinY || v[1] > aMaxY {
		return false
	}
	if v[2] < aMinZ || v[2] > aMaxZ {
		return false
	}
	return true
}

// IntersectBox tests to see if the AABBox parameter intersects the AABBox.
func (aabb *AABBox) IntersectBox(b2 *AABBox) bool {
	aMinX := aabb.Min[0] + aabb.Offset[0]
	aMinY := aabb.Min[1] + aabb.Offset[1]
	aMinZ := aabb.Min[2] + aabb.Offset[2]
	aMaxX := aabb.Max[0] + aabb.Offset[0]
	aMaxY := aabb.Max[1] + aabb.Offset[1]
	aMaxZ := aabb.Max[2] + aabb.Offset[2]

	bMinX := b2.Min[0] + b2.Offset[0]
	bMinY := b2.Min[1] + b2.Offset[1]
	bMinZ := b2.Min[2] + b2.Offset[2]
	bMaxX := b2.Max[0] + b2.Offset[0]
	bMaxY := b2.Max[1] + b2.Offset[1]
	bMaxZ := b2.Max[2] + b2.Offset[2]

	return (max32(aMinX, bMinX) <= min32(aMaxX, bMaxX) &&
		max32(aMinY, bMinY) <= min32(aMaxY, bMaxY) &&
		max32(aMinZ, bMinZ) <= min32(aMaxZ, bMaxZ))
}

// IntersectRay tests to see if a raycast intersects the AABBox.
// Returns intersection status and the length of the ray until intersection
func (aabb *AABBox) IntersectRay(ray *CollisionRay) (bool, float32) {
	aMinX := aabb.Min[0] + aabb.Offset[0]
	aMinY := aabb.Min[1] + aabb.Offset[1]
	aMinZ := aabb.Min[2] + aabb.Offset[2]
	aMaxX := aabb.Max[0] + aabb.Offset[0]
	aMaxY := aabb.Max[1] + aabb.Offset[1]
	aMaxZ := aabb.Max[2] + aabb.Offset[2]

	t1 := (aMinX - ray.Origin[0]) * ray.directionFraction[0]
	t3 := (aMinY - ray.Origin[1]) * ray.directionFraction[1]
	t5 := (aMinZ - ray.Origin[2]) * ray.directionFraction[2]

	t2 := (aMaxX - ray.Origin[0]) * ray.directionFraction[0]
	t4 := (aMaxY - ray.Origin[1]) * ray.directionFraction[1]
	t6 := (aMaxZ - ray.Origin[2]) * ray.directionFraction[2]

	tmin := max32(max32(min32(t1, t2), min32(t3, t4)), min32(t5, t6))
	tmax := min32(min32(max32(t1, t2), max32(t3, t4)), max32(t5, t6))

	// if tmax < 0, ray is intersecting the box, but the whole AABB is behind
	if tmax < 0 {
		return false, tmax
	}

	if tmin > tmax {
		return false, tmax
	}

	return true, tmin
}
