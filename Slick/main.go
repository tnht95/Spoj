//https://www.spoj.com/problems/UCV2013H/cstart=20
package main

import (
	"bufio"
	"fmt"
	"os"
)

type cell struct {
	y int
	x int
}

func solve(n int, m int, arr [][]int) {
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	visited := make([][]bool, n)
	result := make([]int, 250*250)
	total := 0

	for i := range visited {
		visited[i] = make([]bool, m)
	}

	for i := range arr {
		for j := range arr[i] {
			count := 0
			queue := []cell{cell{i, j}}
			for len(queue) > 0 {
				cur := dequeue(&queue)
				if !visited[cur.y][cur.x] && arr[cur.y][cur.x] == 1 {
					visited[cur.y][cur.x] = true
					count++
					for k := range dx {
						newX := cur.x + dx[k]
						newY := cur.y + dy[k]
						if newX >= 0 && newX < m && newY >= 0 && newY < n && !visited[newY][newX] && arr[newY][newX] == 1 {
							queue = append(queue, cell{newY, newX})
						}
					}
				}
			}
			if count > 0 {
				result[count]++
				total++
			}
		}
	}

	fmt.Println(total)
	for i := range result {
		if result[i] != 0 {
			fmt.Fprintf(os.Stdout, "%d %d\n", i, result[i])
		}
	}
}

func dequeue(q *[]cell) cell {
	c := (*q)[0]
	(*q) = (*q)[1:]
	return c
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanf(r, "%d %d\n", &n, &m)
	for n > 0 {
		arr := make([][]int, n)
		for i := range arr {
			arr[i] = make([]int, m)
			for j := range arr[i] {
				fmt.Fscanf(r, "%d", &arr[i][j])
			}
			fmt.Fscanln(r)
		}
		solve(n, m, arr)
		fmt.Fscanf(r, "%d %d\n", &n, &m)
	}
}
