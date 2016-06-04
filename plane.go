// Copyright 2016, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package glider

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
