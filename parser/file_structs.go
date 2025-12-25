package parser

import "fmt"

type YamlDescription struct {
	Colors     []NamedColorModel     `yaml:"colors"`
	Materials  []NamedMaterialModel  `yaml:"materials"`
	Transforms []NamedTransformModel `yaml:"transforms"`
	Patterns   PatternContainer      `yaml:"patterns"`
	Scene      SceneContainer        `yaml:"scene"`
	Light      LightModel            `yaml:"light"`
	Camera     CameraModel           `yaml:"camera"`
	Width      int                   `yaml:"width"`
	Height     int                   `yaml:"height"`
}

type ColorModel struct {
	R int `yaml:"r"`
	G int `yaml:"g"`
	B int `yaml:"b"`
}

type NamedColorModel struct {
	ColorModel `yaml:",inline"`
	Name       string `yaml:"name"`
}

type PatternContainer struct {
	Checker  []CheckerPatternModel  `yaml:"checker"`
	Ring     []RingPatternModel     `yaml:"ring"`
	Gradient []GradientPatternModel `yaml:"gradient"`
	Stripe   []StripePatternModel   `yaml:"stripe"`
}

// embedding struct
type DualColorPattern struct {
	Name       string           `yaml:"name"`
	ColorA     string           `yaml:"colorA"`
	ColorB     string           `yaml:"colorB"`
	Transforms []TransformModel `yaml:"transforms"`
}

type CheckerPatternModel struct {
	DualColorPattern `yaml:",inline"`
}

type GradientPatternModel struct {
	DualColorPattern `yaml:",inline"`
}

type RingPatternModel struct {
	DualColorPattern `yaml:",inline"`
}

type StripePatternModel struct {
	DualColorPattern `yaml:",inline"`
}

type TransformModel struct {
	Type string  `yaml:"type"`
	X    float64 `yaml:"x"`
	Y    float64 `yaml:"y"`
	Z    float64 `yaml:"z"`
}

type NamedTransformModel struct {
	TransformModel `yaml:",inline"`
	Name           string `yaml:"name"`
}

type MaterialModel struct {
	Color           string  `yaml:"color"`
	Pattern         string  `yaml:"pattern"`
	Ambient         float64 `yaml:"ambient"`
	Diffuse         float64 `yaml:"diffuse"`
	Specular        float64 `yaml:"specular"`
	Shininess       float64 `yaml:"shininess"`
	Reflective      float64 `yaml:"reflective"`
	Transparency    float64 `yaml:"transparency"`
	RefractiveIndex float64 `yaml:"refractiveIndex"`
}

type NamedMaterialModel struct {
	MaterialModel `yaml:",inline"`
	Name          string `yaml:"name"`
}

type PointModel struct {
	X float64
	Y float64
	Z float64
}

type VectorModel struct {
	PointModel `yaml:",inline"`
}

type SceneContainer struct {
	Planes    []PlaneModel    `yaml:"planes"`
	Cubes     []CubeModel     `yaml:"cubes"`
	Spheres   []SphereModel   `yaml:"spheres"`
	Groups    []GroupModel    `yaml:"groups"`
	Triangles []TriangleModel `yaml:"triangles"`
	Cylinders []CylinderModel `yaml:"cylinders"`
	Cones     []ConeModel     `yaml:"cones"`
	Objects   []ObjectModel   `yaml:"objects"`
}

type CommonSceneObject struct {
	Name string `yaml:"name"`
	//Parent     string           `yaml:"parent"`
	Material   string           `yaml:"material"`
	Transform  string           `yaml:"transform"`
	Transforms []TransformModel `yaml:"transforms"`
}

type PlaneModel struct {
	CommonSceneObject `yaml:",inline"`
}

type SphereModel struct {
	CommonSceneObject `yaml:",inline"`
}

type CubeModel struct {
	CommonSceneObject `yaml:",inline"`
}

type GroupModel struct {
	CommonSceneObject `yaml:",inline"`
	Children          []string `yaml:"children"`
}

type TriangleModel struct {
	CommonSceneObject `yaml:",inline"`
	P1                *PointModel `yaml:"p1"`
	P2                *PointModel `yaml:"p2"`
	P3                *PointModel `yaml:"p3"`
}

type CylinderModel struct {
	CommonSceneObject `yaml:",inline"`
	Minimum           float64 `yaml:"min"`
	Maximum           float64 `yaml:"max"`
	Closed            bool    `yaml:"closed"`
}

type ConeModel struct {
	CylinderModel `yaml:",inline"`
}

type ObjectModel struct {
	CommonSceneObject `yaml:",inline"`
	File              string `yaml:"file"`
}

type LightModel struct {
	Position  *PointModel `yaml:"p"`
	Intensity *ColorModel `yaml:"intensity"`
}

type CircularCameraAnimation struct {
	Radians float64 `yaml:"radians"`
	Time    float64 `yaml:"timeSec"`
}

type CameraModel struct {
	From      *PointModel              `yaml:"from"`
	To        *PointModel              `yaml:"to"`
	LookAt    string                   `yaml:"lookAt"`
	Up        *VectorModel             `yaml:"up"`
	Animation *CircularCameraAnimation `yaml:"animation"`
}

// validation

func (c *CameraModel) validate() []error {
	valResult := make([]error, 0)

	if c.From == nil {
		valResult = append(valResult, fmt.Errorf("camera requires a 'from' position"))
	}

	if c.To == nil && c.LookAt == "" {
		valResult = append(valResult, fmt.Errorf("camera requires either a valid 'to' or 'lookAt' reference"))
	}

	if c.Up == nil {
		valResult = append(valResult, fmt.Errorf("camera requires an 'up' vector"))
	}

	if c.Animation != nil {
		valResult = append(valResult, c.Animation.validate()...)
	}

	return valResult
}

func (anim *CircularCameraAnimation) validate() []error {
	valResult := make([]error, 0)

	if anim.Time <= 0.0 {
		valResult = append(valResult, fmt.Errorf("the camera animation requires a 'time' value of > 0.0"))
	}

	return valResult
}

func (l *LightModel) validate() []error {
	valResult := make([]error, 0)

	if l.Position == nil {
		valResult = append(valResult, fmt.Errorf("light requires a 'p' (position) value"))
	}

	if l.Intensity == nil {
		valResult = append(valResult, fmt.Errorf("light requires an 'intensity' field for its color"))
	}

	return valResult
}

func (sceneObject *CommonSceneObject) validate() []error {
	valResult := make([]error, 0)

	if sceneObject.Name == "" {
		// I would need yaml line to be more specific
		valResult = append(valResult, fmt.Errorf("scene objects require a non empty 'name' field"))
	}

	return valResult
}

func (obj *ObjectModel) validate() []error {
	valResult := make([]error, 0)

	if obj.File == "" {
		valResult = append(valResult, fmt.Errorf("object '%v' requires a 'file' field from which to load the OBJ", obj.Name))
	}

	valResult = append(valResult, obj.CommonSceneObject.validate()...)

	return valResult
}

func (c *CylinderModel) validate() []error {
	valResult := make([]error, 0)

	if c.Minimum == c.Maximum {
		err := fmt.Errorf("min of scene object '%v' is equal to max", c.Name)
		valResult = append(valResult, err)
	}

	if c.Minimum > c.Maximum {
		err := fmt.Errorf("min(%v) of scene object '%v' is greater than max(%v)", c.Minimum, c.Name, c.Maximum)
		valResult = append(valResult, err)
	}

	valResult = append(valResult, c.CommonSceneObject.validate()...)

	return valResult
}

func (c *ConeModel) validate() []error {
	return c.CylinderModel.validate()
}

func (t *TriangleModel) validate() []error {
	valResult := make([]error, 0)

	if t.P1 == nil {
		valResult = append(valResult, fmt.Errorf("triangle '%v' requires field 'p1'", t.Name))
	}

	if t.P2 == nil {
		valResult = append(valResult, fmt.Errorf("triangle '%v' requires field 'p2'", t.Name))
	}

	if t.P3 == nil {
		valResult = append(valResult, fmt.Errorf("triangle '%v' requires field 'p3'", t.Name))
	}

	valResult = append(valResult, t.CommonSceneObject.validate()...)

	return valResult
}

func (g *GroupModel) validate() []error {
	valResult := make([]error, 0)

	if len(g.Children) == 0 {
		valResult = append(valResult, fmt.Errorf("group '%v' requires children", g.Name))
	}

	valResult = append(valResult, g.CommonSceneObject.validate()...)

	return valResult
}

func (c *CubeModel) validate() []error {
	return c.CommonSceneObject.validate()
}

func (s *SphereModel) validate() []error {
	return s.CommonSceneObject.validate()
}

func (p *PlaneModel) validate() []error {
	return p.CommonSceneObject.validate()
}

func (scene *SceneContainer) validate() []error {
	valResult := make([]error, 0)

	for _, p := range scene.Planes {
		valResult = append(valResult, p.validate()...)
	}

	for _, c := range scene.Cubes {
		valResult = append(valResult, c.validate()...)
	}

	for _, s := range scene.Spheres {
		valResult = append(valResult, s.validate()...)
	}

	for _, g := range scene.Groups {
		valResult = append(valResult, g.validate()...)
	}

	for _, t := range scene.Triangles {
		valResult = append(valResult, t.validate()...)
	}

	for _, c := range scene.Cylinders {
		valResult = append(valResult, c.validate()...)
	}

	for _, c := range scene.Cones {
		valResult = append(valResult, c.validate()...)
	}

	for _, o := range scene.Objects {
		valResult = append(valResult, o.validate()...)
	}

	return valResult
}

func (m *NamedMaterialModel) validate() []error {
	valResult := make([]error, 0)

	if m.Name == "" {
		// I would need yaml line to be more specific
		valResult = append(valResult, fmt.Errorf("named materials require a non empty 'name' field"))
	}

	return valResult
}

func (t *TransformModel) validate() []error {
	valResult := make([]error, 0)

	if t.Type == "" {
		// I would need yaml line to be more specific
		valResult = append(valResult, fmt.Errorf("transforms require a non empty 'type' field"))
	}

	return valResult
}

func (t *NamedTransformModel) validate() []error {
	valResult := make([]error, 0)

	if t.Name == "" {
		// I would need yaml line to be more specific
		valResult = append(valResult, fmt.Errorf("named transforms require a non empty 'name' field"))
	}

	valResult = append(valResult, t.TransformModel.validate()...)

	return valResult
}

func (c *NamedColorModel) validate() []error {
	valResult := make([]error, 0)

	if c.Name == "" {
		// I would need yaml line to be more specific
		valResult = append(valResult, fmt.Errorf("named colors require a non empty 'name' field"))
	}

	return valResult
}

func (p *DualColorPattern) validate() []error {
	valResult := make([]error, 0)

	if p.Name == "" {
		// I would need yaml line to be more specific
		valResult = append(valResult, fmt.Errorf("patterns require a non empty 'name' field"))
	}

	if p.ColorA == "" {
		valResult = append(valResult, fmt.Errorf("pattern '%v' requires a non empty 'colorA' field", p.Name))
	}

	if p.ColorB == "" {
		valResult = append(valResult, fmt.Errorf("pattern '%v' requires a non empty 'colorB' field", p.Name))
	}

	for _, t := range p.Transforms {
		valResult = append(valResult, t.validate()...)
	}

	return valResult
}

func (p *GradientPatternModel) validate() []error {
	return p.DualColorPattern.validate()
}

func (p *StripePatternModel) validate() []error {
	return p.DualColorPattern.validate()
}

func (p *CheckerPatternModel) validate() []error {
	return p.DualColorPattern.validate()
}

func (p *RingPatternModel) validate() []error {
	return p.DualColorPattern.validate()
}

func (patterns *PatternContainer) validate() []error {
	valResult := make([]error, 0)

	for _, p := range patterns.Checker {
		valResult = append(valResult, p.validate()...)
	}

	for _, p := range patterns.Ring {
		valResult = append(valResult, p.validate()...)
	}

	for _, p := range patterns.Gradient {
		valResult = append(valResult, p.validate()...)
	}

	for _, p := range patterns.Stripe {
		valResult = append(valResult, p.validate()...)
	}

	return valResult
}

func (yml *YamlDescription) Validate() []error {
	valResult := make([]error, 0)

	if yml.Width <= 0 {
		valResult = append(valResult, fmt.Errorf("field 'width' is required with a value greater than 0"))
	}

	if yml.Height <= 0 {
		valResult = append(valResult, fmt.Errorf("field 'height' is required with a value greater than 0"))
	}

	for _, c := range yml.Colors {
		valResult = append(valResult, c.validate()...)
	}

	for _, m := range yml.Materials {
		valResult = append(valResult, m.validate()...)
	}

	for _, t := range yml.Transforms {
		valResult = append(valResult, t.validate()...)
	}

	valResult = append(valResult, yml.Patterns.validate()...)
	valResult = append(valResult, yml.Scene.validate()...)
	valResult = append(valResult, yml.Light.validate()...)
	valResult = append(valResult, yml.Camera.validate()...)

	return valResult
}
