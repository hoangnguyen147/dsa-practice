package main

func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    m, n := len(grid), len(grid[0])
    
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    nums := 0
    
    bfs := func(r, c int) {
        visited[r][c] = true
        queue := [][2]int{{r, c}}
        drn := [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

        for len(queue) > 0 {
            nr, nc := queue[0][0], queue[0][1]
            queue = queue[1:]
            for _, d := range drn {
                dr, dc := nr + d[0], nc + d[1]
                if dr >= 0 && dr < m && dc >= 0 && dc < n {
                    if grid[dr][dc] == '1' && !visited[dr][dc] {
                        visited[dr][dc] = true
                        queue = append(queue, [2]int{dr, dc})
                    }
                }
            }
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' && !visited[i][j] {
                nums++
                bfs(i, j)
            }
        }
    }
    
    return nums
}
