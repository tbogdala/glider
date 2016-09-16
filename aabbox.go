// Copyright 2016, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package glider

import (
	mgl "github.com/go-gl/mathgl/mgl32"
)

// AABSquare is a axis aligned sqare shape defined by a minimum and maximum corner.
type AABSquare struct {
	// Min is the corner of the box opposite of Max. (e.g. lower-left corner)
	Min mgl.Vec2

	// Max is the corner of the box opposite of Min. (e.g. top-right corner)
	Max mgl.Vec2

	// Offset is the world-space location of the that can be considered an offset to both Min and Max
	Offset mgl.Vec2

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
func (aabs *AABSquare) IntersectPoint(v *mgl.Vec2) bool {
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
	// and should be the more 'negative' corner (e.g. max={3,3,3} and min={1,1,1} )
	Min mgl.Vec3

	// Max is the corner of the box opposite of Min. (e.g. top-front-right corner)
	// and should be the more 'positive' corner (e.g. max={3,3,3} and min={1,1,1} )
	Max mgl.Vec3

	// Offset is the world-space location of the that can be considered an offset to both Min and Max
	Offset mgl.Vec3

	// Tags provides a way to label an AABB geometry in a custom application
	// (e.g. labelling a collision as "wall" or "floor").
	Tags []string
}

// NewAABBox creates a new AABBox object
func NewAABBox() *AABBox {
	aabb := new(AABBox)
	return aabb
}

// SetOffset changes the offset of the collision object.
func (aabb *AABBox) SetOffset(offset *mgl.Vec3) {
	aabb.Offset = *offset
}

// SetOffset3f changes the offset of the collision object.
func (aabb *AABBox) SetOffset3f(x, y, z float32) {
	aabb.Offset[0] = x
	aabb.Offset[1] = y
	aabb.Offset[2] = z
}

// IntersectPoint tests to see if the point is intersects the AABBox.
func (aabb *AABBox) IntersectPoint(v *mgl.Vec3) bool {
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

// CollideVsAABBox tests to see if the AABBox parameter intersects the AABBox.
func (aabb *AABBox) CollideVsAABBox(b2 *AABBox) int {
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

	if max32(aMinX, bMinX) <= min32(aMaxX, bMaxX) &&
		max32(aMinY, bMinY) <= min32(aMaxY, bMaxY) &&
		max32(aMinZ, bMinZ) <= min32(aMaxZ, bMaxZ) {
		return Intersect
	} else if aMinX >= bMinX && aMaxX <= bMaxX &&
		aMinY >= bMinY && aMaxY <= bMaxY &&
		aMinZ >= bMinZ && aMaxZ <= bMaxZ {
		return Intersect // it's inside
	} else {
		return NoIntersect
	}
}

// CollideVsRay tests to see if a raycast intersects the AABBox.
func (aabb *AABBox) CollideVsRay(ray *CollisionRay) (int, float32) {
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
		return NoIntersect, tmax
	}

	if tmin > tmax {
		return NoIntersect, tmax
	}

	return Intersect, tmin
}

// CollideVsPlane tests to see if the plane is intersects the AABBox.
func (aabb *AABBox) CollideVsPlane(p *Plane) int {
	// implementation based on http://www.lighthouse3d.com/tutorials/view-frustum-culling/
	// and http://www.lighthouse3d.com/tutorials/view-frustum-culling/geometric-approach-testing-boxes-ii/
	min := mgl.Vec3{aabb.Min[0] + aabb.Offset[0], aabb.Min[1] + aabb.Offset[1], aabb.Min[2] + aabb.Offset[2]}
	max := mgl.Vec3{aabb.Max[0] + aabb.Offset[0], aabb.Max[1] + aabb.Offset[1], aabb.Max[2] + aabb.Offset[2]}

	var posCorner, negCorner mgl.Vec3
	if p.Normal[0] >= 0 {
		negCorner[0] = min[0]
		posCorner[0] = max[0]
	} else {
		negCorner[0] = max[0]
		posCorner[0] = min[0]
	}

	if p.Normal[1] >= 0 {
		negCorner[1] = min[1]
		posCorner[1] = max[1]
	} else {
		negCorner[1] = max[1]
		posCorner[1] = min[1]
	}

	if p.Normal[2] >= 0 {
		negCorner[2] = min[2]
		posCorner[2] = max[2]
	} else {
		negCorner[2] = max[2]
		posCorner[2] = min[2]
	}

	// is the positive vertex outside
	if p.Distance(posCorner) < 0 {
		return NoIntersect
	}
	// is the negative vertex outside
	if p.Distance(negCorner) <= 0 {
		return Intersect
	}

	return Intersect // it's inside
}

// CollideVsSphere returns the intersection between an AABB and a sphere.
func (aabb *AABBox) CollideVsSphere(s *Sphere) int {
	min := mgl.Vec3{aabb.Min[0] + aabb.Offset[0], aabb.Min[1] + aabb.Offset[1], aabb.Min[2] + aabb.Offset[2]}
	max := mgl.Vec3{aabb.Max[0] + aabb.Offset[0], aabb.Max[1] + aabb.Offset[1], aabb.Max[2] + aabb.Offset[2]}

	// calc the center of the sphere relative to the box center
	var sphereCenterRelToBox, boxCenter mgl.Vec3
	boxCenter = max.Sub(min)
	boxCenter = boxCenter.Mul(0.5)
	boxCenter = boxCenter.Add(min)

	var offsetSphere mgl.Vec3
	offsetSphere = s.Center.Add(s.Offset)
	sphereCenterRelToBox = offsetSphere.Sub(boxCenter)

	// calculate the point of the cube that is closest to the center of the sphere.
	var boxPoint mgl.Vec3

	// half dimensions
	boxHalfW := (max[0] - min[0]) * 0.5
	boxHalfH := (max[1] - min[1]) * 0.5
	boxHalfD := (max[2] - min[2]) * 0.5

	// calculate x-axis
	if sphereCenterRelToBox[0] < -boxHalfW {
		boxPoint[0] = -boxHalfW
	} else if sphereCenterRelToBox[0] > boxHalfW {
		boxPoint[0] = boxHalfW
	} else {
		boxPoint[0] = sphereCenterRelToBox[0]
	}

	// calculate y-axis
	if sphereCenterRelToBox[1] < -boxHalfH {
		boxPoint[1] = -boxHalfH
	} else if sphereCenterRelToBox[1] > boxHalfH {
		boxPoint[1] = boxHalfH
	} else {
		boxPoint[1] = sphereCenterRelToBox[1]
	}

	// calculate z-axis
	if sphereCenterRelToBox[2] < -boxHalfD {
		boxPoint[2] = -boxHalfD
	} else if sphereCenterRelToBox[2] > boxHalfD {
		boxPoint[2] = boxHalfD
	} else {
		boxPoint[2] = sphereCenterRelToBox[2]
	}

	// use the closest point on the box to get the distance between
	// that and the center of the sphere and see if it's less than
	// the radius.
	var distance mgl.Vec3
	distance = sphereCenterRelToBox.Sub(boxPoint)
	if distance.Dot(distance) <= s.Radius*s.Radius {
		return Intersect
	}

	return NoIntersect
}
