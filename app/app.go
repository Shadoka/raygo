package app

import (
	"fmt"
	"os"
	"raygo/canvas"
	"raygo/parser"
	"slices"
	"strings"
)

func Run(args []string) {
	if fileFlagIndex := slices.Index(args, "-f"); fileFlagIndex != -1 {
		if len(args) <= fileFlagIndex+1 {
			panic("missing file path after -f flag")
		}
		filepath := args[fileFlagIndex+1]

		outputFilename := "default"
		if outputFlagIndex := slices.Index(args, "-o"); outputFlagIndex != -1 {
			if len(args) <= outputFlagIndex+1 {
				panic("missing file name after -o flag")
			}
			outputFilename = args[outputFlagIndex+1]
		}

		data, err := os.ReadFile(filepath)
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

		lastDirSep := strings.LastIndex(filepath, "/")
		dirpath := ""
		if lastDirSep != -1 {
			dirpath = filepath[:lastDirSep]
		}

		world := parser.CreateWorld(yml, dirpath)
		camera := parser.CreateCamera(yml)

		c := camera.Render(world, true)
		if len(c) == 1 {
			c[0].WritePPM(fmt.Sprintf("%v.ppm", outputFilename))
		} else {
			canvas.WriteGif(c, yml.Camera.Animation.Time, fmt.Sprintf("%v.gif", outputFilename))
		}
	}
}
