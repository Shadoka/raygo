package parser

import (
	"fmt"
	"log"
	gomath "math"
	"raygo/geometry"
	"raygo/lighting"
	"raygo/math"
	"raygo/obj"
	"raygo/scene"

	"github.com/goccy/go-yaml"
)

const SHEARING_TF = "shearing"
const TRANSLATION_TF = "translation"
const SCALING_TF = "scaling"
const ROTATION_TF = "rotation"

var yamlColors map[string]*NamedColorModel
var yamlTransforms map[string]*NamedTransformModel
var yamlMaterials map[string]*NamedMaterialModel
var yamlCubes map[string]*CubeModel
var yamlPlanes map[string]*PlaneModel
var yamlSpheres map[string]*SphereModel
var yamlCylinders map[string]*CylinderModel
var yamlCones map[string]*ConeModel
var yamlTriangles map[string]*TriangleModel
var yamlObjects map[string]*ObjectModel
var yamlGroups map[string]*GroupModel
var yamlStripePatterns map[string]*StripePatternModel
var yamlGradientPatterns map[string]*GradientPatternModel
var yamlRingPatterns map[string]*RingPatternModel
var yamlCheckerPatterns map[string]*CheckerPatternModel

var raygoColors map[string]*math.Color
var raygoMaterials map[string]*geometry.Material
var raygoTransforms map[string]math.Matrix
var raygoPatterns map[string]geometry.Pattern
var raygoShapes map[string]geometry.Shape

// we need to keep track of shapes that are children of groups
// because we only want to add the parent to the scene
var childrenObjects map[string]struct{}

func ParseYaml(content string) *YamlDescription {
	description := YamlDescription{}
	if err := yaml.Unmarshal([]byte(content), &description); err != nil {
		log.Fatalf("could not unmarshal yaml: %v", err)
	}
	return &description
}

func initReferences(yml *YamlDescription) {
	yamlColors = make(map[string]*NamedColorModel, 0)
	for _, c := range yml.Colors {
		yamlColors[c.Name] = &c
	}

	yamlMaterials = make(map[string]*NamedMaterialModel, 0)
	for _, m := range yml.Materials {
		yamlMaterials[m.Name] = &m
	}

	yamlTransforms = make(map[string]*NamedTransformModel, 0)
	for _, t := range yml.Transforms {
		yamlTransforms[t.Name] = &t
	}

	yamlPlanes = make(map[string]*PlaneModel, 0)
	for _, p := range yml.Scene.Planes {
		yamlPlanes[p.Name] = &p
	}

	yamlSpheres = make(map[string]*SphereModel, 0)
	for _, s := range yml.Scene.Spheres {
		yamlSpheres[s.Name] = &s
	}

	yamlCubes = make(map[string]*CubeModel, 0)
	for _, c := range yml.Scene.Cubes {
		yamlCubes[c.Name] = &c
	}

	yamlCylinders = make(map[string]*CylinderModel, 0)
	for _, c := range yml.Scene.Cylinders {
		yamlCylinders[c.Name] = &c
	}

	yamlCones = make(map[string]*ConeModel, 0)
	for _, c := range yml.Scene.Cones {
		yamlCones[c.Name] = &c
	}

	yamlTriangles = make(map[string]*TriangleModel, 0)
	for _, t := range yml.Scene.Triangles {
		yamlTriangles[t.Name] = &t
	}

	yamlGroups = make(map[string]*GroupModel, 0)
	for _, g := range yml.Scene.Groups {
		yamlGroups[g.Name] = &g
	}

	yamlObjects = make(map[string]*ObjectModel, 0)
	for _, o := range yml.Scene.Objects {
		yamlObjects[o.Name] = &o
	}

	yamlStripePatterns = make(map[string]*StripePatternModel, 0)
	for _, p := range yml.Patterns.Stripe {
		yamlStripePatterns[p.Name] = &p
	}

	yamlGradientPatterns = make(map[string]*GradientPatternModel, 0)
	for _, p := range yml.Patterns.Gradient {
		yamlGradientPatterns[p.Name] = &p
	}

	yamlRingPatterns = make(map[string]*RingPatternModel, 0)
	for _, p := range yml.Patterns.Ring {
		yamlRingPatterns[p.Name] = &p
	}

	yamlCheckerPatterns = make(map[string]*CheckerPatternModel, 0)
	for _, p := range yml.Patterns.Checker {
		yamlCheckerPatterns[p.Name] = &p
	}
}

func ValidateReferences(yml *YamlDescription) []error {
	validationResult := make([]error, 0)
	initReferences(yml)

	validationResult = append(validationResult, validatePatternReferences(yml)...)
	validationResult = append(validationResult, validateMaterialReferences(yml)...)
	validationResult = append(validationResult, validateSceneObjectReferences(yml)...)
	validationResult = append(validationResult, validateCameraReferences(yml)...)

	return validationResult
}

func validatePatternReferences(yml *YamlDescription) []error {
	validationResult := make([]error, 0)

	for _, p := range yml.Patterns.Checker {
		if yamlColors[p.ColorA] == nil {
			validationResult = append(validationResult, fmt.Errorf("cannot resolve color '%v' for pattern '%v'", p.ColorA, p.Name))
		}
		if yamlColors[p.ColorB] == nil {
			validationResult = append(validationResult, fmt.Errorf("cannot resolve color '%v' for pattern '%v'", p.ColorB, p.Name))
		}
	}

	for _, p := range yml.Patterns.Ring {
		if yamlColors[p.ColorA] == nil {
			validationResult = append(validationResult, fmt.Errorf("cannot resolve color '%v' for pattern '%v'", p.ColorA, p.Name))
		}
		if yamlColors[p.ColorB] == nil {
			validationResult = append(validationResult, fmt.Errorf("cannot resolve color '%v' for pattern '%v'", p.ColorB, p.Name))
		}
	}

	for _, p := range yml.Patterns.Gradient {
		if yamlColors[p.ColorA] == nil {
			validationResult = append(validationResult, fmt.Errorf("cannot resolve color '%v' for pattern '%v'", p.ColorA, p.Name))
		}
		if yamlColors[p.ColorB] == nil {
			validationResult = append(validationResult, fmt.Errorf("cannot resolve color '%v' for pattern '%v'", p.ColorB, p.Name))
		}
	}

	for _, p := range yml.Patterns.Stripe {
		if yamlColors[p.ColorA] == nil {
			validationResult = append(validationResult, fmt.Errorf("cannot resolve color '%v' for pattern '%v'", p.ColorA, p.Name))
		}
		if yamlColors[p.ColorB] == nil {
			validationResult = append(validationResult, fmt.Errorf("cannot resolve color '%v' for pattern '%v'", p.ColorB, p.Name))
		}
	}

	return validationResult
}

func validateMaterialReferences(yml *YamlDescription) []error {
	validationResult := make([]error, 0)

	for _, m := range yml.Materials {
		if m.Color != "" && yamlColors[m.Color] == nil {
			validationResult = append(validationResult, fmt.Errorf("cannot resolve color '%v' for material '%v'", m.Color, m.Name))
		}

		if m.Pattern != "" && !containsPattern(m.Pattern) {
			validationResult = append(validationResult, fmt.Errorf("cannot resolve pattern '%v' for material '%v'", m.Pattern, m.Name))
		}
	}

	return validationResult
}

func validateSceneObjectReferences(yml *YamlDescription) []error {
	validationResult := make([]error, 0)

	for _, p := range yml.Scene.Planes {
		validationResult = append(validationResult, validateCommonSceneObjectReferences(p.CommonSceneObject)...)
	}

	for _, c := range yml.Scene.Cubes {
		validationResult = append(validationResult, validateCommonSceneObjectReferences(c.CommonSceneObject)...)
	}

	for _, s := range yml.Scene.Spheres {
		validationResult = append(validationResult, validateCommonSceneObjectReferences(s.CommonSceneObject)...)
	}

	for _, t := range yml.Scene.Triangles {
		validationResult = append(validationResult, validateCommonSceneObjectReferences(t.CommonSceneObject)...)
	}

	for _, c := range yml.Scene.Cylinders {
		validationResult = append(validationResult, validateCommonSceneObjectReferences(c.CommonSceneObject)...)
	}

	for _, c := range yml.Scene.Cones {
		validationResult = append(validationResult, validateCommonSceneObjectReferences(c.CommonSceneObject)...)
	}

	for _, o := range yml.Scene.Objects {
		validationResult = append(validationResult, validateCommonSceneObjectReferences(o.CommonSceneObject)...)
	}

	for _, g := range yml.Scene.Groups {
		validationResult = append(validationResult, validateCommonSceneObjectReferences(g.CommonSceneObject)...)

		for _, child := range g.Children {
			if !containsSceneObject(child) {
				err := fmt.Errorf("cannot resolve child '%v' for group '%v'", child, g.Name)
				validationResult = append(validationResult, err)
			}
		}
	}

	return validationResult
}

func validateCameraReferences(yml *YamlDescription) []error {
	valResult := make([]error, 0)

	if yml.Camera.To == nil && !containsSceneObject(yml.Camera.LookAt) {
		err := fmt.Errorf("cannot resolve scene object '%v' for camera", yml.Camera.LookAt)
		valResult = append(valResult, err)
	}

	return valResult
}

func validateCommonSceneObjectReferences(sceneObject CommonSceneObject) []error {
	valResult := make([]error, 0)
	if sceneObject.Material != "" && yamlMaterials[sceneObject.Material] == nil {
		err := fmt.Errorf("cannot resolve material '%v' for scene object '%v'",
			sceneObject.Material, sceneObject.Name)
		valResult = append(valResult, err)
	}
	if sceneObject.Transform != "" && yamlTransforms[sceneObject.Transform] == nil {
		err := fmt.Errorf("cannot resolve transform '%v' for scene object '%v'",
			sceneObject.Transform, sceneObject.Name)
		valResult = append(valResult, err)
	}
	return valResult
}

func containsSceneObject(name string) bool {
	return yamlSpheres[name] != nil ||
		yamlPlanes[name] != nil ||
		yamlCubes[name] != nil ||
		yamlCylinders[name] != nil ||
		yamlCones[name] != nil ||
		yamlTriangles[name] != nil ||
		yamlGroups[name] != nil ||
		yamlObjects[name] != nil
}

func containsPattern(name string) bool {
	return yamlStripePatterns[name] != nil ||
		yamlGradientPatterns[name] != nil ||
		yamlRingPatterns[name] != nil ||
		yamlCheckerPatterns[name] != nil
}

func CreateWorld(yml *YamlDescription, directory string) *scene.World {
	world := scene.EmptyWorld()

	createRaygoColors()
	createRaygoTransformations()
	createRaygoPatterns()
	createRaygoMaterials()
	createRaygoShapes(directory)

	sceneObjects := collectRootElements()
	world.Objects = sceneObjects

	light := createLight(yml.Light)
	world.Light = &light

	return world
}

func CreateCamera(yml *YamlDescription) scene.Camera {
	camera := scene.CreateCamera(yml.Width, yml.Height, gomath.Pi/3.0)
	from := mapPoint(yml.Camera.From)
	var to math.Point
	if yml.Camera.LookAt != "" {
		to = geometry.GetCenter(raygoShapes[yml.Camera.LookAt])
	} else {
		to = mapPoint(yml.Camera.To)
	}

	up := math.CreateVector(0.0, 1.0, 0.0)
	camera.Position = scene.CreateCameraPosition(from, to, up)
	if yml.Camera.Animation != nil {
		camera.Animation = createCameraAnimation(yml.Camera.Animation)
	}

	return *camera
}

func createCameraAnimation(yamlAnimation *CircularCameraAnimation) *scene.CameraAnimation {
	return scene.CreateCameraAnimation(math.Radians(yamlAnimation.Degrees), yamlAnimation.Time, yamlAnimation.Fps)
}

func createLight(yamlLight LightModel) lighting.Light {
	p := mapPoint(yamlLight.Position)
	intensity := mapColor(yamlLight.Intensity)
	return lighting.CreateLight(p, intensity)
}

func collectRootElements() []geometry.Shape {
	elements := make([]geometry.Shape, 0)

	for name, shape := range raygoShapes {
		if _, ok := childrenObjects[name]; !ok {
			elements = append(elements, shape)
		}
	}

	return elements
}

func createRaygoColors() {
	raygoColors = make(map[string]*math.Color)
	for _, v := range yamlColors {
		rc := math.CreateColor(math.BToF(v.R), math.BToF(v.G), math.BToF(v.B))
		raygoColors[v.Name] = &rc
	}
}

func createRaygoTransformations() {
	raygoTransforms = make(map[string]math.Matrix)
	for name, ytf := range yamlTransforms {
		tf, err := mapTransform(&ytf.TransformModel)
		if err != nil {
			panic(fmt.Sprintf("could not map transform '%v'", name))
		}
		raygoTransforms[name] = tf
	}
}

func mapTransform(ymlTransform *TransformModel) (math.Matrix, error) {
	switch ymlTransform.Type {
	case SCALING_TF:
		return math.Scaling(ymlTransform.X, ymlTransform.Y, ymlTransform.Z), nil
	case TRANSLATION_TF:
		return math.Translation(ymlTransform.X, ymlTransform.Y, ymlTransform.Z), nil
	case SHEARING_TF:
		return math.Shearing(ymlTransform.XY, ymlTransform.XZ, ymlTransform.YX, ymlTransform.YZ,
			ymlTransform.ZX, ymlTransform.ZY), nil
	case ROTATION_TF:
		var tf math.Matrix
		if ymlTransform.X != 0.0 {
			tf = math.Rotation_X(ymlTransform.X)
		} else if ymlTransform.Y != 0.0 {
			tf = math.Rotation_Y(ymlTransform.Y)
		} else {
			tf = math.Rotation_Z(ymlTransform.Z)
		}
		return tf, nil
	}
	return math.Matrix{}, fmt.Errorf("unknown transform type '%v'", ymlTransform.Type)
}

func createRaygoPatterns() {
	raygoPatterns = make(map[string]geometry.Pattern)

	for name, yp := range yamlCheckerPatterns {
		c1 := raygoColors[yp.ColorA]
		c2 := raygoColors[yp.ColorB]
		raygoPatterns[name] = geometry.CreateCheckerPattern(*c1, *c2)
	}

	for name, yp := range yamlGradientPatterns {
		c1 := raygoColors[yp.ColorA]
		c2 := raygoColors[yp.ColorB]
		raygoPatterns[name] = geometry.CreateGradientPattern(*c1, *c2)
	}

	for name, yp := range yamlRingPatterns {
		c1 := raygoColors[yp.ColorA]
		c2 := raygoColors[yp.ColorB]
		raygoPatterns[name] = geometry.CreateRingPattern(*c1, *c2)
	}

	for name, yp := range yamlStripePatterns {
		c1 := raygoColors[yp.ColorA]
		c2 := raygoColors[yp.ColorB]
		raygoPatterns[name] = geometry.CreateStripePattern(*c1, *c2)
	}
}

func createRaygoMaterials() {
	raygoMaterials = make(map[string]*geometry.Material)

	for name, ym := range yamlMaterials {
		m := geometry.DefaultMaterial()
		if ym.Ambient != nil {
			m.Ambient = *ym.Ambient
		}
		if ym.Diffuse != nil {
			m.Diffuse = *ym.Diffuse
		}
		if ym.Specular != nil {
			m.Specular = *ym.Specular
		}
		if ym.Shininess != nil {
			m.Shininess = *ym.Shininess
		}
		if ym.Transparency != nil {
			m.Transparency = *ym.Transparency
		}
		if ym.Reflective != nil {
			m.Reflective = *ym.Reflective
		}
		if ym.RefractiveIndex != nil {
			m.RefractiveIndex = *ym.RefractiveIndex
		}

		if ym.RawColor != nil {
			m.Color = mapColor(ym.RawColor)
		}

		if ym.Color != "" {
			m.Color = *raygoColors[ym.Color]
		}

		if ym.Pattern != "" {
			m.Pattern = raygoPatterns[ym.Pattern]
		}

		raygoMaterials[name] = &m
	}
}

func createRaygoShapes(directory string) {
	raygoShapes = make(map[string]geometry.Shape)

	createRaygoPlanes()
	createRaygoSpheres()
	createRaygoCubes()
	createRaygoCylinders()
	createRaygoCones()
	createRaygoTriangles()
	createRaygoObjects(directory)
	// groups need to be last, they can reference all other shapes
	createRaygoGroups()
}

func createRaygoPlanes() {
	for name, yp := range yamlPlanes {
		plane := geometry.CreatePlane()
		if yp.Transform != "" {
			plane.Transform = raygoTransforms[name]
		} else {
			plane.Transform = createTransformFromList(yp.Transforms)
		}

		if yp.Material != "" {
			plane.Material = *raygoMaterials[yp.Material]
		}

		raygoShapes[name] = plane
	}
}

func createRaygoSpheres() {
	for name, ys := range yamlSpheres {
		sphere := geometry.CreateSphere()

		if ys.Transform != "" {
			sphere.Transform = raygoTransforms[name]
		} else {
			sphere.Transform = createTransformFromList(ys.Transforms)
		}

		if ys.Material != "" {
			sphere.Material = *raygoMaterials[ys.Material]
		}

		raygoShapes[name] = sphere
	}
}

func createRaygoCubes() {
	for name, yc := range yamlCubes {
		cube := geometry.CreateCube()
		if yc.Transform != "" {
			cube.Transform = raygoTransforms[name]
		} else {
			cube.Transform = createTransformFromList(yc.Transforms)
		}

		if yc.Material != "" {
			cube.Material = *raygoMaterials[yc.Material]
		}

		raygoShapes[name] = cube
	}
}

func createRaygoCylinders() {
	for name, yc := range yamlCylinders {
		cylinder := geometry.CreateCylinder()
		if yc.Transform != "" {
			cylinder.Transform = raygoTransforms[name]
		} else {
			cylinder.Transform = createTransformFromList(yc.Transforms)
		}

		if yc.Material != "" {
			cylinder.Material = *raygoMaterials[yc.Material]
		}

		if yc.Minimum != nil {
			cylinder.Minimum = *yc.Minimum
		}
		if yc.Maximum != nil {
			cylinder.Maximum = *yc.Maximum
		}

		cylinder.Closed = yc.Closed
		raygoShapes[name] = cylinder
	}
}

func createRaygoCones() {
	for name, yc := range yamlCones {
		cone := geometry.CreateCone()
		if yc.Transform != "" {
			cone.Transform = raygoTransforms[name]
		} else {
			cone.Transform = createTransformFromList(yc.Transforms)
		}

		if yc.Material != "" {
			cone.Material = *raygoMaterials[yc.Material]
		}

		if yc.Minimum != nil {
			cone.Minimum = *yc.Minimum
		}
		if yc.Maximum != nil {
			cone.Maximum = *yc.Maximum
		}

		cone.Closed = yc.Closed
		raygoShapes[name] = cone
	}
}

func createRaygoTriangles() {
	for name, yt := range yamlTriangles {
		triangle := geometry.CreateTriangle(mapPoint(yt.P1), mapPoint(yt.P2), mapPoint(yt.P3))
		if yt.Transform != "" {
			triangle.Transform = raygoTransforms[name]
		} else {
			triangle.Transform = createTransformFromList(yt.Transforms)
		}

		if yt.Material != "" {
			triangle.Material = *raygoMaterials[yt.Material]
		}

		raygoShapes[name] = triangle
	}
}

func createRaygoObjects(directory string) {
	for name, yo := range yamlObjects {
		objData := obj.ParseFile(fmt.Sprintf("%v%v", directory, yo.File))
		objGroup := objData.ToGroup(true)

		if yo.Transform != "" {
			objGroup.Transform = raygoTransforms[name]
		} else {
			objGroup.Transform = createTransformFromList(yo.Transforms)
		}

		if yo.Material != "" {
			objGroup.Material = *raygoMaterials[yo.Material]
		}

		raygoShapes[name] = objGroup
	}
}

func createRaygoGroups() {
	for name, yg := range yamlGroups {
		group := geometry.EmptyGroup()
		if yg.Transform != "" {
			group.Transform = raygoTransforms[name]
		} else {
			group.Transform = createTransformFromList(yg.Transforms)
		}

		if yg.Material != "" {
			group.Material = *raygoMaterials[yg.Material]
		}

		for _, child := range yg.Children {
			// mark child as dependent on group
			childrenObjects[child] = struct{}{}
			// get raygo child and add to group
			raygoChild := raygoShapes[child]
			group.AddChild(raygoChild)
		}

		raygoShapes[name] = group
	}
}

func createTransformFromList(tfList []TransformModel) math.Matrix {
	result := math.IdentityMatrix()

	for _, tf := range tfList {
		mappedTf, err := mapTransform(&tf)
		if err != nil {
			panic(err)
		}
		result = result.MulM(mappedTf)
	}

	return result
}

func mapPoint(yamlPoint *PointModel) math.Point {
	return math.CreatePoint(yamlPoint.X, yamlPoint.Y, yamlPoint.Z)
}

func mapColor(yamlColor *ColorModel) math.Color {
	return math.CreateColor(math.BToF(yamlColor.R), math.BToF(yamlColor.G), math.BToF(yamlColor.B))
}
