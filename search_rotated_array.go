package main

import "fmt"

func main(){
	nums := []int{3,4,5,6,1,2}
	target := 2
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	pivot := find_index(nums)
	low := 0; high := len(nums)-1
	for low <= high{
		if target > nums[pivot]{
			if target > nums[high]{
				high = pivot -1
			}else{
				low = pivot+1
			}
		}else if target < nums[pivot]{
			high = pivot -1
		}else{
			return pivot
		}

		if low == high{
			pivot = low
		}else{
			pivot = (high-low)/2
		}
	}
	return -1
}

func find_index(nums []int)int{
	low := 0
	high := len(nums)-1
	var pivot_ind int
	for low < high{
		pivot_ind = (high+low)/2
		if nums[low] > nums[pivot_ind] {
			high = pivot_ind
		}else if nums[pivot_ind + 1] > nums[high]{
			low = pivot_ind + 1
		}else if nums[pivot_ind] > nums[pivot_ind + 1]{
			return pivot_ind + 1
		}else{
			break
		}
	}
	if low == high{
		return low
	}
	return pivot_ind
}
