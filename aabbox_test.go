// Copyright 2015, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package glider

import (
	"testing"
)

func TestAABSquareCollisionVsPoint(t *testing.T) {
	var s1 AABSquare
	var p1 Vec2

	s1.Min = Vec2{0.0, 0.0}
	s1.Max = Vec2{1.0, 1.0}
	s1.Offset = Vec2{0.0, 0.0}

	p1 = Vec2{0.5, 0.5}
	if s1.IntersectPoint(&p1) == false {
		t.Error("AABSquare.IntersectPoint() indicated false with a unit square and a point that intersect.")
	}

	s1.Offset = Vec2{10.0, 5.0}
	p1 = Vec2{0.5, 0.5}
	if s1.IntersectPoint(&p1) == true {
		t.Error("AABSquare.IntersectPoint() indicated false with a unit square and a point that intersect.")
	}

	p1 = Vec2{10.5, 5.5}
	if s1.IntersectPoint(&p1) == false {
		t.Error("AABSquare.IntersectPoint() indicated false with a unit square and a point that intersect.")
	}

}

func TestAABBoxNoCollision(t *testing.T) {
	var b1, b2 AABBox

	b1.Min = Vec3{0.0, 0.0, 0.0}
	b1.Max = Vec3{1.0, 1.0, 1.0}

	b2.Min = Vec3{2.0, 2.0, 2.0}
	b2.Max = Vec3{3.0, 3.0, 3.0}

	if b1.IntersectBox(&b2) == true {
		t.Error("AABBox.IntersectBox() indicated true with two unit cubes that don't intersect.")
	}
}

func TestAABBoxCollision(t *testing.T) {
	var b1, b2 AABBox

	b1.Min = Vec3{0.0, 0.0, 0.0}
	b1.Max = Vec3{1.0, 1.0, 1.0}

	b2.Min = Vec3{0.5, 0.5, 0.5}
	b2.Max = Vec3{3.0, 3.0, 3.0}

	if b1.IntersectBox(&b2) == false {
		t.Error("AABBox.IntersectBox() indicated false with two unit cubes that intersect.")
	}
}

func TestAABBoxSiblingCollision(t *testing.T) {
	var b1, b2 AABBox

	b1.Min = Vec3{0.0, 0.0, 0.0}
	b1.Max = Vec3{1.0, 1.0, 1.0}

	b2.Min = Vec3{1.0, 1.0, 1.0}
	b2.Max = Vec3{3.0, 3.0, 3.0}

	if b1.IntersectBox(&b2) == false {
		t.Error("AABBox.IntersectBox() indicated false with two unit cubes that share an edge.")
	}
}

func TestAABBoxCollisionVsPoint(t *testing.T) {
	var b1 AABBox
	var p1 Vec3

	b1.Min = Vec3{0.0, 0.0, 0.0}
	b1.Max = Vec3{1.0, 1.0, 1.0}
	b1.Offset = Vec3{0.0, 0.0, 0.0}

	p1 = Vec3{0.5, 0.5, 0.5}
	if b1.IntersectPoint(&p1) == false {
		t.Error("AABBox.IntersectPoint() indicated false with a unit cube and a point that intersect.")
	}

	b1.Offset = Vec3{3.0, 3.0, 3.0}
	if b1.IntersectPoint(&p1) == true {
		t.Error("AABBox.IntersectPoint() indicated false with a unit cube and a point that intersect.")
	}

	p1 = Vec3{3.5, 3.5, 3.5}
	if b1.IntersectPoint(&p1) == false {
		t.Error("AABBox.IntersectPoint() indicated false with a unit cube and a point that intersect.")
	}
}

func TestAABBoxCollisionVsEdgePoint(t *testing.T) {
	var b1 AABBox
	var p1 Vec3

	b1.Min = Vec3{0.0, 0.0, 0.0}
	b1.Max = Vec3{1.0, 1.0, 1.0}

	p1 = Vec3{1.0, 1.0, 1.0}

	if b1.IntersectPoint(&p1) == false {
		t.Error("AABBox.IntersectPoint() indicated false with a unit cube and a point along an edge.")
	}
}

func TestAABBoxCollisionVsRay(t *testing.T) {
	var b1 AABBox
	var r1 CollisionRay

	b1.Min = Vec3{-1.0, -1.0, -1.0}
	b1.Max = Vec3{1.0, 1.0, 1.0}

	// cast at the center
	r1.Origin = Vec3{5.0, 0.0, 0.0}
	r1.SetDirection(Vec3{-1.0, 0.0, 0.0})

	intersect, _ := b1.IntersectRay(&r1)
	if intersect == false {
		t.Error("AABBox.IntersectRay() indicated false with a ray pointed at it's center.")
	}

	// cast away from it
	r1.Origin = Vec3{5.0, 0.0, 0.0}
	r1.SetDirection(Vec3{1.0, 0.0, 0.0})

	intersect, _ = b1.IntersectRay(&r1)
	if intersect == true {
		t.Error("AABBox.IntersectRay() indicated true with a ray pointed away from it.")
	}

	// cast at the edge
	r1.Origin = Vec3{10.0, 1.0, 1.0}
	r1.SetDirection(Vec3{-1.0, 0.0, 0.0})

	intersect, _ = b1.IntersectRay(&r1)
	if intersect == false {
		t.Error("AABBox.IntersectRay() indicated false with a ray pointed at it's edge.")
	}

	// cast from inside
	r1.Origin = Vec3{0.0, 0.0, 0.0}
	r1.SetDirection(Vec3{1.0, 1.0, 1.0})

	intersect, _ = b1.IntersectRay(&r1)
	if intersect == false {
		t.Error("AABBox.IntersectRay() indicated false with a ray starting at the center of the box.")
	}
}

func TestAABBoxCollisionVsRay2(t *testing.T) {
	var b1 AABBox
	var r1 CollisionRay

	b1.Min = Vec3{-1.0, -1.0, -1.0}
	b1.Max = Vec3{1.0, 1.0, 1.0}
	b1.Offset = Vec3{10.0, 0.0, 0.0}

	// cast at the center
	r1.Origin = Vec3{15.0, 0.0, 0.0}
	r1.SetDirection(Vec3{-1.0, 0.0, 0.0})

	intersect, _ := b1.IntersectRay(&r1)
	if intersect == false {
		t.Error("AABBox.IntersectRay2() indicated false with a ray pointed at it's center.")
	}

	// cast away from it
	r1.Origin = Vec3{15.0, 0.0, 0.0}
	r1.SetDirection(Vec3{1.0, 0.0, 0.0})

	intersect, _ = b1.IntersectRay(&r1)
	if intersect == true {
		t.Error("AABBox.IntersectRay2() indicated true with a ray pointed away from it.")
	}

	// cast at the edge
	r1.Origin = Vec3{10.0, 1.0, 1.0}
	r1.SetDirection(Vec3{-1.0, 0.0, 0.0})

	intersect, _ = b1.IntersectRay(&r1)
	if intersect == false {
		t.Error("AABBox.IntersectRay2() indicated false with a ray pointed at it's edge.")
	}

	// cast from inside
	r1.Origin = Vec3{10.0, 0.0, 0.0}
	r1.SetDirection(Vec3{1.0, 1.0, 1.0})

	intersect, _ = b1.IntersectRay(&r1)
	if intersect == false {
		t.Error("AABBox.IntersectRay2() indicated false with a ray starting at the center of the box.")
	}
}

func TestAABBoxCollisionVsPlane(t *testing.T) {
	var b1 AABBox
	var p *Plane

	b1.Min = Vec3{-10.0, -10.0, -10.0}
	b1.Max = Vec3{10.0, 10.0, 10.0}
	b1.Offset = Vec3{0.0, 0.0, 0.0}

	// Plane @ {0, 0, 0}   Normal---> {1, 0, 0}
	planeNormal := Vec3{1.0, 0.0, 0.0}
	p = NewPlaneFromNormalAndPoint(planeNormal, Vec3{0, 0, 0})
	if b1.IntersectPlane(p) != Intersect {
		t.Errorf("AABBox.InstersectPlane() indicated a box didn't intersect that should have.")
	}

	// Plane @ {20, 0, 0}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, Vec3{20, 0, 0})
	if b1.IntersectPlane(p) != Outside {
		t.Errorf("AABBox.InstersectPlane() indicated a box wasn't Outside that should have been.")
	}

	// Plane @ {-20, 0, 0}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, Vec3{-20, 0, 0})
	if b1.IntersectPlane(p) != Inside {
		t.Errorf("AABBox.InstersectPlane() indicated a box wasn't Inside that should have been.")
	}

	// Now do the same tests but with the box having an Offset
	b1.Offset = Vec3{25, 25, 25}

	// Plane @ {0, 0, 0}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, Vec3{0, 0, 0})
	if b1.IntersectPlane(p) != Inside {
		t.Errorf("AABBox.InstersectPlane() indicated a box wasn't Inside that should have been.")
	}

	// Plane @ {50, 0, 0}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, Vec3{50, 0, 0})
	if b1.IntersectPlane(p) != Outside {
		t.Errorf("AABBox.InstersectPlane() indicated a box wasn't Outside that should have been.")
	}

	// Plane @ {25, 25, 25}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, Vec3{25, 25, 25})
	if b1.IntersectPlane(p) != Intersect {
		t.Errorf("AABBox.InstersectPlane() indicated a box didn't intersect that should have.")
	}

	// Reset the box to origin as a 16^3 cube
	b1.Min = Vec3{0.0, 0.0, 0.0}
	b1.Max = Vec3{16.0, 16.0, 16.0}
	b1.Offset = Vec3{0.0, 0.0, 0.0}

	// Add a second box some distance off
	var b2 AABBox
	b2.Min = Vec3{0.0, 0.0, 0.0}
	b2.Max = Vec3{16.0, 16.0, 16.0}
	b2.Offset = Vec3{32.0, 0.0, 0.0}

	// Plane @ {8, 8, 8}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, Vec3{8, 8, 8})
	if b1.IntersectPlane(p) != Intersect {
		t.Errorf("AABBox.InstersectPlane() indicated a box didn't intersect that should have.")
	}
	if b2.IntersectPlane(p) != Inside {
		t.Errorf("AABBox.InstersectPlane() indicated a box wasn't Inside that should have been.")
	}

	// Plane @ {18, 8, 8}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, Vec3{18, 8, 8})
	if b1.IntersectPlane(p) != Outside {
		t.Errorf("AABBox.InstersectPlane() indicated a box wasn't Outside that should have been.")
	}
	if b2.IntersectPlane(p) != Inside {
		t.Errorf("AABBox.InstersectPlane() indicated a box wasn't Inside that should have been.")
	}

	// last test along border of box
	// Plane @ {0, 0, 0}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, Vec3{0, 0, 0})
	if b1.IntersectPlane(p) != Intersect {
		t.Errorf("AABBox.InstersectPlane() indicated a box didn't intersect that should have.")
	}

}

func TestAABBoxCollisionVsSphere(t *testing.T) {
	var b1 AABBox
	var sphere Sphere

	b1.Min = Vec3{-10.0, -10.0, -10.0}
	b1.Max = Vec3{10.0, 10.0, 10.0}
	b1.Offset = Vec3{0.0, 0.0, 0.0}

	// Sphere {0, 0, 0} | r = 5.0
	sphere = Sphere{Center: Vec3{0.0, 0.0, 0.0}, Radius: 5.0}
	if b1.IntersectSphere(&sphere) != Intersect {
		t.Errorf("AABBox.IntersectSphere() indicated a box didn't intersect that should have.")
	}

	// Sphere {15, 0, 0} | r = 5.0
	sphere = Sphere{Center: Vec3{15.0, 0.0, 0.0}, Radius: 5.0}
	if b1.IntersectSphere(&sphere) != Intersect {
		t.Errorf("AABBox.IntersectSphere() indicated a box didn't intersect that should have.")
	}

	// Sphere {16, 0, 0} | r = 5.0
	sphere = Sphere{Center: Vec3{16.0, 0.0, 0.0}, Radius: 5.0}
	if b1.IntersectSphere(&sphere) != Outside {
		t.Errorf("AABBox.IntersectSphere() indicated a box intersected that should not have.")
	}

	// change the box offset ... effective {0, 0, 0}->{20, 20, 20}
	b1.Offset = Vec3{10.0, 10.0, 10.0}
	// Sphere {0, 0, 0} | r = 5.0
	sphere = Sphere{Center: Vec3{0.0, 0.0, 0.0}, Radius: 5.0}
	if b1.IntersectSphere(&sphere) != Intersect {
		t.Errorf("AABBox.IntersectSphere() indicated a box didn't intersect that should have.")
	}

	sphere = Sphere{Center: Vec3{-6.0, 0.0, 0.0}, Radius: 5.0}
	if b1.IntersectSphere(&sphere) != Outside {
		t.Errorf("AABBox.IntersectSphere() indicated a box intersected that should not have.")
	}

}
