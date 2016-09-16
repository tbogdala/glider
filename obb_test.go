// Copyright 2016, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package glider

import (
	"testing"

	mgl "github.com/go-gl/mathgl/mgl32"
)

func TestOOBCollisionVsSphere(t *testing.T) {
	var sphere Sphere

	obb := NewOBBox()
	obb.HalfSize = mgl.Vec3{1, 1, 1}
	obb.SetOffset(mgl.Vec3{0, 0, 0})
	obb.SetOrientation(mgl.QuatIdent())

	// Sphere {0, 0, 0} | r = 1.0
	sphere = Sphere{Center: mgl.Vec3{0.0, 0.0, 0.0}, Radius: 1.0}
	if obb.CollideVsSphere(&sphere) != Intersect {
		t.Error("OBBox.CollideVsSphere() indicated a sphere didn't collide that should have.")
	}

	// Sphere {2, 0, 0} | r = 1.0
	sphere = Sphere{Center: mgl.Vec3{0.0, 0.0, 0.0}, Radius: 1.0}
	if obb.CollideVsSphere(&sphere) != Intersect {
		t.Error("OBBox.CollideVsSphere() indicated a sphere didn't collide that should have.")
	}

	// Sphere {2.1, 0, 0} | r = 1.0
	sphere = Sphere{Center: mgl.Vec3{2.1, 0.0, 0.0}, Radius: 1.0}
	if obb.CollideVsSphere(&sphere) != NoIntersect {
		t.Error("OBBox.CollideVsSphere() indicated a sphere collided that should not have.")
	}
	// Sphere {0, 5, 0} | r = 1.0
	sphere = Sphere{Center: mgl.Vec3{0.0, 5.0, 0.0}, Radius: 1.0}
	if obb.CollideVsSphere(&sphere) != NoIntersect {
		t.Error("OBBox.CollideVsSphere() indicated a sphere collided that should not have.")
	}

	// ---------------------------------------------------------------------------
	// ROTATE OBB 45 DEG
	q := mgl.QuatRotate(mgl.DegToRad(45.0), mgl.Vec3{0, 0, 1})
	obb.SetOrientation(q)

	// Sphere {2.1, 0, 0} | r = 1.0
	sphere = Sphere{Center: mgl.Vec3{2.1, 0.0, 0.0}, Radius: 1.0}
	if obb.CollideVsSphere(&sphere) != Intersect {
		t.Error("OBBox.CollideVsSphere() indicated a sphere didn't collide that should have.")
	}

	// Sphere {2.5, 0, 0} | r = 1.0
	sphere = Sphere{Center: mgl.Vec3{2.5, 0.0, 0.0}, Radius: 1.0}
	if obb.CollideVsSphere(&sphere) != NoIntersect {
		t.Error("OBBox.CollideVsSphere() indicated a sphere collided that should not have.")
	}

}
