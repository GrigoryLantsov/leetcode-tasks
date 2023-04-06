package main
import (
	"fmt"
)
func closedIslandInit() {
	grid := [][]int{{1,1,1,1,1,1,1,0},{1,0,0,0,0,1,1,0},{1,0,1,0,1,1,1,0},{1,0,0,0,0,1,0,1},{1,1,1,1,1,1,1,0}}
	numClosedIslands := closedIsland(grid)
	fmt.Println(numClosedIslands)
}

func closedIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	count := 0

	// Initialize all border cells to 1 to mark as visited
	for i := 0; i < len(grid); i++ {
		dfs(grid, i, 0)
		dfs(grid, i, len(grid[0])-1)
	}
	for j := 0; j < len(grid[0]); j++ {
		dfs(grid, 0, j)
		dfs(grid, len(grid)-1, j)
	}

	// Check all inner cells
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if grid[i][j] == 0 {
				count++
				dfs(grid, i, j)
			}
		}
	}

	return count
}
func dfs(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 1 {
		return
	}

	grid[i][j] = 1
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}