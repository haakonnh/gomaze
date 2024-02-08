# gomaze
A Go project which implements different MST algorithms to generate and search algorithms solve mazes.

## Algorithms
Prims and Kruskals algorithms are implemented to generate the maze. Prim grows the maze from the currently visited cells randomly, whilst Kruskal picks a random wall to break down in the grid and connects the maze over time.
DFS is currently the only algorithm implemented to solve the maze. It goes depth first to try to find the end cell, if not, it backtracks to try another path. 

## Dependencies
Ebitengine is used for the rendering of the maze. 
