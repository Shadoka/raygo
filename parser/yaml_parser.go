package parser

import (
	"fmt"
	"log"

	"github.com/goccy/go-yaml"
)

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
		err := fmt.Errorf("cannot resolve material '%v' for scene object '%v'", sceneObject.Material, sceneObject.Name)
		valResult = append(valResult, err)
	}
	if sceneObject.Transform != "" && yamlTransforms[sceneObject.Transform] == nil {
		err := fmt.Errorf("cannot resolve transform '%v' for scene object '%v'", sceneObject.Transform, sceneObject.Name)
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
