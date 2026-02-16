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

const VERTEX_PREFIX = "v "
const FACE_PREFIX = "f "
const GROUP_PREFIX = "g "
const NORMAL_PREFIX = "vn "
const TEXTURE_PREFIX = "vt "

var currentGroup *ObjGroup

type ObjData struct {
	Vertices           []math.Point
	Faces              []*Face
	Normals            []math.Vector
	TextureCoordinates []math.Point
	Groups             []*ObjGroup
	IgnoredLines       int
}

type Face struct {
	VertIndices    []int
	TextureIndices []int
	NormalIndices  []int
}

type ObjGroup struct {
	Faces []*Face
}

func (o *ObjData) GetV(index int) math.Point {
	return o.Vertices[index-1]
}

func (o *ObjData) GetN(index int) math.Vector {
	return o.Normals[index-1]
}

func (o *ObjData) GetT(index int) math.Vector {
	return o.TextureCoordinates[index-1]
}

func CreateObjData() *ObjData {
	return &ObjData{
		Vertices:     make([]math.Point, 0, 300),
		Faces:        make([]*Face, 0, 100),
		Normals:      make([]math.Vector, 0, 100),
		Groups:       make([]*ObjGroup, 0, 2),
		IgnoredLines: 0,
	}
}

func CreateFace(cap int) *Face {
	return &Face{
		VertIndices:    make([]int, 0, cap),
		TextureIndices: make([]int, 0, cap),
		NormalIndices:  make([]int, 0, cap),
	}
}

func CreateObjGroup() *ObjGroup {
	return &ObjGroup{
		Faces: make([]*Face, 0, 100),
	}
}

func (o *ObjData) PrintStats() {
	objBounds := o.ToGroup(true).Bounds()
	fmt.Printf("Vertices: %v\n", len(o.Vertices))
	fmt.Printf("Faces(root): %v\n", len(o.Faces))
	fmt.Printf("Normals: %v\n", len(o.Normals))
	fmt.Printf("Texture Coordinates: %v\n", len(o.TextureCoordinates))
	fmt.Printf("Groups: %v\n", len(o.Groups))
	fmt.Printf("Bounds:\n")
	fmt.Printf("\tMin: %v\n", objBounds.Minimum.ToString())
	fmt.Printf("\tMax: %v\n", objBounds.Maximum.ToString())
}

func (o *ObjData) ToGroup(preCalcBB bool) *geometry.Group {
	root := geometry.EmptyGroup()
	root.GetMaterial().SetShininess(50.0) // ???

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

	if preCalcBB {
		root.Bounds()
	}

	return root
}

func (f *Face) ToTriangles(o *ObjData) []*geometry.Triangle {
	triangles := make([]*geometry.Triangle, 0, 1)
	for i := 1; i < len(f.VertIndices)-1; i++ {
		p1 := o.GetV(f.VertIndices[0])
		p2 := o.GetV(f.VertIndices[i])
		p3 := o.GetV(f.VertIndices[i+1])

		triangle := geometry.CreateTriangle(p1, p2, p3)
		if len(f.NormalIndices) > 0 {
			n1 := o.GetN(f.NormalIndices[0])
			n2 := o.GetN(f.NormalIndices[i])
			n3 := o.GetN(f.NormalIndices[i+1])
			triangle.AddSmoothingInformation(n1, n2, n3)
		}
		if len(f.TextureIndices) > 0 {
			t1 := o.GetT(f.TextureIndices[0])
			t2 := o.GetT(f.TextureIndices[i])
			t3 := o.GetT(f.TextureIndices[i+1])
			triangle.AddTextureInformation(t1, t2, t3)
		}
		triangles = append(triangles, triangle)
	}
	return triangles
}

func ParseFile(objPath string) *ObjData {
	content, err := os.ReadFile(objPath)
	if err != nil {
		panic(fmt.Sprintf("cannot open obj file: '%v'", objPath))
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
	} else if strings.HasPrefix(*line, NORMAL_PREFIX) {
		processNormal(objData, line)
	} else if strings.HasPrefix(*line, TEXTURE_PREFIX) {
		processTextureCoordinates(objData, line)
	} else if strings.HasPrefix(*line, GROUP_PREFIX) {
		currentGroup = CreateObjGroup()
		objData.Groups = append(objData.Groups, currentGroup)
	} else {
		objData.IgnoredLines += 1
	}
}

func processNormal(objData *ObjData, line *string) {
	normalComponents := strings.Split(*line, " ")
	normalComponents = slices.DeleteFunc(normalComponents, isEmptyString)
	normalx := getStringAsFloat(normalComponents[1])
	normaly := getStringAsFloat(normalComponents[2])
	normalz := getStringAsFloat(strings.TrimSuffix(normalComponents[3], "\r"))

	objData.Normals = append(objData.Normals,
		math.CreateVector(normalx, normaly, normalz))
}

func processFace(objData *ObjData, line *string, currentGroup *ObjGroup) {
	faceComponents := strings.Split(*line, " ")
	faceComponents = slices.DeleteFunc(faceComponents, isEmptyString)
	face := CreateFace(len(faceComponents) - 1)
	for index := range faceComponents {
		if index == 0 {
			continue
		}
		vertexIndex, textureIndex, normalIndex := extractFaceIndices(faceComponents[index])
		face.VertIndices = append(face.VertIndices, vertexIndex)
		if textureIndex != -1 {
			face.TextureIndices = append(face.TextureIndices, textureIndex)
		}
		if normalIndex != -1 {
			face.NormalIndices = append(face.NormalIndices, normalIndex)
		}
	}

	if currentGroup == nil {
		objData.Faces = append(objData.Faces, face)
	} else {
		currentGroup.Faces = append(currentGroup.Faces, face)
	}
}

func processTextureCoordinates(objData *ObjData, line *string) {
	textureComponents := strings.Split(*line, " ")
	textureComponents = slices.DeleteFunc(textureComponents, isEmptyString)
	// only u is mandatory, v and w default to 0
	textureU := getStringAsFloat(textureComponents[1])
	textureV := 0.0
	textureW := 0.0
	if len(textureComponents) > 2 {
		textureV = getStringAsFloat(textureComponents[2])
	}
	if len(textureComponents) > 3 {
		textureW = getStringAsFloat(textureComponents[3])
	}

	objData.TextureCoordinates = append(objData.TextureCoordinates, math.CreatePoint(textureU, textureV, textureW))
}

// input => 1/3/5
func extractFaceIndices(face string) (int, int, int) {
	// face format: vertexIndex/textureIndex/vertexNormal
	indices := strings.Split(face, "/")
	vertexIndex := -1
	textureIndex := -1
	normalIndex := -1
	for i, stringIndex := range indices {
		if i == 0 {
			vertexIndex, _ = strconv.Atoi(stringIndex)
		} else if i == 1 && stringIndex != "" {
			textureIndex, _ = strconv.Atoi(stringIndex)
		} else if i == 2 {
			normalIndex, _ = strconv.Atoi(stringIndex)
		}
	}
	return vertexIndex, textureIndex, normalIndex
}

func isEmptyString(s string) bool {
	return s == "" || s == "\r"
}

func processVertex(objData *ObjData, line *string) {
	vertexComponents := strings.Split(*line, " ")
	vertexComponents = slices.DeleteFunc(vertexComponents, isEmptyString)
	if len(vertexComponents) != 4 {
		panic(fmt.Sprintf("a vertex line must consist of 4 elements: %v", vertexComponents))
	}

	vertx := getStringAsFloat(vertexComponents[1])
	verty := getStringAsFloat(vertexComponents[2])
	vertz := getStringAsFloat(strings.TrimSuffix(vertexComponents[3], "\r"))
	objData.Vertices = append(objData.Vertices,
		math.CreatePoint(vertx,
			verty,
			vertz))
}

func getStringAsFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
