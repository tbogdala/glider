// Copyright 2015, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package glider

import (
	"testing"

	mgl "github.com/go-gl/mathgl/mgl32"
)

func TestAABSquareCollisionVsPoint(t *testing.T) {
	var s1 AABSquare
	var p1 mgl.Vec2

	s1.Min = mgl.Vec2{0.0, 0.0}
	s1.Max = mgl.Vec2{1.0, 1.0}
	s1.Offset = mgl.Vec2{0.0, 0.0}

	p1 = mgl.Vec2{0.5, 0.5}
	if s1.IntersectPoint(&p1) == false {
		t.Error("AABSquare.IntersectPoint() indicated false with a unit square and a point that intersect.")
	}

	s1.Offset = mgl.Vec2{10.0, 5.0}
	p1 = mgl.Vec2{0.5, 0.5}
	if s1.IntersectPoint(&p1) == true {
		t.Error("AABSquare.IntersectPoint() indicated false with a unit square and a point that intersect.")
	}

	p1 = mgl.Vec2{10.5, 5.5}
	if s1.IntersectPoint(&p1) == false {
		t.Error("AABSquare.IntersectPoint() indicated false with a unit square and a point that intersect.")
	}

}

func TestAABBoxNoCollision(t *testing.T) {
	var b1, b2 AABBox

	b1.Min = mgl.Vec3{0.0, 0.0, 0.0}
	b1.Max = mgl.Vec3{1.0, 1.0, 1.0}

	b2.Min = mgl.Vec3{2.0, 2.0, 2.0}
	b2.Max = mgl.Vec3{3.0, 3.0, 3.0}

	if b1.CollideVsAABBox(&b2) == Intersect {
		t.Error("AABBox.IntersectBox() indicated true with two unit cubes that don't intersect.")
	}
}

func TestAABBoxCollision(t *testing.T) {
	var b1, b2 AABBox

	b1.Min = mgl.Vec3{0.0, 0.0, 0.0}
	b1.Max = mgl.Vec3{1.0, 1.0, 1.0}

	b2.Min = mgl.Vec3{0.5, 0.5, 0.5}
	b2.Max = mgl.Vec3{3.0, 3.0, 3.0}

	if b1.CollideVsAABBox(&b2) == NoIntersect {
		t.Error("AABBox.IntersectBox() indicated false with two unit cubes that intersect.")
	}
}

func TestAABBoxSiblingCollision(t *testing.T) {
	var b1, b2 AABBox

	b1.Min = mgl.Vec3{0.0, 0.0, 0.0}
	b1.Max = mgl.Vec3{1.0, 1.0, 1.0}

	b2.Min = mgl.Vec3{1.0, 1.0, 1.0}
	b2.Max = mgl.Vec3{3.0, 3.0, 3.0}

	if b1.CollideVsAABBox(&b2) == NoIntersect {
		t.Error("AABBox.IntersectBox() indicated false with two unit cubes that share an edge.")
	}
}

func TestAABBoxCollisionVsPoint(t *testing.T) {
	var b1 AABBox
	var p1 mgl.Vec3

	b1.Min = mgl.Vec3{0.0, 0.0, 0.0}
	b1.Max = mgl.Vec3{1.0, 1.0, 1.0}
	b1.Offset = mgl.Vec3{0.0, 0.0, 0.0}

	p1 = mgl.Vec3{0.5, 0.5, 0.5}
	if b1.IntersectPoint(&p1) == false {
		t.Error("AABBox.IntersectPoint() indicated false with a unit cube and a point that intersect.")
	}

	b1.Offset = mgl.Vec3{3.0, 3.0, 3.0}
	if b1.IntersectPoint(&p1) == true {
		t.Error("AABBox.IntersectPoint() indicated false with a unit cube and a point that intersect.")
	}

	p1 = mgl.Vec3{3.5, 3.5, 3.5}
	if b1.IntersectPoint(&p1) == false {
		t.Error("AABBox.IntersectPoint() indicated false with a unit cube and a point that intersect.")
	}
}

func TestAABBoxCollisionVsEdgePoint(t *testing.T) {
	var b1 AABBox
	var p1 mgl.Vec3

	b1.Min = mgl.Vec3{0.0, 0.0, 0.0}
	b1.Max = mgl.Vec3{1.0, 1.0, 1.0}

	p1 = mgl.Vec3{1.0, 1.0, 1.0}

	if b1.IntersectPoint(&p1) == false {
		t.Error("AABBox.IntersectPoint() indicated false with a unit cube and a point along an edge.")
	}
}

func TestAABBoxCollisionVsRay(t *testing.T) {
	var b1 AABBox
	var r1 CollisionRay

	b1.Min = mgl.Vec3{-1.0, -1.0, -1.0}
	b1.Max = mgl.Vec3{1.0, 1.0, 1.0}

	// cast at the center
	r1.Origin = mgl.Vec3{5.0, 0.0, 0.0}
	r1.SetDirection(mgl.Vec3{-1.0, 0.0, 0.0})

	intersect, _ := b1.CollideVsRay(&r1)
	if intersect == NoIntersect {
		t.Error("AABBox.IntersectRay() indicated false with a ray pointed at it's center.")
	}

	// cast away from it
	r1.Origin = mgl.Vec3{5.0, 0.0, 0.0}
	r1.SetDirection(mgl.Vec3{1.0, 0.0, 0.0})

	intersect, _ = b1.CollideVsRay(&r1)
	if intersect == Intersect {
		t.Error("AABBox.IntersectRay() indicated true with a ray pointed away from it.")
	}

	// cast at the edge
	r1.Origin = mgl.Vec3{10.0, 1.0, 1.0}
	r1.SetDirection(mgl.Vec3{-1.0, 0.0, 0.0})

	intersect, _ = b1.CollideVsRay(&r1)
	if intersect == NoIntersect {
		t.Error("AABBox.IntersectRay() indicated false with a ray pointed at it's edge.")
	}

	// cast from inside
	r1.Origin = mgl.Vec3{0.0, 0.0, 0.0}
	r1.SetDirection(mgl.Vec3{1.0, 1.0, 1.0})

	intersect, _ = b1.CollideVsRay(&r1)
	if intersect == NoIntersect {
		t.Error("AABBox.IntersectRay() indicated false with a ray starting at the center of the box.")
	}
}

func TestAABBoxCollisionVsRay2(t *testing.T) {
	var b1 AABBox
	var r1 CollisionRay

	b1.Min = mgl.Vec3{-1.0, -1.0, -1.0}
	b1.Max = mgl.Vec3{1.0, 1.0, 1.0}
	b1.Offset = mgl.Vec3{10.0, 0.0, 0.0}

	// cast at the center
	r1.Origin = mgl.Vec3{15.0, 0.0, 0.0}
	r1.SetDirection(mgl.Vec3{-1.0, 0.0, 0.0})

	intersect, _ := b1.CollideVsRay(&r1)
	if intersect == NoIntersect {
		t.Error("AABBox.IntersectRay2() indicated false with a ray pointed at it's center.")
	}

	// cast away from it
	r1.Origin = mgl.Vec3{15.0, 0.0, 0.0}
	r1.SetDirection(mgl.Vec3{1.0, 0.0, 0.0})

	intersect, _ = b1.CollideVsRay(&r1)
	if intersect == Intersect {
		t.Error("AABBox.IntersectRay2() indicated true with a ray pointed away from it.")
	}

	// cast at the edge
	r1.Origin = mgl.Vec3{10.0, 1.0, 1.0}
	r1.SetDirection(mgl.Vec3{-1.0, 0.0, 0.0})

	intersect, _ = b1.CollideVsRay(&r1)
	if intersect == NoIntersect {
		t.Error("AABBox.IntersectRay2() indicated false with a ray pointed at it's edge.")
	}

	// cast from inside
	r1.Origin = mgl.Vec3{10.0, 0.0, 0.0}
	r1.SetDirection(mgl.Vec3{1.0, 1.0, 1.0})

	intersect, _ = b1.CollideVsRay(&r1)
	if intersect == NoIntersect {
		t.Error("AABBox.IntersectRay2() indicated false with a ray starting at the center of the box.")
	}
}

func TestAABBoxCollisionVsPlane(t *testing.T) {
	var b1 AABBox
	var p *Plane

	b1.Min = mgl.Vec3{-10.0, -10.0, -10.0}
	b1.Max = mgl.Vec3{10.0, 10.0, 10.0}
	b1.Offset = mgl.Vec3{0.0, 0.0, 0.0}

	// Plane @ {0, 0, 0}   Normal---> {1, 0, 0}
	planeNormal := mgl.Vec3{1.0, 0.0, 0.0}
	p = NewPlaneFromNormalAndPoint(planeNormal, mgl.Vec3{0, 0, 0})
	if b1.CollideVsPlane(p) != Intersect {
		t.Errorf("AABBox.InstersectPlane() indicated a box didn't intersect that should have.")
	}

	// Plane @ {20, 0, 0}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, mgl.Vec3{20, 0, 0})
	if b1.CollideVsPlane(p) != NoIntersect {
		t.Errorf("AABBox.InstersectPlane() indicated a box wasn't Outside that should have been.")
	}

	// Plane @ {-20, 0, 0}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, mgl.Vec3{-20, 0, 0})
	if b1.CollideVsPlane(p) != Intersect {
		t.Errorf("AABBox.InstersectPlane() indicated a box wasn't Inside that should have been.")
	}

	// Now do the same tests but with the box having an Offset
	b1.Offset = mgl.Vec3{25, 25, 25}

	// Plane @ {0, 0, 0}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, mgl.Vec3{0, 0, 0})
	if b1.CollideVsPlane(p) != Intersect {
		t.Errorf("AABBox.InstersectPlane() indicated a box wasn't Inside that should have been.")
	}

	// Plane @ {50, 0, 0}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, mgl.Vec3{50, 0, 0})
	if b1.CollideVsPlane(p) != NoIntersect {
		t.Errorf("AABBox.InstersectPlane() indicated a box wasn't Outside that should have been.")
	}

	// Plane @ {25, 25, 25}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, mgl.Vec3{25, 25, 25})
	if b1.CollideVsPlane(p) != Intersect {
		t.Errorf("AABBox.InstersectPlane() indicated a box didn't intersect that should have.")
	}

	// Reset the box to origin as a 16^3 cube
	b1.Min = mgl.Vec3{0.0, 0.0, 0.0}
	b1.Max = mgl.Vec3{16.0, 16.0, 16.0}
	b1.Offset = mgl.Vec3{0.0, 0.0, 0.0}

	// Add a second box some distance off
	var b2 AABBox
	b2.Min = mgl.Vec3{0.0, 0.0, 0.0}
	b2.Max = mgl.Vec3{16.0, 16.0, 16.0}
	b2.Offset = mgl.Vec3{32.0, 0.0, 0.0}

	// Plane @ {8, 8, 8}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, mgl.Vec3{8, 8, 8})
	if b1.CollideVsPlane(p) != Intersect {
		t.Errorf("AABBox.InstersectPlane() indicated a box didn't intersect that should have.")
	}
	if b2.CollideVsPlane(p) != Intersect {
		t.Errorf("AABBox.InstersectPlane() indicated a box wasn't Inside that should have been.")
	}

	// Plane @ {18, 8, 8}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, mgl.Vec3{18, 8, 8})
	if b1.CollideVsPlane(p) != NoIntersect {
		t.Errorf("AABBox.InstersectPlane() indicated a box wasn't Outside that should have been.")
	}
	if b2.CollideVsPlane(p) != Intersect {
		t.Errorf("AABBox.InstersectPlane() indicated a box wasn't Inside that should have been.")
	}

	// last test along border of box
	// Plane @ {0, 0, 0}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, mgl.Vec3{0, 0, 0})
	if b1.CollideVsPlane(p) != Intersect {
		t.Errorf("AABBox.InstersectPlane() indicated a box didn't intersect that should have.")
	}

}

func TestAABBoxCollisionVsSphere(t *testing.T) {
	var b1 AABBox
	var sphere Sphere

	b1.Min = mgl.Vec3{-10.0, -10.0, -10.0}
	b1.Max = mgl.Vec3{10.0, 10.0, 10.0}
	b1.Offset = mgl.Vec3{0.0, 0.0, 0.0}

	// Sphere {0, 0, 0} | r = 5.0
	sphere = Sphere{Center: mgl.Vec3{0.0, 0.0, 0.0}, Radius: 5.0}
	if b1.CollideVsSphere(&sphere) != Intersect {
		t.Errorf("AABBox.IntersectSphere() indicated a box didn't intersect that should have.")
	}

	// Sphere {15, 0, 0} | r = 5.0
	sphere = Sphere{Center: mgl.Vec3{15.0, 0.0, 0.0}, Radius: 5.0}
	if b1.CollideVsSphere(&sphere) != Intersect {
		t.Errorf("AABBox.IntersectSphere() indicated a box didn't intersect that should have.")
	}

	// Sphere {16, 0, 0} | r = 5.0
	sphere = Sphere{Center: mgl.Vec3{16.0, 0.0, 0.0}, Radius: 5.0}
	if b1.CollideVsSphere(&sphere) != NoIntersect {
		t.Errorf("AABBox.IntersectSphere() indicated a box intersected that should not have.")
	}

	// change the box offset ... effective {0, 0, 0}->{20, 20, 20}
	b1.Offset = mgl.Vec3{10.0, 10.0, 10.0}
	// Sphere {0, 0, 0} | r = 5.0
	sphere = Sphere{Center: mgl.Vec3{0.0, 0.0, 0.0}, Radius: 5.0}
	if b1.CollideVsSphere(&sphere) != Intersect {
		t.Errorf("AABBox.IntersectSphere() indicated a box didn't intersect that should have.")
	}

	sphere = Sphere{Center: mgl.Vec3{-6.0, 0.0, 0.0}, Radius: 5.0}
	if b1.CollideVsSphere(&sphere) != NoIntersect {
		t.Errorf("AABBox.IntersectSphere() indicated a box intersected that should not have.")
	}

}
