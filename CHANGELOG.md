Changes since v0.1.0
====================

* APIBREAK: Changed results from being Intersect/Outside/Inside to just being Intersect/NoIntersect.
  The theory behind this is that this is supposed to be a 'coarse' collision library and
  'quick' so a hit or no-hit test should be sufficient. The exception will be Ray's which
  should also return a distance.

* NEW: Added OBBox collisions. They're oriented bounding boxes and have rotations. The support
  is incomplete at present.
