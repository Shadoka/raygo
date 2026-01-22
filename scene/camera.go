package scene

import (
	gomath "math"
	"raygo/canvas"
	g "raygo/geometry"
	"raygo/math"
	"raygo/progress"
	"sync"
)

type Camera struct {
	Hsize            int
	Vsize            int
	FieldOfView      float64
	Transform        math.Matrix
	HalfWidth        float64
	HalfHeight       float64
	PixelSize        float64
	Animation        *CameraAnimation
	Position         CameraPosition
	PositionStates   []CameraPosition
	Antialias        bool
	ColorCache       ColorCache
	InverseTransform *math.Matrix // <-- invalidate after rendering of frame
}

// only circular motion around point for now
type CameraAnimation struct {
	FullMotionRadians float64
	MovementTime      float64
	TargetFps         float64
}

type CameraPosition struct {
	From math.Point
	To   math.Point
	Up   math.Vector
}

type ColorCache struct {
	mu               sync.Mutex
	CanvasColorCache map[math.Point]*math.Color // <-- invalidate after rendering of frame
}

func (cc *ColorCache) Set(p math.Point, color *math.Color) {
	cc.mu.Lock()
	defer cc.mu.Unlock()
	cc.CanvasColorCache[p] = color
}

func (cc *ColorCache) Get(p math.Point) *math.Color {
	cc.mu.Lock()
	defer cc.mu.Unlock()
	return cc.CanvasColorCache[p]
}

func CreateCamera(hsize int, vsize int, fov float64) *Camera {
	c := &Camera{
		Hsize:            hsize,
		Vsize:            vsize,
		FieldOfView:      fov,
		Transform:        math.IdentityMatrix(),
		InverseTransform: nil,
		ColorCache: ColorCache{
			CanvasColorCache: make(map[math.Point]*math.Color, 0),
		},
		Antialias: false,
	}
	c.calculateCameraProperties()

	return c
}

func CreateCameraAnimation(radians float64, time float64, fps float64) *CameraAnimation {
	ca := &CameraAnimation{
		FullMotionRadians: radians,
		MovementTime:      time,
		TargetFps:         fps,
	}
	return ca
}

func CreateCameraPosition(from math.Point, to math.Point, up math.Vector) CameraPosition {
	cp := CameraPosition{
		From: from,
		To:   to,
		Up:   up,
	}
	return cp
}

func (cp *CameraPosition) Equals(other *CameraPosition) bool {
	return other != nil &&
		cp.From.Equals(other.From) &&
		cp.To.Equals(other.To) &&
		cp.Up.Equals(other.Up)
}

func (c *Camera) calculateCameraProperties() {
	halfView := gomath.Tan(c.FieldOfView / 2.0)
	aspect := float64(c.Hsize) / float64(c.Vsize)

	if aspect >= 1.0 {
		c.HalfWidth = halfView
		c.HalfHeight = halfView / aspect
	} else {
		c.HalfWidth = halfView * aspect
		c.HalfHeight = halfView
	}

	c.PixelSize = (c.HalfWidth * 2.0) / float64(c.Hsize)
}

// the initial camera position has to be set before calling this function
func (c *Camera) createAnimationStates() {
	if c.Animation == nil {
		states := make([]CameraPosition, 0, 1)
		states = append(states, c.Position)
		c.PositionStates = states
		return
	}

	totalFrames := c.Animation.MovementTime * c.Animation.TargetFps
	// if I want to enable camera smoothing radianFrameDelta needs to be variable
	// -1 because we start with the non-rotated starting position
	radianFrameDelta := c.Animation.FullMotionRadians / (totalFrames - 1.0)
	totalFramesInt := int(totalFrames)

	states := make([]CameraPosition, 0, totalFramesInt)
	states = append(states, c.Position)

	for frameIndex := range totalFramesInt {
		if frameIndex == 0 {
			continue
		}
		previousFrom := states[frameIndex-1].From
		currentFrameFrom := math.RotateAroundPoint(previousFrom, radianFrameDelta, c.Position.To)
		states = append(states, CreateCameraPosition(currentFrameFrom, c.Position.To, c.Position.Up))
	}

	c.PositionStates = states
}

func (c *Camera) RayForPixel(x int, y int) g.Ray {
	coordinate := c.calculateWorldCoordinateWithOffset(float64(x), float64(y), 0.5, 0.5)

	// using the camera matrix, transform the canvas point and the origin,
	// and then compute the rays direction vector.
	// (remember that the canvas is at z = -1)
	pixel := c.GetInverseTransform().MulT(coordinate)
	origin := c.GetInverseTransform().MulT(math.CreatePoint(0.0, 0.0, 0.0))
	direction := pixel.Subtract(origin).Normalize()

	return g.CreateRay(origin, direction)
}

func (c *Camera) RayForCoordinate(coordinate math.Point) g.Ray {
	// using the camera matrix, transform the canvas point and the origin,
	// and then compute the rays direction vector.
	pixel := c.GetInverseTransform().MulT(coordinate)
	origin := c.GetInverseTransform().MulT(math.CreatePoint(0.0, 0.0, 0.0))
	direction := pixel.Subtract(origin).Normalize()

	return g.CreateRay(origin, direction)
}

func (c *Camera) SetTransform(tf math.Matrix) {
	c.Transform = tf
}

func (c *Camera) Render(w *World, multithreaded bool) []*canvas.Canvas {
	c.createAnimationStates()
	totalFrames := len(c.PositionStates)
	progress.TotalFrames(totalFrames)

	images := make([]*canvas.Canvas, 0, len(c.PositionStates))
	for frameIndex, currentPosition := range c.PositionStates {
		c.Position = currentPosition
		c.InverseTransform = nil
		if multithreaded {
			images = append(images, c.RenderMultithreaded(w, c.Hsize/2))
		} else {
			images = append(images, c.RenderSinglethreaded(w))
		}
		progress.SetFrameInfo(frameIndex+1, totalFrames)
	}
	return images
}

func (c *Camera) RenderSinglethreaded(w *World) *canvas.Canvas {
	c.Transform = math.ViewTransform(c.Position.From, c.Position.To, c.Position.Up)
	canv := canvas.CreateCanvas(c.Hsize, c.Vsize)

	for y := range c.Vsize {
		for x := range c.Hsize {
			r := c.RayForPixel(x, y)
			color := w.ColorAt(r, MAX_REFLECTION_LIMIT)
			if c.Antialias {
				relevantPixelColors := c.getCornerColors(w, x, y)
				relevantPixelColors = append(relevantPixelColors, color)
				color = getMeanColor(relevantPixelColors)
			}
			canv.WritePixel(x, y, color)
		}
	}

	return &canv
}

func (c *Camera) RenderMultithreaded(w *World, workerThreads int) *canvas.Canvas {
	c.Transform = math.ViewTransform(c.Position.From, c.Position.To, c.Position.Up)
	var wg sync.WaitGroup
	canv := canvas.CreateCanvas(c.Hsize, c.Vsize)

	rowsPerWorker := c.Vsize / workerThreads
	remainingRows := c.Vsize % workerThreads

	for worker := range workerThreads {
		from := rowsPerWorker * worker
		to := from + rowsPerWorker
		if worker == workerThreads-1 {
			to = to + remainingRows
		}

		wg.Add(1)

		go func() {
			defer wg.Done()
			c.renderPartially(from, to, w, &canv)
		}()
	}

	wg.Wait()
	return &canv
}

func (c *Camera) renderPartially(fromY int, toY int, w *World, cv *canvas.Canvas) {
	for y := fromY; y < toY; y++ {
		for x := range c.Hsize {
			r := c.RayForPixel(x, y)
			color := w.ColorAt(r, MAX_REFLECTION_LIMIT)
			if c.Antialias {
				relevantPixelColors := c.getCornerColors(w, x, y)
				relevantPixelColors = append(relevantPixelColors, color)
				color = getMeanColor(relevantPixelColors)
			}
			cv.WritePixel(x, y, color)
		}
	}
}

/**
* Gets the colors the 4 corners of the pixel.
 */
func (c *Camera) getCornerColors(w *World, x int, y int) []math.Color {
	xFloat := float64(x)
	yFloat := float64(y)
	points := make([]math.Point, 0)
	points = append(points, c.calculateWorldCoordinateWithOffset(xFloat, yFloat, 0.0, 0.0))
	points = append(points, c.calculateWorldCoordinateWithOffset(xFloat, yFloat, 0.0, 1.0))
	points = append(points, c.calculateWorldCoordinateWithOffset(xFloat, yFloat, 1.0, 0.0))
	points = append(points, c.calculateWorldCoordinateWithOffset(xFloat, yFloat, 1.0, 1.0))

	colors := make([]math.Color, 0)
	for _, corner := range points {
		if cornerColor := c.ColorCache.Get(corner); cornerColor == nil {
			cornerRay := c.RayForCoordinate(corner)
			color := w.ColorAt(cornerRay, MAX_REFLECTION_LIMIT)
			c.ColorCache.Set(corner, &color)
			colors = append(colors, color)
		} else {
			colors = append(colors, *cornerColor)
		}
	}

	return colors
}

func getMeanColor(colors []math.Color) math.Color {
	r := 0.0
	g := 0.0
	b := 0.0

	for _, color := range colors {
		r += color.X
		g += color.Y
		b += color.Z
	}

	colorCount := float64(len(colors))
	return math.CreateColor(r/colorCount, g/colorCount, b/colorCount)
}

func (c *Camera) calculateWorldCoordinateWithOffset(x float64, y float64, xOff float64, yOff float64) math.Point {
	// the offset from the edge of the canvas to the pixels center
	xOffset := (float64(x) + xOff) * c.PixelSize
	yOffset := (float64(y) + yOff) * c.PixelSize

	// the untransformed coordinates of the pixel in world space
	// (remember that the camera looks toward -z, so +x is to the *left*)
	return math.CreatePoint(c.HalfWidth-xOffset, c.HalfHeight-yOffset, -1.0)
}

func (c *Camera) GetInverseTransform() math.Matrix {
	if c.InverseTransform != nil {
		return *c.InverseTransform
	}

	inverse := c.Transform.Inverse()
	c.InverseTransform = &inverse
	return *c.InverseTransform
}
