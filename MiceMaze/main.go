//https://www.spoj.com/SSCMTA/problems/MICEMAZE/
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

var (
	maze [][]int
)

type node struct {
	v    int
	cost int
}

type minHeap []node

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h minHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(node))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func solve(exit int, timer int, n int) int {
	result := 0
	time := make([]int, n+1)
	h := &minHeap{node{exit, 0}}
	heap.Init(h)

	for i := range time {
		time[i] = math.MaxInt32
	}
	time[exit] = 0

	for h.Len() > 0 {
		cur := heap.Pop(h).(node)
		for i := range maze[cur.v] {
			t := cur.cost + maze[cur.v][i]
			if maze[cur.v][i] > 0 && t < time[i] {
				time[i] = t
				heap.Push(h, node{i, t})
			}
		}
	}

	for i := range time {
		if time[i] <= timer {
			result++
		}
	}

	return result
}

func main() {
	r := bufio.NewReader(os.Stdin)

	var n, exit, timer, m int
	fmt.Fscanf(r, "%d\n", &n)
	fmt.Fscanf(r, "%d\n", &exit)
	fmt.Fscanf(r, "%d\n", &timer)
	fmt.Fscanf(r, "%d\n", &m)

	maze = make([][]int, n+1)
	for i := range maze {
		maze[i] = make([]int, n+1)
	}

	//directed graph
	//maze present which nodes(column) can go to this node(row)
	//0 1 2 3
	//1 0 4 0 -> only node 2 can go to 1 with cost 4
	//2	0 0 0
	//3 0 0 0

	for i := 0; i < m; i++ {
		var from, to, t int
		fmt.Fscanf(r, "%d %d %d\n", &from, &to, &t)
		maze[to][from] = t
	}

	fmt.Fprintf(os.Stdout, "%d\n", solve(exit, timer, n))
}
