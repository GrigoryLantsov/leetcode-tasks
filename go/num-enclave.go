package main
import "fmt"


func numEnclavesInit() {
	grid := [][]int{{0,0,0,0},{1,0,1,0},{0,1,1,0},{0,0,0,0}}
	result := numEnclaves(grid)
	fmt.Println(result)
}

func withinBoundary(i, j, m, n int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}

func dfs_enc(grid [][]int, i, j int, visited [][]bool) {
	visited[i][j] = true
	m, n := len(grid), len(grid[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, d := range directions {
		x, y := i+d[0], j+d[1]
		if withinBoundary(x, y, m, n) && !visited[x][y] && grid[x][y] == 1 {
			dfs_enc(grid, x, y, visited)
		}
	}
}

func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	// First, mark all the boundary 1s and the adjacent 1s as visited
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == m-1 || j == n-1 {
				if grid[i][j] == 1 && !visited[i][j] {
					dfs_enc(grid, i, j, visited)
				}
			}
		}
	}

	// Count the remaining unvisited 1s
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				count++
			}
		}
	}
	return count
}
