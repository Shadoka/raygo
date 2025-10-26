## raygo

Go lang implementation of [The Ray Tracer Challenge](http://raytracerchallenge.com/)

## Implementation progress

1. Tuples, Points and Vectors ✔
2. Drawing on a Canvas ✔
3. Matrices ✔
4. Matrix Transformations ✔
5. Ray-Sphere Intersections ✔
6. Light and Shading ✔
7. Making a Scene ✔
8. Shadows ✔
9. Planes ✔
10. Patterns ✔
11. Reflection and Refraction ✔
12. Cubes ✔
13. Cylinders ✔
14. Groups ✔
15. Triangles ✔
16. Constructive Solid Geometry (CSG)
17. Next Steps

I don't think that I'll implement CSG, because I don't see the use case for myself at this point.
Feel free to submit a PR though.

## Custom improvements

1. Performance optimisations (multithreaded rendering + caching)
2. Camera movement & multi frame rendering

## Examples

### Chapter 5
![Exercise 5](examples/chapter5.png)

Drawing of a sphere via ray-sphere intersections.

### Chapter 6
![Exercise 6](examples/chapter6.png)

Drawing of a sphere with lightning via phong shader.

### Chapter 7
![Exercise 7](examples/chapter7.png)

Scene drawn from the viewpoint of a camera

### Chapter 8
![Exercise 8](examples/chapter8.png)

Shapes now cast shadows

### Chapter 9
![Exercise 9](examples/chapter9.png)

Plane got added as additional shape

![Exercise 9 with backdrop](examples/chapter9_backdrop.png)

An additional plane got added as backdrop.

I think the reason why the ground looks so much darker is because I lowered the light. The flat reflection angle is probably why
the floor looks so dark compared with the other image.

The following image is fundamentally the same besides the light being higher.

![Exercise 9 with backdrop and higher light](examples/chapter9_backdrop_higherlight.png)

### Chapter 10
![Exercise 10](examples/chapter10_stripes.png)

Spheres with attached stripe patterns

![Exercise 10](examples/chapter10_gradients.png)

Spheres with attached gradient patterns

![Exercise 10](examples/chapter10_patterns.png)

Multiple patterns at once

### Chapter 11
![Exercise 11](examples/chapter11_reflections.png)

Material can now be reflective

![Exercise 11](examples/chapter11_refractions.png)

A glass sphere on the outside that contains a sphere of air in the inside

### Chapter 12
![Exercise 12](examples/chapter12_cubes.png)

Cubes got added as shapes

### Chapter 13
![Exercise 13](examples/chapter13_cylinders.png)

Cylinders and cones got added

### Chapter 14
![Exercise 14](examples/chapter14_groups.png)

Groups are now added, allowing for creating reusable complex shapes and performance improvements.

The implementation differs slightly from the book. The implementation of the bounding box intersection is
according to [this blog article](https://tavianator.com/2011/ray_box.html)

### Chapter 15
![Exercise 15](examples/chapter15_teapot.png)

### Multi frame rendering & camera movement

It is now possible to define a simple, circular camera movement around a given point. Following parameters are
configurable:
* Camera rotation in Radians
* Duration of the movement in seconds
* How many frames per second are to be rendered
For still images it is easiest to set the duration to 1 and define via FPS parameter how many images you want
to be rendered.

Here is an example with the following parameters:
```golang
animation := scene.CreateCameraAnimation(math.Radians(90), 1, 3)
cam.Animation = animation
```

Image #1
![Multiframe 1](examples/multiframe1.png)
Image #2
![Multiframe 1](examples/multiframe2.png)
Image #3
![Multiframe 1](examples/multiframe3.png)

## Outlook

These are the next steps for me:

* Scene descriptions in YAML format
* Camera smoothing for stop & start of movement
* More sophisticated camera movement options (define anchor points and have the camera move to them in order)