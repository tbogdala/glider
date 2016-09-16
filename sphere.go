// Copyright 2016, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package glider

import (
	mgl "github.com/go-gl/mathgl/mgl32"
)

// Sphere is defined by a center point and a radius and can be used in collisions
// against AABB.
type Sphere struct {
	// Center is the center point of the sphere, in local space (model-space in 3d graphics)
	Center mgl.Vec3

	// Offset is the world-space location of the that can be considered an offset to Center
	Offset mgl.Vec3

	// Radius determines the size of the sphere
	Radius float32
}

// SetOffset changes the offset of the collision object.
func (s1 *Sphere) SetOffset(offset *mgl.Vec3) {
	s1.Offset = *offset
}

// SetOffset3f changes the offset of the collision object.
func (s1 *Sphere) SetOffset3f(x, y, z float32) {
	s1.Offset[0] = x
	s1.Offset[1] = y
	s1.Offset[2] = z
}

// CollideVsSphere tests a collision between two spheres.
func (s1 *Sphere) CollideVsSphere(s2 *Sphere) int {
	rSquared := s1.Radius + s2.Radius
	rSquared *= rSquared

	var offsetS1, offsetS2 mgl.Vec3
	offsetS1 = s1.Center.Add(s1.Offset)
	offsetS2 = s2.Center.Add(s2.Offset)

	var delta mgl.Vec3
	delta = offsetS2.Sub(offsetS1)
	distSquared := delta.Dot(delta)

	if distSquared > rSquared {
		return NoIntersect
	}

	return Intersect
}

// CollideVsAABBox tests a collision between a sphere and an AABBox.
func (s1 *Sphere) CollideVsAABBox(b *AABBox) int {
	return b.CollideVsSphere(s1)
}

// CollideVsRay tests a collision between a sphere and a ray.
// FIXME: sphere's don't return a distance at present
func (s1 *Sphere) CollideVsRay(ray *CollisionRay) (int, float32) {
	var offsetSphere mgl.Vec3
	offsetSphere = s1.Center.Add(s1.Offset)

	var w mgl.Vec3
	w = offsetSphere.Sub(ray.Origin)
	rsq := s1.Radius * s1.Radius

	// see if the ray starts inside the sphere
	wsq := w.Dot(w)
	if wsq < rsq {
		return Intersect, 0.0
	}

	// is the ray start point 'behind' the sphere so that the ray
	// might cast into the sphere?
	proj := w.Dot(ray.direction)
	if proj < 0.0 {
		return NoIntersect, 0.0
	}

	// we're behind the sphere, so test the length of the difference
	// vs the radius of the sphere
	vsq := ray.direction.Dot(ray.direction)
	if vsq*wsq-proj*proj <= vsq*rsq {
		return Intersect, 0.0
	}

	return NoIntersect, 0.0
}

// CollideVsPlane tests a collision between a sphere and a plane.
func (s1 *Sphere) CollideVsPlane(p *Plane) int {
	var offsetSphere mgl.Vec3
	offsetSphere = s1.Center.Add(s1.Offset)

	dist := p.Distance(offsetSphere)
	if dist < 0.0 && -dist > s1.Radius {
		return NoIntersect
	}

	return Intersect
}
