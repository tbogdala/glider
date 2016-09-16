// Copyright 2016, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package glider

import (
	"testing"

	mgl "github.com/go-gl/mathgl/mgl32"
)

func TestSphereCollisionVsSphere(t *testing.T) {
	var s1, s2, s3, s4 Sphere

	// Sphere {0, 0, 0} | r = 5.0
	s1 = Sphere{Center: mgl.Vec3{0.0, 0.0, 0.0}, Radius: 5.0}

	// Sphere {0, 0, 0} | r = 12.0
	s2 = Sphere{Center: mgl.Vec3{0.0, 0.0, 0.0}, Radius: 12.0}

	// Sphere {0, 10, 0} | r = 2.0
	s3 = Sphere{Center: mgl.Vec3{0.0, 10.0, 0.0}, Radius: 2.0}

	// Sphere {0, 12, 0} | r = 2.0
	s4 = Sphere{Center: mgl.Vec3{0.0, 12.0, 0.0}, Radius: 2.0}

	if s1.CollideVsSphere(&s3) != NoIntersect {
		t.Errorf("Sphere.IntersectSphere() indicated a sphere intersected that shouldn't have.")
	}

	if s1.CollideVsSphere(&s4) != NoIntersect {
		t.Errorf("Sphere.IntersectSphere() indicated a sphere intersected that shouldn't have.")
	}

	if s1.CollideVsSphere(&s2) != Intersect {
		t.Errorf("Sphere.IntersectSphere() indicated a sphere didn't intersect inside that should have.")
	}

	if s3.CollideVsSphere(&s4) != Intersect {
		t.Errorf("Sphere.IntersectSphere() indicated a sphere didn't intersect that should have.")
	}

	// now try some changes to the offset
	s1.Offset = mgl.Vec3{-20, 0, 0}
	if s1.CollideVsSphere(&s2) != NoIntersect {
		t.Errorf("Sphere.IntersectSphere() indicated a sphere didn intersected that shouldn't have.")
	}
}

func TestSphereCollisionVsRay(t *testing.T) {
	var s1, s2, s3 Sphere

	// Sphere {0, 0, 0} | r = 5.0
	s1 = Sphere{Center: mgl.Vec3{0.0, 0.0, 0.0}, Radius: 5.0}

	// Sphere {0, 0, 0} | r = 12.0
	s2 = Sphere{Center: mgl.Vec3{0.0, 0.0, 0.0}, Radius: 12.0}

	// Sphere {0, 10, 0} | r = 2.0
	s3 = Sphere{Center: mgl.Vec3{0.0, 10.0, 0.0}, Radius: 2.0}

	var r1 CollisionRay
	r1.Origin = mgl.Vec3{0.0, 0.0, 0.0}
	r1.SetDirection(mgl.Vec3{1.0, 0.0, 0.0})

	var r2 CollisionRay
	r2.Origin = mgl.Vec3{-100.0, 0.0, 0.0}
	r2.SetDirection(mgl.Vec3{1.0, 0.0, 0.0})

	intersect, _ := s1.CollideVsRay(&r1)
	if intersect != Intersect {
		t.Errorf("Sphere.IntersectRay() indicated a sphere didn't intersect that should have.")
	}

	intersect, _ = s3.CollideVsRay(&r1)
	if intersect != NoIntersect {
		t.Errorf("Sphere.IntersectRay() indicated a sphere intersected that shouldn't have.")
	}

	intersect, _ = s2.CollideVsRay(&r2)
	if intersect != Intersect {
		t.Errorf("Sphere.IntersectRay() indicated a sphere didn't intersect that should have.")
	}

	r2.Origin = mgl.Vec3{100.0, 0.0, 0.0}
	intersect, _ = s2.CollideVsRay(&r2)
	if intersect != NoIntersect {
		t.Errorf("Sphere.IntersectRay() indicated a sphere intersected that shouldn't have.")
	}

	// try change to the origin
	r2.Origin = mgl.Vec3{-100.0, 0.0, 0.0}
	s2.Offset = mgl.Vec3{0, 20.0, 0}
	intersect, _ = s2.CollideVsRay(&r2)
	if intersect != NoIntersect {
		t.Errorf("Sphere.IntersectRay() indicated a sphere intersected that shouldn't have.")
	}
}

func TestSphereCollisionVsAABBox(t *testing.T) {
	// these tests were copied from aabbox_test and adjusted to reverse the polarity of test conditions
	var b1 AABBox
	var sphere Sphere

	b1.Min = mgl.Vec3{-10.0, -10.0, -10.0}
	b1.Max = mgl.Vec3{10.0, 10.0, 10.0}
	b1.Offset = mgl.Vec3{0.0, 0.0, 0.0}

	// Sphere {0, 0, 0} | r = 5.0
	sphere = Sphere{Center: mgl.Vec3{0.0, 0.0, 0.0}, Radius: 5.0}
	if sphere.CollideVsAABBox(&b1) != Intersect {
		t.Errorf("Sphere.IntersectAABBox() indicated a box didn't intersect that should have.")
	}

	// try a change to the offset
	sphere.Center = mgl.Vec3{20, 20, 20}
	if sphere.CollideVsAABBox(&b1) != NoIntersect {
		t.Logf("sphere: %v ; box: %v\n", sphere, b1)
		t.Errorf("Sphere.IntersectAABBox() indicated a box intersected that shouldn't have.")
	}
	sphere.Offset = mgl.Vec3{0, 0, 0}

	// Sphere {15, 0, 0} | r = 5.0
	sphere = Sphere{Center: mgl.Vec3{15.0, 0.0, 0.0}, Radius: 5.0}
	if sphere.CollideVsAABBox(&b1) != Intersect {
		t.Errorf("Sphere.IntersectAABBox() indicated a box didn't intersect that should have.")
	}

	// change the box offset ... effective {0, 0, 0}->{20, 20, 20}
	b1.Offset = mgl.Vec3{10.0, 10.0, 10.0}
	// Sphere {0, 0, 0} | r = 5.0
	sphere = Sphere{Center: mgl.Vec3{0.0, 0.0, 0.0}, Radius: 5.0}
	if sphere.CollideVsAABBox(&b1) != Intersect {
		t.Errorf("Sphere.IntersectAABBox() indicated a box didn't intersect that should have.")
	}
}

func TestSphereCollisionVsPlane(t *testing.T) {
	var s1 Sphere
	s1.Radius = 10.0
	s1.Center = mgl.Vec3{0.0, 0.0, 0.0}

	// Plane @ {0, 0, 0}   Normal---> {1, 0, 0}
	var p *Plane
	planeNormal := mgl.Vec3{1.0, 0.0, 0.0}
	p = NewPlaneFromNormalAndPoint(planeNormal, mgl.Vec3{0, 0, 0})
	if s1.CollideVsPlane(p) != Intersect {
		t.Errorf("Sphere.InstersectPlane() indicated a sphere didn't intersect that should have.")
	}

	// try a change to the offset
	s1.Offset = mgl.Vec3{20, 0, 0}
	if s1.CollideVsPlane(p) != Intersect {
		t.Errorf("Sphere.InstersectPlane() indicated a sphere wasn't inside that should have been.")
	}
	s1.Offset = mgl.Vec3{0, 0, 0}

	// Plane @ {20, 0, 0}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, mgl.Vec3{20, 0, 0})
	if s1.CollideVsPlane(p) != NoIntersect {
		t.Errorf("Sphere.InstersectPlane() indicated a sphere wasn't Outside that should have been.")
	}

	// Plane @ {-20, 0, 0}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, mgl.Vec3{-20, 0, 0})
	if s1.CollideVsPlane(p) != Intersect {
		t.Errorf("Sphere.InstersectPlane() indicated a sphere wasn't Inside that should have been.")
	}

	// Now do the same tests but with the sphere having an Offset
	s1.Center = mgl.Vec3{25, 25, 25}

	// Plane @ {0, 0, 0}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, mgl.Vec3{0, 0, 0})
	if s1.CollideVsPlane(p) != Intersect {
		t.Errorf("Sphere.InstersectPlane() indicated a sphere wasn't Inside that should have been.")
	}

	// Plane @ {50, 0, 0}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, mgl.Vec3{50, 0, 0})
	if s1.CollideVsPlane(p) != NoIntersect {
		t.Errorf("Sphere.InstersectPlane() indicated a sphere wasn't Outside that should have been.")
	}

	// Plane @ {25, 25, 25}   Normal---> {1, 0, 0}
	p = NewPlaneFromNormalAndPoint(planeNormal, mgl.Vec3{25, 25, 25})
	if s1.CollideVsPlane(p) != Intersect {
		t.Errorf("Sphere.InstersectPlane() indicated a sphere didn't intersect that should have.")
	}
}
