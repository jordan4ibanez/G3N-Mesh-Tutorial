package engine

import (
	"fmt"
	"log"
	"os"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/math32"

	//"github.com/g3n/engine/gls"
	//"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/texture"
)

var myMesh *graphic.Mesh

func DebugTest(scene *core.Node) {

	//first we shall create the geometric data for G3N to utilize

	//create a blank geometry pointer
	var myGeometry *geometry.Geometry = geometry.NewGeometry()

	// Create buffers
	positions := math32.NewArrayF32(0, 16)
	indices := math32.NewArrayU32(0, 16)
	normals := math32.NewArrayF32(0, 16)
	uvs := math32.NewArrayF32(0, 16)

	//positions are the "dots" which make up the mesh's vertex layout
	positions.Append(
		//tri 1
		-0.5, 0.5, 0.0, //top left
		-0.5, -0.5, 0.0, //bottom left
		0.5, 0.5, 0.0, //top right

		//tri 2
		0.5, 0.5, 0.0, //top right
		-0.5, -0.5, 0.0, //bottom left
		0.5, -0.5, 0.0, //bottom right
	)
	//indices are the "lines" which draw the connection between the vertex dots
	//they are drawn as triangles, even to make quads for gpu optimizations
	indices.Append(
		//tri 1
		0, 1, 3,

		//tri 2
		3, 1, 2,
	)

	//normals are the direction that the textures are facing
	//in this case, the "plane" (2 tris) are facing +Z
	normals.Append(
		0.0, 0.0, 1.0, //all face one direction (+Z coordinate on plane of Z); invert by switching to -1.0
		0.0, 0.0, 1.0,
		0.0, 0.0, 1.0,
		0.0, 0.0, 1.0,
		0.0, 0.0, 1.0,
		0.0, 0.0, 1.0,
	)

	//uv is the texture mapping
	uvs.Append(
		0.0, 0.0,
		0.0, 1.0,
		1.0, 1.0,
		1.0, 0.0,
	)

	//apply the geometric indice data to the geometry object
	myGeometry.SetIndices(indices)

	//apply the position data to the geometry object and inline set it's data attribute to Vertex Position
	myGeometry.AddVBO(gls.NewVBO(positions).AddAttrib(gls.VertexPosition))

	//apply the normals data to the geometry object and inline set it's data attribute to Vertex Normal
	myGeometry.AddVBO(gls.NewVBO(normals).AddAttrib(gls.VertexNormal))

	//apply the uvs (texture coordinates) data to the geometry object and inline set it's data attribute to Vertex Texcoord (Texture Coordinate)
	myGeometry.AddVBO(gls.NewVBO(uvs).AddAttrib(gls.VertexTexcoord))

	//we are now done with the geometric structure of the mesh

	//we are now moving into the material of the mesh

	//create a blank material object - this NewBasic()Material object is able to intake textures
	var myMaterial *material.Material = &material.NewBasic().Material

	//turns the 2D myTexture.png image in /textures/ into a texture that OpenGL understands
	myTexture, error := texture.NewTexture2DFromImage("textures/myTexture.png")

	if error != nil {
		fmt.Println("HEY SOMETHING GOT FUCKED HERE")
		fmt.Println(error)
		fmt.Println("HEY SOMETHING GOT FUCKED HERE")
	}

	//we finally apply the texture to the material
	myMaterial.AddTexture(myTexture)

	//we are now done with the material object

	//we now are able to create a mesh from the data we told the computer to create from raw numbers/files
	myMesh = graphic.NewMesh(myGeometry, myMaterial)

	//finally we add the mesh into the scene
	scene.Add(myMesh)
}

//this is a "getter" that the main.go file can use to import the myMesh object into
func GetMyMesh() *graphic.Mesh {
	return myMesh
}
