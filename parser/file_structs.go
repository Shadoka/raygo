package parser

type YamlDescription struct {
	Colors    []ColorModel         `yaml:"colors"`
	Materials []NamedMaterialModel `yaml:"materials"`
	Patterns  PatternContainer     `yaml:"patterns"`
	Scene     SceneContainer       `yaml:"scene"`
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
	Name       string            `yaml:"name"`
	ColorA     string            `yaml:"colorA"`
	ColorB     string            `yaml:"colorB"`
	Transforms []*TransformModel `yaml:"transforms"`
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

type SceneContainer struct {
	Planes []*PlaneModel `yaml:"planes"`
}

type CommonSceneObject struct {
	Name       string            `yaml:"name"`
	Parent     string            `yaml:"parent"`
	Material   string            `yaml:"material"`
	Transform  string            `yaml:"transform"`
	Transforms []*TransformModel `yaml:"transforms"`
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
