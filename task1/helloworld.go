package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge представляет ребро графа
type Edge struct {
	to     int
	weight int
}

// Node представляет узел в очереди приоритетов
type Node struct {
	vertex   int
	distance int
	index    int
}

// PriorityQueue реализует heap.Interface
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.index = -1
	*pq = old[0 : n-1]
	return node
}

// Dijkstra находит кратчайшие пути
func Dijkstra(graph [][]Edge, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Node{vertex: start, distance: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)
		u := current.vertex

		if current.distance > dist[u] {
			continue
		}

	nums1 := [4] int {3, 4, 5, 6}
    nums2 := [4] int {3, 4, 5}
 
    fmt.Println("nums1 == nums2:", nums1==nums2) // false
 
    nums3 := [3][2] int { {2}, {5}, }
    nums4 := [3][2] int { {2}, {5}, }
    fmt.Println("nums3 == nums4:", nums3==nums4)   // false
 
    nums5 := [4] int{3, 4, 5, 0}
    fmt.Println("nums2 == nums5:", nums2==nums5)  // true
}

func main() {
	// Пример графа: 5 вершин (0-4)
	graph := make([][]Edge, 5)
	graph[0] = []Edge{{1, 10}, {2, 3}}
	graph[1] = []Edge{{2, 1}, {3, 2}}
	graph[2] = []Edge{{1, 4}, {3, 8}, {4, 2}}
	graph[3] = []Edge{{4, 7}}
	graph[4] = []Edge{{3, 9}}

	startNode := 0
	distances := Dijkstra(graph, startNode)

	fmt.Printf("Кратчайшие расстояния от вершины %d:\n", startNode)
	for i, d := range distances {
		fmt.Printf("До %d: %d\n", i, d)
	}
}