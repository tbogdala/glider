// Copyright 2016, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package glider

// Sphere is defined by a center point and a radius and can be used in collisions
// against AABB.
type Sphere struct {
	// Center is the center point of the sphere, in local space (model-space in 3d graphics)
	Center Vec3

	// Offset is the world-space location of the that can be considered an offset to Center
	Offset Vec3

	// Radius determines the size of the sphere
	Radius float32
}

// SetOffset changes the offset of the collision object.
func (s1 *Sphere) SetOffset(offset *Vec3) {
	s1.Offset = *offset
}

// SetOffset3f changes the offset of the collision object.
func (s1 *Sphere) SetOffset3f(x,y,z float32) {
	s1.Offset[0] = x
	s1.Offset[1] = y
	s1.Offset[2] = z
}

// CollideVsSphere tests a collision between two spheres.
func (s1 *Sphere) CollideVsSphere(s2 *Sphere) int {
    rSquared := s1.Radius + s2.Radius
    rSquared *= rSquared

	var offsetS1, offsetS2 Vec3
	offsetS1.AddInto(&s1.Center, &s1.Offset)
	offsetS2.AddInto(&s2.Center, &s2.Offset)

    var delta Vec3
    delta.SubInto(&offsetS2, &offsetS1)
    distSquared := delta.Dot(&delta)

    if distSquared > rSquared {
        return Outside
    }
    if distSquared - s1.Radius < 0.0 {
        return Inside
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
	var offsetSphere Vec3
	offsetSphere.AddInto(&s1.Center, &s1.Offset)

	var w Vec3
	w.SubInto(&offsetSphere, &ray.Origin)
    rsq := s1.Radius * s1.Radius

    // see if the ray starts inside the sphere
    wsq := w.Dot(&w)
    if wsq < rsq {
        return Intersect, 0.0
    }

    // is the ray start point 'behind' the sphere so that the ray
    // might cast into the sphere?
    proj := w.Dot(&ray.direction)
    if proj < 0.0 {
        return Outside, 0.0
    }

    // we're behind the sphere, so test the length of the difference
    // vs the radius of the sphere
    vsq := ray.direction.Dot(&ray.direction)
    if vsq*wsq - proj*proj <= vsq*rsq {
        return Intersect, 0.0
    }

    return Outside, 0.0
}

// CollideVsPlane tests a collision between a sphere and a plane.
func (s1 *Sphere) CollideVsPlane(p *Plane) int {
	var offsetSphere Vec3
	offsetSphere.AddInto(&s1.Center, &s1.Offset)

    dist := p.Distance(&offsetSphere)
    if dist < 0.0 && -dist > s1.Radius {
        return Outside
    }

    if dist > 0.0 && dist > s1.Radius {
        return Inside
    }

    return Intersect
}
