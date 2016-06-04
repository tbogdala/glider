// Copyright 2016, Timothy Bogdala <tdb@animal-machine.com>
// See the LICENSE file for more details.

package glider

import (
	"testing"
)


func TestSphereCollisionVsSphere(t *testing.T) {
	var s1, s2, s3, s4 Sphere

	// Sphere {0, 0, 0} | r = 5.0
	s1 = Sphere{Center: Vec3{0.0, 0.0, 0.0}, Radius: 5.0}

    // Sphere {0, 0, 0} | r = 12.0
	s2 = Sphere{Center: Vec3{0.0, 0.0, 0.0}, Radius: 12.0}

    // Sphere {0, 10, 0} | r = 2.0
	s3 = Sphere{Center: Vec3{0.0, 10.0, 0.0}, Radius: 2.0}

    // Sphere {0, 12, 0} | r = 2.0
	s4 = Sphere{Center: Vec3{0.0, 12.0, 0.0}, Radius: 2.0}

    if s1.IntersectSphere(&s3) != Outside {
		t.Errorf("Sphere.IntersectSphere() indicated a sphere intersected that shouldn't have.")
	}

    if s1.IntersectSphere(&s4) != Outside {
        t.Errorf("Sphere.IntersectSphere() indicated a sphere intersected that shouldn't have.")
    }

	if s1.IntersectSphere(&s2) != Inside {
		t.Errorf("Sphere.IntersectSphere() indicated a sphere didn't intersect inside that should have.")
	}

    if s3.IntersectSphere(&s4) != Intersect {
		t.Errorf("Sphere.IntersectSphere() indicated a sphere didn't intersect that should have.")
	}
}

func TestSphereCollisionVsRay(t *testing.T) {
	var s1, s2, s3 Sphere

	// Sphere {0, 0, 0} | r = 5.0
	s1 = Sphere{Center: Vec3{0.0, 0.0, 0.0}, Radius: 5.0}

    // Sphere {0, 0, 0} | r = 12.0
	s2 = Sphere{Center: Vec3{0.0, 0.0, 0.0}, Radius: 12.0}

    // Sphere {0, 10, 0} | r = 2.0
	s3 = Sphere{Center: Vec3{0.0, 10.0, 0.0}, Radius: 2.0}

    var r1 CollisionRay
	r1.Origin = Vec3{0.0, 0.0, 0.0}
	r1.SetDirection(Vec3{1.0, 0.0, 0.0})

    var r2 CollisionRay
	r2.Origin = Vec3{-100.0, 0.0, 0.0}
	r2.SetDirection(Vec3{1.0, 0.0, 0.0})

    if s1.IntersectRay(&r1) != true {
		t.Errorf("Sphere.IntersectRay() indicated a sphere didn't intersect that should have.")
	}

    if s3.IntersectRay(&r1) != false {
		t.Errorf("Sphere.IntersectRay() indicated a sphere intersected that shouldn't have.")
	}

    if s2.IntersectRay(&r2) != true {
		t.Errorf("Sphere.IntersectRay() indicated a sphere didn't intersect that should have.")
	}

    r2.Origin = Vec3{100.0, 0.0, 0.0}
    if s2.IntersectRay(&r2) != false {
        t.Errorf("Sphere.IntersectRay() indicated a sphere intersected that shouldn't have.")
	}

}

func TestSphereCollisionVsAABBox(t *testing.T) {
    // these tests were copied from aabbox_test and adjusted to reverse the polarity of test conditions
	var b1 AABBox
	var sphere Sphere

	b1.Min = Vec3{-10.0, -10.0, -10.0}
	b1.Max = Vec3{10.0, 10.0, 10.0}
	b1.Offset = Vec3{0.0, 0.0, 0.0}

	// Sphere {0, 0, 0} | r = 5.0
	sphere = Sphere{Center: Vec3{0.0, 0.0, 0.0}, Radius: 5.0}
	if sphere.IntersectAABBox(&b1) != Intersect {
		t.Errorf("Sphere.IntersectAABBox() indicated a box didn't intersect that should have.")
	}

	// Sphere {15, 0, 0} | r = 5.0
	sphere = Sphere{Center: Vec3{15.0, 0.0, 0.0}, Radius: 5.0}
	if sphere.IntersectAABBox(&b1) != Intersect {
		t.Errorf("Sphere.IntersectAABBox() indicated a box didn't intersect that should have.")
	}

	// Sphere {16, 0, 0} | r = 5.0
	sphere = Sphere{Center: Vec3{16.0, 0.0, 0.0}, Radius: 5.0}
	if sphere.IntersectAABBox(&b1) != Inside {
		t.Errorf("Sphere.IntersectAABBox() indicated a box didn't intersec that should not have.")
	}

	// change the box offset ... effective {0, 0, 0}->{20, 20, 20}
	b1.Offset = Vec3{10.0, 10.0, 10.0}
	// Sphere {0, 0, 0} | r = 5.0
	sphere = Sphere{Center: Vec3{0.0, 0.0, 0.0}, Radius: 5.0}
	if sphere.IntersectAABBox(&b1) != Intersect {
		t.Errorf("Sphere.IntersectAABBox() indicated a box didn't intersect that should have.")
	}

	sphere = Sphere{Center: Vec3{-6.0, 0.0, 0.0}, Radius: 5.0}
	if sphere.IntersectAABBox(&b1) != Inside {
		t.Errorf("Sphere.IntersectAABBox() indicated a box didn't intersec that should not have.")
	}

}
