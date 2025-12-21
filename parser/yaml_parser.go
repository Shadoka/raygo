package parser

import (
	"log"

	"github.com/goccy/go-yaml"
)

func ParseYaml(content string) *YamlDescription {
	description := YamlDescription{}
	if err := yaml.Unmarshal([]byte(content), &description); err != nil {
		log.Fatalf("could not unmarshal yaml: %v", err)
	}
	return &description
}
