package progress

import (
	"fmt"
	"strings"
	"sync"
)

// Progress represents a progress tracker for rendering operations
type Progress struct {
	mu            sync.Mutex
	currentStep   string
	frameNumber   int
	totalFrames   int
	showFrameInfo bool
}

// NewProgress creates a new progress tracker
func NewProgress() *Progress {
	return &Progress{
		currentStep:   "",
		frameNumber:   0,
		totalFrames:   0,
		showFrameInfo: false,
	}
}

// Step prints a progress step message
func (p *Progress) Step(message string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.currentStep = message
	fmt.Printf("%s\n", message)
}

// SetFrameInfo sets up frame tracking information
func (p *Progress) SetFrameInfo(currentFrame, totalFrames int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.frameNumber = currentFrame
	p.totalFrames = totalFrames
	p.showFrameInfo = true
	percentage := int((float64(currentFrame) / float64(totalFrames)) * 100)
	bar := p.createProgressBar(percentage, 30)
	fmt.Printf("\rRendering frame %d of %d [%s] %d%%", currentFrame, totalFrames, bar, percentage)
	if currentFrame == totalFrames {
		fmt.Printf("\n")
	}
}

// TotalFrames prints how many frames will be rendered
func (p *Progress) TotalFrames(count int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.totalFrames = count
	fmt.Printf("Total frames to render: %d\n", count)
}

// Complete marks the entire process as complete
func (p *Progress) Complete(totalTime string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	fmt.Printf("\nRendering complete! Total time: %s\n", totalTime)
}

// createProgressBar creates a visual progress bar
func (p *Progress) createProgressBar(percentage, width int) string {
	filled := (percentage * width) / 100
	empty := width - filled
	return "[" + strings.Repeat("=", filled) + strings.Repeat(" ", empty) + "]"
}
