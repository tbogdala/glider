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

	p1 = Vec2{0.5, 0.5}

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

	p1 = Vec3{0.5, 0.5, 0.5}

	if b1.IntersectPoint(&p1) == false {
		t.Error("AABBox.IntersectPoint() indicated false with a unit cube and a point that intersect.")
	}
}

func TestAABBoxNoCollisionVsPoint(t *testing.T) {
	var b1 AABBox
	var p1 Vec3

	b1.Min = Vec3{0.0, 0.0, 0.0}
	b1.Max = Vec3{1.0, 1.0, 1.0}

	p1 = Vec3{1.5, 1.5, 1.5}

	if b1.IntersectPoint(&p1) == true {
		t.Error("AABBox.IntersectPoint() indicated true with a unit cube and a point that do not intersect.")
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
