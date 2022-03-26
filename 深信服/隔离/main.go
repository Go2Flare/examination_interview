package main

/*
一个区域范围，有部分新冠病人， 需要用障碍物将其隔离
请求出至少需要做为隔离的障碍物
grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]

*/
import "fmt"

func main() {
	//grid := [][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}}
	grid := [][]int{{1, 0, 0, 0, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 1, 1, 0}}
	for _, v := range grid {
		for _, v1 := range v {
			fmt.Print(v1, "\t")
		}
		fmt.Println()
	}
	fmt.Println("--------------------------------")
	res := getClosedRegion(grid)
	for _, v := range grid {
		for _, v1 := range v {
			fmt.Print(v1, "\t")
		}
		fmt.Println()
	}
	fmt.Printf("需设立%v个障碍物\n", res)
}
var (
	dx = [4]int{0, 1, 0, -1}
	dy = [4]int{1, 0, -1, 0}
)
var n,m int
func getClosedRegion(grid [][]int) (res int) {

	m = len(grid)
	n = len(grid[0])
	var bfs func(x, y int) int
	bfs = func(x, y int) int {
		count := 0
		grid[x][y] = 3
		lockDown(grid, x, y, &count)

		queue := [][]int{}
		queue = append(queue, []int{x, y})

		for i := 0; i < len(queue); i++ {
			e := queue[i]
			for j := 0; j < 4; j++ {
				//新的转向
				mx, my := e[0]+dx[j], e[1]+dy[j]
				//注意要用新的坐标
				if mx >= 0 && mx < m && my >= 0 && my < n && grid[mx][my] == 0 {
					grid[mx][my] = 3
					queue = append(queue, []int{mx, my})
					lockDown(grid, mx, my, &count)
				}
			}
		}
		return count
	}
	for i, v := range grid {
		for j, v1 := range v {
			if v1 == 0 {
				res += bfs(i, j)
			}
		}
	}
	return
}

func lockDown(grid [][]int, x,y int, count *int){
	for k := 0; k < 4; k++ {
		//如果现在所在的是病人，需要在周围设立障碍物
		nx, ny := x+dx[k], y+dy[k]
		if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
			*count++
			grid[nx][ny] = 2
		}
	}
}
