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
13. Cylinders
14. Groups
15. Triangles
16. Constructive Solid Geometry (CSG)
17. Next Steps

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