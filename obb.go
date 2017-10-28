// Copyright 2016, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package glider

import (
	"math"

	mgl "github.com/go-gl/mathgl/mgl32"
)

// OBBox is a oriented bounding box shape defined by a center ('Offset') and
// a HalfSize. HalfSize is used instead of Min/Max corners because the math is
// quicker.
type OBBox struct {
	// HalfSize holds the cube's half-sizes along each of its local axes.
	HalfSize mgl.Vec3

	// Offset is the world-space center location of cube.
	Offset mgl.Vec3

	orientation mgl.Quat

	transform mgl.Mat4

	// Tags provides a way to label an OBB geometry in a custom application
	// (e.g. labelling a collision as "wall" or "floor").
	Tags []string
}

// NewOBBox creates a new OBBox object
func NewOBBox() *OBBox {
	obb := new(OBBox)
	obb.transform = mgl.Ident4()
	return obb
}

// transformInverse transforms the vector by the transformational
// inverse of this matrix.
// NOTE: will not work on matrixes with scale or shears.
func transformInverse(m *mgl.Mat4, v *mgl.Vec3) mgl.Vec3 {
	tmp := *v
	tmp[0] -= m[12]
	tmp[1] -= m[13]
	tmp[2] -= m[14]

	return mgl.Vec3{
		tmp[0]*m[0] + tmp[1]*m[1] + tmp[2]*m[2],
		tmp[0]*m[4] + tmp[1]*m[5] + tmp[2]*m[6],
		tmp[0]*m[8] + tmp[1]*m[9] + tmp[2]*m[10],
	}
}

func fabs32(a float32) float32 {
	return float32(math.Abs(float64(a)))
}

// SetOffset changes the offset of the collision object.
func (obb *OBBox) SetOffset(offset mgl.Vec3) {
	obb.Offset = offset
	obb.syncOffset()
}

// GetOffset returns the offset of the collision object.
func (obb *OBBox) GetOffset() mgl.Vec3 {
	return obb.Offset
}

// SetOffset3f changes the offset of the collision object.
func (obb *OBBox) SetOffset3f(x, y, z float32) {
	obb.Offset[0] = x
	obb.Offset[1] = y
	obb.Offset[2] = z
	obb.syncOffset()
}

// SetOrientation sets the rotation of the box.
func (obb *OBBox) SetOrientation(q mgl.Quat) {
	obb.orientation = q
	obb.transform = q.Mat4()
	obb.syncOffset()
}

func (obb *OBBox) syncOffset() {
	obb.transform[12] = obb.Offset[0]
	obb.transform[13] = obb.Offset[1]
	obb.transform[14] = obb.Offset[2]
}

// CollideVsSphere tests an OBBox vs Sphere collision.
func (obb *OBBox) CollideVsSphere(sphere *Sphere) int {
	// transform the center of the sphere into cube coordinates
	position := sphere.Offset.Add(sphere.Center)
	relCenter := transformInverse(&obb.transform, &position)

	// check to see if we can exclude contact
	if fabs32(relCenter[0])-sphere.Radius > obb.HalfSize[0] ||
		fabs32(relCenter[1])-sphere.Radius > obb.HalfSize[1] ||
		fabs32(relCenter[2])-sphere.Radius > obb.HalfSize[2] {
		return NoIntersect
	}

	var closestPoint mgl.Vec3

	// clamp the coordinates to the box
	for i := 0; i < 3; i++ {
		dist := relCenter[i]
		if dist > obb.HalfSize[i] {
			dist = obb.HalfSize[i]
		} else if dist < -obb.HalfSize[i] {
			dist = -obb.HalfSize[i]
		}
		closestPoint[i] = dist
	}

	// check to see if we're in contact
	distCheck := closestPoint.Sub(relCenter)
	dist := distCheck.Dot(distCheck)
	if dist > sphere.Radius*sphere.Radius {
		return NoIntersect
	}

	return Intersect
}
