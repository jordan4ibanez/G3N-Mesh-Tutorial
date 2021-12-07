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

A truly beautiful square. OpenGL has built in magic to automatically know how to assemble these together, so really, the only thing you have to worry about is making those vertex positions turn into a square using the indices. OpenGL knows what to do automatically. Truly incredible to be honest with you!


