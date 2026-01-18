package progress

import (
	"fmt"
	"strings"
)

// Step prints a progress step message
func Step(message string) {
	fmt.Printf("%s\n", message)
}

// SetFrameInfo sets up frame tracking information
func SetFrameInfo(currentFrame, totalFrames int) {
	percentage := int((float64(currentFrame) / float64(totalFrames)) * 100)
	bar := createProgressBar(percentage, 30)
	fmt.Printf("\rRendering frame %d of %d [%s] %d%%", currentFrame, totalFrames, bar, percentage)
	if currentFrame == totalFrames {
		fmt.Printf("\n")
	}
}

// TotalFrames prints how many frames will be rendered
func TotalFrames(count int) {
	fmt.Printf("Total frames to render: %d\n", count)
}

// Complete marks the entire process as complete
func Complete(totalTime string) {
	fmt.Printf("\nRendering complete! Total time: %s\n", totalTime)
}

// createProgressBar creates a visual progress bar
func createProgressBar(percentage, width int) string {
	filled := (percentage * width) / 100
	empty := width - filled
	return "[" + strings.Repeat("=", filled) + strings.Repeat(" ", empty) + "]"
}
