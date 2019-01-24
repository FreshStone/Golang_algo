package main

import (
	"fmt"
	"math"
)

func main(){
	key := 2
	tree := []int{1,2,3,4,-1,-1,-1,5,-1,6} // breath first representation (-1 is null node)
//	tree := []int{1,2,3,-1,-1,-1,4,-1,5,6}
	l := len(tree)
	if l == 0 || tree[0] == -1{
		fmt.Println("key not found")
		return
	}
	var key_id int
	info := make([][]int, l)
	var previous, child, next int
	nums := 2
	info[0] = []int{0}
	for i, v := range tree{
		if v == -1{
			nums -= 2
			info[i] = append(info[i], -1)
		}else{
			child = next + 2*(i-previous) - 1 - (2*(next-previous) - nums)
			if child <= l-2{
				info[i] = append(info[i], child)
				info[child] = []int{i}
				info[child+1] = []int{i}
			}else if child == l-1{//no right child
				info[i] = append(info[i], child)
                                info[child] = []int{i}
                        }else{// no child
				info[i] = append(info[i], -1)
			}
		}

		if i == next{
			previous = i
			next = i + nums
			nums = 2*nums
		}
		if v == key{
			key_id = i
		}
	}
	leaves, dis := findclosest(tree, info, key_id, math.MaxInt32)
	fmt.Printf("distance-%v\n", dis)
	fmt.Println("leaves-",leaves)
}

func findclosest(tree []int, info [][]int, id, max int) ([]int, int){
	left := info[id][1]
	right := left + 1

	v := tree[id]
	tree[id] = -2
	if left < 0 ||(tree[left] == -1 && (right < len(tree) &&tree[right] == -1)){
		return []int{v}, 0
	}

	tree[id] = -2
	if max == 0{
		return []int{}, math.MaxInt32
	}

	var leaves, leaves2 []int
	var dis int
	parent := info[id][0]

	if tree[left] >= 0{
		leaves, dis = findclosest(tree, info, left, max-1)
		dis += 1
		if dis <= max{
			max = dis
		}else {
			leaves = []int{}
		}
	}

	if right < len(tree) && tree[right] >= 0{
		leaves2, dis = findclosest(tree, info, right, max-1)
		dis += 1
		if dis < max{
			max = dis
			leaves = leaves2
	        }else if dis  == max {
			leaves = append(leaves, leaves2...)
	        }
	}

	if tree[parent] >= 0{
		leaves2, dis = findclosest(tree, info, parent, max-1)
		dis += 1
		if dis  < max{
                        max = dis
                        leaves = leaves2
                }else if dis  == max {
                        leaves = append(leaves, leaves2...)
                }
	}
	return leaves, max
}
