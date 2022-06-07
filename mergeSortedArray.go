package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 6, 0} //{3,0,0,0}, {0,0,0,0}
	nums2 := []int{5}                //{1,2,4}, {1,2,3,4}
	m := 5
	n := 1
	merge(nums1, m, nums2, n)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	var i int
	for i = m - 1; i > -1; i-- {
		nums1[i+n] = nums1[i]
	}
	m = n
	for i, n = 0, 0; i < len(nums1); i++ {
		if m == len(nums1) || n == len(nums2) {
			break
		}
		if nums1[m] < nums2[n] {
			nums1[i] = nums1[m]
			m += 1
		} else {
			nums1[i] = nums2[n]
			n += 1
		}
	}
	if n == len(nums2) {
		fmt.Println(nums1)
		return
	}
	for ; i < len(nums1); i++ {
		nums1[i] = nums2[n]
		n++
	}
	fmt.Println(nums1)
	return
}
