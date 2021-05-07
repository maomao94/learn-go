package main

import (
	"fmt"

	"github.com/chentaihan/container/queue"
)

// 深度优先
func getProvince(citysConnected [][]int) int {
	citys := len(citysConnected)
	visited := make([]bool, citys)
	province := 0 // 计数器
	for i := 0; i < citys; i++ {
		if !visited[i] {
			// 深度优先
			dfs(i, citys, visited, citysConnected)
			province++
		}
	}
	return province
}

// 广度优先
func bfs(citysConnected [][]int) int {
	citys := len(citysConnected)
	visited := make([]bool, citys)
	province := 0 // 计数器
	queueLink := queue.NewQueueLink()
	for i := 0; i < citys; i++ {
		if !visited[i] {
			queueLink.Enqueue(i)
			for !queueLink.Empty() {
				k, _ := queueLink.Dequeue()
				visited[k.(int)] = true
				for j := 0; j < citys; j++ {
					if citysConnected[i][j] == 1 && !visited[j] {
						queueLink.Enqueue(j)
					}
				}
			}
			province++
		}
	}
	return province
}

func dfs(i int, citys int, visited []bool, citysConnected [][]int) {
	for j := 0; j < citys; j++ {
		if citysConnected[i][j] == 1 && !visited[j] {
			visited[j] = true
			dfs(j, citys, visited, citysConnected)
		}
	}
}

// 省份数量
func main() {
	fmt.Println(getProvince([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}})) // 2
	fmt.Println(getProvince([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}})) // 3
	fmt.Println(bfs([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))         // 2
	fmt.Println(bfs([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))         // 3
}
