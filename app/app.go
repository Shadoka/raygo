package app

import (
	"fmt"
	"os"
	"raygo/parser"
	"slices"
)

func Run(args []string) {
	if fileFlagIndex := slices.Index(args, "-f"); fileFlagIndex != -1 {
		if len(args) <= fileFlagIndex+1 {
			panic("missing file path after -f flag")
		}
		data, err := os.ReadFile(args[fileFlagIndex+1])
		if err != nil {
			panic(err)
		}
		yml := parser.ParseYaml(string(data))
		validationResult := yml.Validate()
		validationResult = append(validationResult, parser.ValidateReferences(yml)...)
		if len(validationResult) != 0 {
			for i, vr := range validationResult {
				fmt.Printf("%v. %v\n", i, vr.Error())
			}
		}
	}
}
