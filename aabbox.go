// Copyright 2015, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package glider

import "math"

const (
	// Inside means the collision was considered to be inside the object
	Inside = 1

	// Outside means the collision was considered to be outside the object
	Outside = 2

	// Intersect means there was a collision
	Intersect = 4
)

// Vec3 is a 3 dimenional vector
type Vec3 [3]float32

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

// Vec2 is a 2 dimenional vector
type Vec2 [2]float32

// Plane represents an infinite plane defined by a point and its normal.
type Plane struct {
	// Normal is the direction the plane is facing; the normal of the plane.
	Normal Vec3

	// D is the plane constant, considered to be the distance from the origin.
	D float32
}

// NewPlaneFromNormalAndPoint makes a new Plane object based on a normal
// and point in space.
func NewPlaneFromNormalAndPoint(normal, point Vec3) *Plane {
	p := new(Plane)
	p.Normal = normal
	p.D = -(normal.Dot(&point))
	return p
}

// Distance calculates the distance of the plane to the vertex
func (p *Plane) Distance(v *Vec3) float32 {
	return p.D + p.Normal.Dot(v)
}

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

// Sphere is defined by a center point and a radius and can be used in collisions
// against AABB.
type Sphere struct {
	// Center is the center point of the sphere
	Center Vec3

	// Radius determines the size of the sphere
	Radius float32
}

// AABBox is a axis aligned cube shape defined by a minimum and maximum corner.
type AABBox struct {
	// Min is the corner of the box opposite of Max. (e.g. lower-back-left corner)
	// and should be the more 'negative' corner (e.g. max={3,3,3} and min={1,1,1} )
	Min Vec3

	// Max is the corner of the box opposite of Min. (e.g. top-front-right corner)
	// and should be the more 'positive' corner (e.g. max={3,3,3} and min={1,1,1} )
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

// IntersectPlane tests to see if the plane is intersects the AABBox.
func (aabb *AABBox) IntersectPlane(p *Plane) int {
	// implementation based on http://www.lighthouse3d.com/tutorials/view-frustum-culling/
	// and http://www.lighthouse3d.com/tutorials/view-frustum-culling/geometric-approach-testing-boxes-ii/
	min := Vec3{aabb.Min[0] + aabb.Offset[0], aabb.Min[1] + aabb.Offset[1], aabb.Min[2] + aabb.Offset[2]}
	max := Vec3{aabb.Max[0] + aabb.Offset[0], aabb.Max[1] + aabb.Offset[1], aabb.Max[2] + aabb.Offset[2]}

	var posCorner, negCorner Vec3
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
	if p.Distance(&posCorner) < 0 {
		return Outside
	}
	// is the negative vertex outside
	if p.Distance(&negCorner) <= 0 {
		return Intersect
	}

	return Inside
}

// IntersectSphere returns the intersection between an AABB and a sphere. Will
// return Intersect or Outside depending on the result.
func (aabb *AABBox) IntersectSphere(s *Sphere) int {
	min := Vec3{aabb.Min[0] + aabb.Offset[0], aabb.Min[1] + aabb.Offset[1], aabb.Min[2] + aabb.Offset[2]}
	max := Vec3{aabb.Max[0] + aabb.Offset[0], aabb.Max[1] + aabb.Offset[1], aabb.Max[2] + aabb.Offset[2]}

	// calc the center of the sphere relative to the box center
	var sphereCenterRelToBox, boxCenter Vec3
	boxCenter.SubInto(&max, &min)
	boxCenter.MulWith(0.5)
	boxCenter.AddInto(&boxCenter, &min)
	sphereCenterRelToBox.SubInto(&s.Center, &boxCenter)

	// calculate the point of the cube that is closest to the center of the sphere.
	var boxPoint Vec3

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
	var distance Vec3
	distance.SubInto(&sphereCenterRelToBox, &boxPoint)
	if distance.Dot(&distance) <= s.Radius*s.Radius {
		return Intersect
	}

	return Outside
}
