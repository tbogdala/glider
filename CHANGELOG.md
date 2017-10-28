Version v0.2.2
=============

* NEW: Added Collider.GetOffset() to the Collider Interface.


Version v0.2.0
==============

* APIBREAK: Changed results from being Intersect/Outside/Inside to just being Intersect/NoIntersect.
  The theory behind this is that this is supposed to be a 'coarse' collision library and
  'quick' so a hit or no-hit test should be sufficient. The exception will be Ray's which
  should also return a distance.

* APIBREAK: No longer supplying vector types with this library and will instead be using
  Mathgl for the necessary vectors.

* NEW: Added OBBox collisions. They're oriented bounding boxes and have rotations. The support
  is incomplete at present.


