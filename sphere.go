// Copyright 2016, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package glider

// Sphere is defined by a center point and a radius and can be used in collisions
// against AABB.
type Sphere struct {
	// Center is the center point of the sphere
	Center Vec3

	// Radius determines the size of the sphere
	Radius float32
}


// IntersectSphere tests a collision between two spheres.
func (s1 *Sphere) IntersectSphere(s2 *Sphere) int {
    rSquared := s1.Radius + s2.Radius
    rSquared *= rSquared

    var delta Vec3
    delta.SubInto(&s2.Center, &s1.Center)
    distSquared := delta.Dot(&delta)

    if distSquared > rSquared {
        return Outside
    }
    if distSquared - s1.Radius < 0.0 {
        return Inside
    }

    return Intersect
}

// IntersectAABBox tests a collision between a sphere and an AABBox.
func (s1 *Sphere) IntersectAABBox(b *AABBox) int {
    result := b.IntersectSphere(s1)
    if result == Intersect {
        return Intersect
    } else if result == Inside {
        return Outside
    } else {
        return Inside
    }
}

// IntersectRay tests a collision between a sphere and a ray.
func (s1 *Sphere) IntersectRay(ray *CollisionRay) bool {
    var w Vec3
    w.SubInto(&s1.Center, &ray.Origin)
    rsq := s1.Radius * s1.Radius

    // see if the ray starts inside the sphere
    wsq := w.Dot(&w)
    if wsq < rsq {
        return true
    }

    // is the ray start point 'behind' the sphere so that the ray
    // might cast into the sphere?
    proj := w.Dot(&ray.direction)
    if proj < 0.0 {
        return false
    }

    // we're behind the sphere, so test the length of the difference
    // vs the radius of the sphere
    vsq := ray.direction.Dot(&ray.direction)
    if vsq*wsq - proj*proj <= vsq*rsq {
        return true
    }

    return false
}
