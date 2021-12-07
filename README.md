# G3N-Mesh-Tutorial
 This is a basic tutorial which shows you how to create a mesh in G3N from scratch using only a .png texture.

 This tutorial was built on how I learned LWJGL3 from this tutorial: https://lwjglgamedev.gitbooks.io/3d-game-development-with-lwjgl/content/

 If you see any ways to improve this tutorial, please let me know.

 If you are brand new to Go I recommend: [Tech With Tim Golang Tutorial](https://www.youtube.com/playlist?list=PLzMcBGfZo4-mtY_SE3HuzQJzuj4VlUG0q)

I'm going to assume you're using VSCODE, Code, or CODE OSS, and your IDE's terminal is already in the directory of the program in this tutorial.

Let's get started.

First of all, this is how this tutorial looks when you run:

    $ go build && ./G3N-Mesh-Tutorial

On Linux or

    $ go build; .\G3N-Mesh-Tutorial.exe

On Windows.

![Wow, a square!](https://raw.githubusercontent.com/jordan4ibanez/G3N-Mesh-Tutorial/main/screenshots/program.png)

As you can see, this is extremely basic. But every program starts off extremely basic. It only matters where your imagination can and will take you with it.

**Question:** Okay what am I looking at? Isn't this from the game Crafting Mine or something?

**Answer:** No, this is a simple .png image used as a texture material, created from 2 tris with a helper coordinate axis in G3N labeled "A beautiful mesh".

Okay maybe that is a bit complex to unload to someone new to this, let me make it a little bit simpler.

Verbatem from Google: A mesh is a collection of vertices, edges and faces that define the shape of a 3D object. 

Sadly they left off the part where a mesh usually has a material plopped on it, but we'll get to that.

### Step 1, Vertex Positions:

Let's start off with the simplest thing in that definition, edges.

In OpenGL an "edge" is a vertex position, a literal point in 3D space which is assembled into a list, otherwise known as an array. A vertex position can be by itself in a list, but that wouldn't be visible, or helpful. So we assemble them into an array. [Here is a visualization of what is happening in line 32 of meshBuilder.go.](https://github.com/jordan4ibanez/G3N-Mesh-Tutorial/blob/main/engine/meshBuilder.go#L32)

![Wow, dots, amazing](https://raw.githubusercontent.com/jordan4ibanez/G3N-Mesh-Tutorial/main/screenshots/vertexPosition.png)

Truly incredible, isn't it? Well, sarcasm aside, it is actually quite incredible that it is that easy to tell the computer where to put those points in 3D space.

Basically all we're doing is saying, hey program, I want these dots here, thanks.

Then they are stored as such.

### Step 2, Indices:

In OpenGL, you have your vertex positions, but those alone are not very helpful. We need a way to tell the graphics card how to connect these dots, quite literally. That's where indices come into play. In OpenGL these dots are connected together in what is known as "tris". The term tri, literally means triangle. 

When these dots are connected together, your graphics card can fill them in utilizing it's pixel shader cores, but that is beyond the scope of this tutorial. Just know, you need three points to connect the dots, and keep them as triangles. This will allow you to, let's say, put a color to the triangle created between those points. Or maybe you could put a texture there instead. It is up to you. But for simplicity sake, in this tutorial we will be texturing our tris.

Before we get into texturing, [I will show you our two quads that we created at line 40 of meshBuilder.go.](https://github.com/jordan4ibanez/G3N-Mesh-Tutorial/blob/main/engine/meshBuilder.go#L40)

The first image is tri 1, the second is tri 2. Pretty simple isn't it?

![TRI 1](https://raw.githubusercontent.com/jordan4ibanez/G3N-Mesh-Tutorial/main/screenshots/tri1.png)

![TRI 2](https://raw.githubusercontent.com/jordan4ibanez/G3N-Mesh-Tutorial/main/screenshots/tri2.png)

And now we show them together as it's written in the code!

![A SQUARE!](https://raw.githubusercontent.com/jordan4ibanez/G3N-Mesh-Tutorial/main/screenshots/bothTris.png)

A truly beautiful square. OpenGL has built in magic and knows how to assemble these together, so really, the only thing you have to worry about is making those vertex positions turn into a square using the indices formed into two tris. OpenGL knows what to do automatically. Truly incredible!

### Step 3, Normals:

**Question:** I thought all of this code was normal?

**Answer:** Well, normals are simply the term used to define the front of a mesh.

Think of it as the face of the mesh, literally. Normals are used to calculate the visibility of parts of your mesh so that the graphics card does not have to do unnecessary work when rendering your game! It's also used for lighting calculations. You can technically work with OpenGL without using normals, but it becomes a horrible mess and I'd recommend not doing that. [If you want to read more about it, you can click this link.](https://en.wikipedia.org/wiki/Back-face_culling) It's actually fascinating.

But onto the code. [At line 50 of meshBuilder.go](https://github.com/jordan4ibanez/G3N-Mesh-Tutorial/blob/main/engine/meshBuilder.go#L50), you can see we are telling the normals to face on the positive Z axis (it's defined as XYZ). Our plane (or square if you want to call it that) is utilizing X and Y for it's 2D alignment in the 3D space so the Z axis becomes the default plane in which the pixels are rendered on due to how 3D space works. 

I'm going to show you what I mean using the actual program.

![Pretty normal looking](https://raw.githubusercontent.com/jordan4ibanez/G3N-Mesh-Tutorial/main/screenshots/normalsShowcase.png)

As you can see, the +Z face is now rendering outwards where it should be. All is good!

But what happens when we turn this mesh around?

![Oh no, it's gone](https://raw.githubusercontent.com/jordan4ibanez/G3N-Mesh-Tutorial/main/screenshots/normalsInverseShowcase.png)

As you can see, backface culling has kicked in. The GPU is basically stopping itself from rendering things it doesn't need to exactly as we told it to.

**Question:** And how is this helpful at all? Wouldn't you want to see this mesh from all angles?

**Answer:** Maybe in this extremely basic tutorial. But now imagine we had millions of tris in this mesh, all textured, all at different angles. If the GPU had to render them all no matter if it didn't need to, your game would basically slow to a crawl. This is why backface culling was created in OpenGL, to stop the GPU from getting worked unnecessarily.

This also has an extreme impact on complex scenes that utilize lots of lighting. If you think just rendering with the pixel shader is taxing in the scenario I just described, imagine if you had a few hundred lights in the scene. Your GPU would basically melt. Backface culling is truly a marvel. 

This is why normals are very important.

### Step 4, UVs:

Now I know what you're thinking, UV as in UV light. But UV are actually the texture mapping. UV is a coordinate system like XYZ. In OpenGL when you hear the term UV, just think "texture map". This is a very simple way to sum that up, [but you can read about it here if you want to know more.](https://en.wikipedia.org/wiki/UV_mapping) I'm just going to use texture mapping from here on out, just know that I am referring to UVs.

Texture mapping is one of the most essential parts of meshes. They allow you to do complex model wrapping to make your model look more detailed than it is. You can do some truly crazy things with texture mapping. But for our tutorial scenario, we have made a square that has a texture on it. [You can see this being applied to the geometry on line 73 of meshBuilder.go.](https://github.com/jordan4ibanez/G3N-Mesh-Tutorial/blob/main/engine/meshBuilder.go#L73)

Something very important I would like to get out of the way first. The texture coordinate system on the Y axis (up and down) of the texure map is inverted. Left is 0, right is 1, up is 0, down is 1. Very simple, but this could get extremely confusing if you ever forget this.
Hopefully this example helps explain what I mean.

![This guy was very happy to be drawn](https://raw.githubusercontent.com/jordan4ibanez/G3N-Mesh-Tutorial/main/screenshots/flippedUVExample.png)

You want your texture definition to start at Vertex Position 0 and end at 3. Basically make it so that it goes to the Vertex Positions in order, 0,1,2,3. You can see this on the line 73 link above. It goes from 0, to 3. If you mess up your UV mapping, you can keep brute forcing it if you really want to, it just becomes a pain in more complex programs to do this.

If your image is not mapped the way you want it to be, you can do a UV shift. This is simply done by taking the bottom coordinate and moving it to the top. This will either move the orientation 90 degrees clockwise or counter clockwise. On line 16, you can turn on the debugDirection boolean to change the texture to a simple "up" texture.](https://github.com/jordan4ibanez/G3N-Mesh-Tutorial/blob/main/engine/meshBuilder.go#L16) Then, [on line 70, you can change the uvShift boolean to true](https://github.com/jordan4ibanez/G3N-Mesh-Tutorial/blob/main/engine/meshBuilder.go#L70) to make it turn 90 degrees counter clockwise. I recommend changing debugDirection and doing a run/rebuild, then changing uvShift so you can see it in action.



