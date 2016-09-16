Glider v0.2.0-development
=========================

Glider is a simple collision library written in the [Go][golang] programming language.


Installation
------------

The library can be installed with the following command:

```bash
go get github.com/go-gl/mathgl/mgl32
go get github.com/tbogdala/glider
```


Current Features
----------------

* Basic 2d box collision
* AABB intersection vs AABB
* AABB intersection vs Sphere
* AABB intersection vs Ray
* AABB intersection vs Plane
* Sphere intersection tests vs AABB
* Sphere intersection tests vs Sphere
* Sphere intersection tests vs Ray
* Sphere intersection tests vs Plane

Documentation
-------------

Currently, you'll have to use godoc to read the API documentation and check
out the unit tests to figure out how to use the library.

It should be mostly easy to figure out, though.


LICENSE
=======

Glider is released under the BSD license. See the [LICENSE][license-link] file for more details.


[golang]: https://golang.org/
[license-link]: https://raw.githubusercontent.com/tbogdala/glider/master/LICENSE
