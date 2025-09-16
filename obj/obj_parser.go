package obj

import (
	"fmt"
	"os"
	"raygo/geometry"
	"raygo/math"
	"slices"
	"strconv"
	"strings"
)

const VERTEX_PREFIX = "v"
const FACE_PREFIX = "f"
const GROUP_PREFIX = "g"

var currentGroup *ObjGroup

type ObjData struct {
	Vertices     []math.Point
	Faces        []*Face
	Groups       []*ObjGroup
	IgnoredLines int
}

type Face struct {
	VertIndices []int
}

type ObjGroup struct {
	Faces []*Face
}

func (o *ObjData) GetV(index int) math.Point {
	return o.Vertices[index-1]
}

func CreateObjData() *ObjData {
	return &ObjData{
		Vertices:     make([]math.Point, 0, 300),
		Faces:        make([]*Face, 0, 100),
		Groups:       make([]*ObjGroup, 0, 2),
		IgnoredLines: 0,
	}
}

func CreateFace(cap int) *Face {
	return &Face{
		VertIndices: make([]int, 0, cap),
	}
}

func CreateObjGroup() *ObjGroup {
	return &ObjGroup{
		Faces: make([]*Face, 0, 100),
	}
}

func (o *ObjData) PrintStats() {
	fmt.Printf("Vertices: %v\n", len(o.Vertices))
	fmt.Printf("Faces(root): %v\n", len(o.Faces))
	fmt.Printf("Groups: %v\n", len(o.Groups))
}

func (o *ObjData) ToGroup() *geometry.Group {
	root := geometry.EmptyGroup()

	for _, face := range o.Faces {
		for _, t := range face.ToTriangles(o) {
			root.AddChild(t)
		}
	}

	for _, objGroup := range o.Groups {
		grp := geometry.EmptyGroup()
		for _, face := range objGroup.Faces {
			for _, t := range face.ToTriangles(o) {
				grp.AddChild(t)
			}
		}
		root.AddChild(grp)
	}

	return root
}

func (f *Face) ToTriangles(o *ObjData) []*geometry.Triangle {
	triangles := make([]*geometry.Triangle, 0, 1)
	for i := 1; i < len(f.VertIndices)-1; i++ {
		p1 := o.GetV(f.VertIndices[0])
		p2 := o.GetV(f.VertIndices[i])
		p3 := o.GetV(f.VertIndices[i+1])
		triangles = append(triangles, geometry.CreateTriangle(p1, p2, p3))
	}
	return triangles
}

func ParseFile(objPath string) *ObjData {
	content, err := os.ReadFile(objPath)
	if err != nil {
		panic("cannot open obj file")
	}

	data := CreateObjData()
	ParseData(data, string(content))

	return data
}

func ParseData(objData *ObjData, data string) {
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		ParseLine(objData, &line)
	}
}

func ParseLine(objData *ObjData, line *string) {
	if strings.HasPrefix(*line, VERTEX_PREFIX) {
		processVertex(objData, line)
	} else if strings.HasPrefix(*line, FACE_PREFIX) {
		processFace(objData, line, currentGroup)
	} else if strings.HasPrefix(*line, GROUP_PREFIX) {
		currentGroup = CreateObjGroup()
		objData.Groups = append(objData.Groups, currentGroup)
	} else {
		objData.IgnoredLines += 1
	}
}

func processFace(objData *ObjData, line *string, currentGroup *ObjGroup) {
	faceComponents := strings.Split(*line, " ")
	faceComponents = slices.DeleteFunc(faceComponents, isEmptyString)
	face := CreateFace(len(faceComponents) - 1)
	for index := range faceComponents {
		if index == 0 {
			continue
		}
		// face format: vertexIndex/textureIndex/vertexNormal
		vertIndexString, _, _ := strings.Cut(faceComponents[index], "/")
		vertIndex, _ := strconv.Atoi(vertIndexString)
		face.VertIndices = append(face.VertIndices, vertIndex)
	}

	if currentGroup == nil {
		objData.Faces = append(objData.Faces, face)
	} else {
		currentGroup.Faces = append(currentGroup.Faces, face)
	}
}

func isEmptyString(s string) bool {
	return s == ""
}

func processVertex(objData *ObjData, line *string) {
	vertexComponents := strings.Split(*line, " ")
	vertexComponents = slices.DeleteFunc(vertexComponents, isEmptyString)
	if len(vertexComponents) != 4 {
		panic(fmt.Sprintf("a vertex line must consist of 4 elements: %v", vertexComponents))
	}

	vertx, err := getStringAsFloat(vertexComponents[1])
	if err != nil {
		fmt.Printf("current line: %s\n", *line)
		panic("")
	}
	verty, err := getStringAsFloat(vertexComponents[2])
	if err != nil {
		fmt.Printf("current line: %s\n", *line)
		panic("")
	}
	vertz, err := getStringAsFloat(vertexComponents[3])
	if err != nil {
		fmt.Printf("current line: %s\n", *line)
		panic("")
	}
	objData.Vertices = append(objData.Vertices,
		math.CreatePoint(vertx,
			verty,
			vertz))
}

func getStringAsFloat(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("cannot parse '%v' as float\n", s)
		return 0.0, fmt.Errorf("cannot parse '%v' as float\n", s)
	}
	return f, nil
}
