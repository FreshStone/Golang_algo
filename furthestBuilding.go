package main

import (
	"fmt"
)

func main() {
	tests := [][]int{
		{4, 2, 7, 6, 9, 14, 12},
		{4, 12, 2, 7, 3, 18, 20, 3, 19},
		{14, 3, 19, 3},
	}
	bricks := []int{5, 10, 17}
	ladders := []int{1, 2, 0}
	ans := []int{4, 7, 3}
	for i := 0; i < len(tests); i++ {
		if ans[i] != furthestBuilding(tests[i], bricks[i], ladders[i]) {
			fmt.Println("error output", furthestBuilding(tests[i], bricks[i], ladders[i]), "expected", ans[i])
		} else {
			fmt.Println("correct")
		}
	}
}

/* leetcode 1642
maintain a priority queque(min_heap of top l(ladders) height differences) and
a heap_total counter (summation of all values in heap) and check
 if (bricks_to_use-heap_total > bricks) == false than move forward
*/
func furthestBuilding(heights []int, bricks int, ladders int) int {
	if len(heights) == 1 {
		return 0
	}
	if ladders > len(heights)-2 {
		return len(heights) - 1
	}
	var i, id, dif, j, heap_total, heap_len, bricks_to_use int
	min_heap := make([]int, ladders)
	min := func(a, b int) int {
		if min_heap[a] < min_heap[b] {
			return a
		}
		return b
	}
	add_to_heap := func() {
		if heap_len < ladders {
			heap_total += dif
		} else if min_heap[0] < dif { //pop top val from heap
			heap_total += dif - min_heap[0]
			heap_len -= 1
			min_heap[0] = min_heap[heap_len]
			j = 0
			for 2*j+2 < heap_len {
				id = min(2*j+1, 2*j+2)
				if min_heap[j] < min_heap[id] {
					break
				}
				min_heap[j], min_heap[id] = min_heap[id], min_heap[j]
				j = id
			}
			if (2*j+1) < heap_len && min_heap[2*j+1] < min_heap[j] {
				min_heap[j], min_heap[2*j+1] = min_heap[2*j+1], min_heap[j]
			}
		} else {
			return
		}
		//push val to heap
		j = heap_len
		heap_len += 1
		for j > 0 {
			if min_heap[(j-1)/2] > dif {
				min_heap[j] = min_heap[(j-1)/2]
				j = (j - 1) / 2
			} else {
				break
			}
		}
		min_heap[j] = dif
	}
	for i = 1; i < len(heights); i++ {
		if heights[i] > heights[i-1] {
			dif = heights[i] - heights[i-1]
			bricks_to_use += dif
			if ladders > 0 {
				add_to_heap()
			}
			if bricks_to_use-heap_total > bricks {
				return i - 1
			}
		}
	}
	return len(heights) - 1
}
