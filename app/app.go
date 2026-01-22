package app

import (
	"fmt"
	"os"
	"path/filepath"
	"raygo/canvas"
	"raygo/obj"
	"raygo/parser"
	"raygo/progress"
	"slices"
	"strings"
	"time"
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
		fp := args[fileFlagIndex+1]

		switch determineFileType(fp) {
		case OBJ:
			handleObjStats(fp)
		case YAML:
			handleRendering(args, fp)
		case UNKNOWN:
			fmt.Printf("encountered input file with unfamiliar file ending: '%v'\n", fp)
		}
	}
}

func handleObjStats(fp string) {
	object := obj.ParseFile(fp)
	object.PrintStats()
}

func handleRendering(args []string, fp string) {
	startTime := time.Now()

	outputFilename := getOutputFilename(args)
	writePng := checkPngFlag(args)
	antialias := checkAntialiasFlag(args)

	progress.Step("Parsing Yaml")
	yml := parseYamlFile(fp)

	absolutePath, err := filepath.Abs(fp)
	if err != nil {
		panic("unable to get absolute file path for yaml file")
	}
	lastDirSep := strings.LastIndex(absolutePath, string(os.PathSeparator))
	dirpath := ""
	if lastDirSep != -1 {
		dirpath = absolutePath[:lastDirSep+1]
	}

	progress.Step("Creating Scene from Yaml")
	world := parser.CreateWorld(yml, dirpath)
	camera := parser.CreateCamera(yml)
	camera.Antialias = antialias

	c := camera.Render(world, true)
	progress.Step("Writing output file")
	if len(c) == 1 {
		if writePng {
			c[0].WritePng(fmt.Sprintf("%v.png", outputFilename))
		} else {
			c[0].WritePPM(fmt.Sprintf("%v.ppm", outputFilename))
		}
	} else {
		canvas.WriteGif(c, yml.Camera.Animation.Time, fmt.Sprintf("%v.gif", outputFilename))
	}
	elapsed := time.Since(startTime)
	progress.Complete(fmt.Sprintf("%.2f seconds", elapsed.Seconds()))
}

func parseYamlFile(path string) *parser.YamlDescription {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	yml := parser.ParseYaml(string(data))
	progress.Step("Validating Yaml")
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

func checkAntialiasFlag(args []string) bool {
	if antialiasFlagIndex := slices.Index(args, "--aa"); antialiasFlagIndex != -1 {
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
