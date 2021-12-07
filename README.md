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

Question: Okay what am I looking at? Isn't this from the game Crafting Mine or something?

Answer: No, this is a simple .png image used as a texture material, created from 2 tris with a helper coordinate axis in G3N labeled "A beautiful mesh".

Okay maybe that is a bit complex to unload to someone new to this, let me make it a little bit simpler.

Verbatem from Google: A mesh is a collection of vertices, edges and faces that define the shape of a 3D object. 

### Step 1, Indices:

Let's start off with the simplest thing in that definition, edges.

In OpenGL an "edge" is an indice, a literal point in 3D space which is assembled into a list, otherwise known as an array. An indice can be by itself in a list, but that wouldn't be visible.


