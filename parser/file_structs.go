package parser

type YamlDescription struct {
	Colors    []ColorModel         `yaml:"colors"`
	Materials []NamedMaterialModel `yaml:"materials"`
	Patterns  PatternContainer     `yaml:"patterns"`
	Scene     SceneContainer       `yaml:"scene"`
	Light     LightModel           `yaml:"light"`
	Camera    CameraModel          `yaml:"camera"`
	Width     int                  `yaml:"width"`
	Height    int                  `yaml:"height"`
}

type ColorModel struct {
	Name string `yaml:"name"`
	R    int    `yaml:"r"`
	G    int    `yaml:"g"`
	B    int    `yaml:"b"`
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
	Name       string           `yaml:"name"`
	Parent     string           `yaml:"parent"`
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
