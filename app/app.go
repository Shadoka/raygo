package app

import (
	"fmt"
	"os"
	"raygo/canvas"
	"raygo/obj"
	"raygo/parser"
	"slices"
	"strings"
)

type FileType int

const (
	OBJ FileType = iota
	YAML
	UNKNOWN
)

func Run(args []string) {
	if fileFlagIndex := slices.Index(args, "-f"); fileFlagIndex != -1 {
		if len(args) <= fileFlagIndex+1 {
			panic("missing file path after -f flag")
		}
		filepath := args[fileFlagIndex+1]

		switch determineFileType(filepath) {
		case OBJ:
			handleObjStats(filepath)
		case YAML:
			handleRendering(args, filepath)
		case UNKNOWN:
			fmt.Printf("encountered input file with unfamiliar file ending: '%v'\n", filepath)
		}
	}
}

func handleObjStats(filepath string) {
	object := obj.ParseFile(filepath)
	object.PrintStats()
}

func handleRendering(args []string, filepath string) {
	outputFilename := getOutputFilename(args)
	writePng := checkPngFlag(args)

	yml := parseYamlFile(filepath)

	lastDirSep := strings.LastIndex(filepath, "/")
	dirpath := ""
	if lastDirSep != -1 {
		dirpath = filepath[:lastDirSep+1]
	}

	world := parser.CreateWorld(yml, dirpath)
	camera := parser.CreateCamera(yml)

	c := camera.Render(world, true)
	if len(c) == 1 {
		if writePng {
			c[0].WritePng(fmt.Sprintf("%v.png", outputFilename))
		} else {
			c[0].WritePPM(fmt.Sprintf("%v.ppm", outputFilename))
		}
	} else {
		canvas.WriteGif(c, yml.Camera.Animation.Time, fmt.Sprintf("%v.gif", outputFilename))
	}
}

func parseYamlFile(path string) *parser.YamlDescription {
	data, err := os.ReadFile(path)
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
		os.Exit(1)
	}

	return yml
}

func getOutputFilename(args []string) string {
	outputFilename := "default"
	if outputFlagIndex := slices.Index(args, "-o"); outputFlagIndex != -1 {
		if len(args) <= outputFlagIndex+1 {
			panic("missing file name after -o flag")
		}
		outputFilename = args[outputFlagIndex+1]
	}
	return outputFilename
}

func checkPngFlag(args []string) bool {
	if outputFlagIndex := slices.Index(args, "--png"); outputFlagIndex != -1 {
		return true
	}
	return false
}

func determineFileType(file string) FileType {
	if strings.HasSuffix(file, ".obj") {
		return OBJ
	} else if strings.HasSuffix(file, ".yaml") || strings.HasSuffix(file, ".yml") {
		return YAML
	} else {
		return UNKNOWN
	}
}
