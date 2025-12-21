package parser

type YamlDescription struct {
	Colors []ColorModel `name:"colors"`
}

type ColorModel struct {
	Name string `yaml:"name"`
	R    int    `yaml:"r"`
	G    int    `yaml:"g"`
	B    int    `yaml:"b"`
}
