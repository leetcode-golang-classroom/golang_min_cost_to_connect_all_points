package sol

import "container/heap"

type AdjacentNode struct {
	Cost, Point int
}

type AdjacentMinHeap []AdjacentNode

func (h *AdjacentMinHeap) Len() int {
	return len(*h)
}
func (h *AdjacentMinHeap) Less(i, j int) bool {
	return (*h)[i].Cost < (*h)[j].Cost
}
func (h *AdjacentMinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *AdjacentMinHeap) Push(val interface{}) {
	*h = append(*h, val.(AdjacentNode))
}
func (h *AdjacentMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func abs(value int) int {
	if value > 0 {
		return value
	}
	return -value
}
func minCostConnectPoints(points [][]int) int {
	// make adjacency map
	n := len(points)
	adjacencyMap := make(map[int]AdjacentMinHeap, n)
	for i := 0; i < n; i++ {
		point1 := points[i]
		for j := i + 1; j < n; j++ {
			point2 := points[j]
			dist := abs(point1[0]-point2[0]) + abs(point1[1]-point2[1])
			adjacencyMap[i] = append(adjacencyMap[i], AdjacentNode{Cost: dist, Point: j})
			adjacencyMap[j] = append(adjacencyMap[j], AdjacentNode{Cost: dist, Point: i})
		}
	}

	totalCost := 0
	visit := make(map[int]struct{})
	priorityQueue := &AdjacentMinHeap{AdjacentNode{Cost: 0, Point: 0}}
	heap.Init(priorityQueue)
	// Prim's algorithm
	for len(visit) < n {
		node := heap.Pop(priorityQueue).(AdjacentNode)
		if _, exist := visit[node.Point]; exist {
			continue
		}
		totalCost += node.Cost
		visit[node.Point] = struct{}{}
		adjList := adjacencyMap[node.Point]
		for _, adjNode := range adjList {
			if _, exist := visit[adjNode.Point]; !exist {
				heap.Push(priorityQueue, adjNode)
			}
		}
	}

	return totalCost
}
