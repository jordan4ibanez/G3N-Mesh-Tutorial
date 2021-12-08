package engine

import (
	"fmt"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/math32"

	"github.com/g3n/engine/material"
	"github.com/g3n/engine/texture"
)

var debugDirection bool = false

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

	//positions are the "dots" which make up the mesh's vertex layout - this is a square
	positions.Append(
		-0.5, 0.5, 0.0, //top left
		-0.5, -0.5, 0.0, //bottom left
		0.5, -0.5, 0.0, //bottom right
		0.5, 0.5, 0.0, //top right
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
		0.0, 0.0, 1.0, //all face one direction (+Z coordinate on plane of Z)
		0.0, 0.0, 1.0, //this is used to optimize GPU lighting calculations
		0.0, 0.0, 1.0,

		0.0, 0.0, 1.0,
		0.0, 0.0, 1.0,
		0.0, 0.0, 1.0,
	)

	//uv is the texture mapping
	//this is a literal interpretation of the 2D texture that is in /textures/ from 0.0 to 1.0000~INF
	//the y on the UV map is actually inverted of your normal coordinate instincts if you commonly work with higher level engines
	//0 on the y is the top of the image
	//1 on the y is the bottom of the image
	//(y is the second float32 in this case)

	var uvShift bool = false

	if !uvShift {
		uvs.Append(
			0.0, 1.0, //bottom left
			0.0, 0.0, //top left
			1.0, 0.0, //top right
			1.0, 1.0, //bottom right
		)
	} else {
		uvs.Append(
			1.0, 1.0, //bottom right
			0.0, 1.0, //bottom left
			0.0, 0.0, //top left
			1.0, 0.0, //top right
		)
	}

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
	var myMaterial *material.Standard = material.NewStandard(math32.NewColor("White"))

	//turns the 2D myTexture.png image in /textures/ into a texture that OpenGL understands


	var myTexture *texture.Texture2D

	var error error

	if (debugDirection) {
		myTexture, error = texture.NewTexture2DFromImage("textures/up.png")

	} else {
		myTexture, error = texture.NewTexture2DFromImage("textures/myTexture.png")
	}

	//prints an error if there is one
	if error != nil {
		fmt.Println(error)
	}

	//this sets the texture filter to not "automatically smooth" and blur the texture, aka low poly voxel game look
	myTexture.SetMagFilter(gls.NEAREST)

	//we finally apply the texture to the material
	myMaterial.AddTexture(myTexture)

	//we are now done with the material object

	//we now are able to create a mesh from the data we told the computer to create from raw numbers/files
	myMesh = graphic.NewMesh(myGeometry, myMaterial)

	//finally we add the mesh into the scene
	scene.Add(myMesh)

	//if you do this a few million more times you have created Minetest
}

//this is a "getter" that the main.go file can use to import the myMesh object into
func GetMyMesh() *graphic.Mesh {
	return myMesh
}
