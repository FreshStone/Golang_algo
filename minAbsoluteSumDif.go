package main

import (
	"fmt"
	"sort"
)

func main() {
	tests_a := [][]int{
		{1, 7, 5},
		{57, 42, 21, 28, 30, 25, 22, 12, 55, 3, 47, 18, 43, 29, 20, 44, 59, 9, 43, 7, 8, 5, 42, 53, 99, 34, 37, 88, 87, 62, 38, 68, 31, 3, 11, 61, 93, 34, 63, 27, 20, 48, 38, 5, 71, 100, 88, 54, 52, 15, 98, 59, 74, 26, 81, 38, 11, 44, 25, 69, 79, 81, 51, 85, 59, 84, 83, 99, 31, 47, 31, 23, 83, 70, 82, 79, 86, 31, 50, 17, 11, 100, 55, 15, 98, 11, 90, 16, 46, 89, 34, 33, 57, 53, 82, 34, 25, 70, 5, 1},
	}
	tests_b := [][]int{
		{2, 3, 5},
		{76, 3, 5, 29, 18, 53, 55, 79, 30, 33, 87, 3, 56, 93, 40, 80, 9, 91, 71, 38, 35, 78, 32, 58, 77, 41, 63, 5, 21, 67, 21, 84, 52, 80, 65, 38, 62, 99, 80, 13, 59, 94, 21, 61, 43, 82, 29, 97, 31, 24, 95, 52, 90, 92, 37, 26, 65, 89, 90, 32, 27, 3, 42, 47, 93, 25, 14, 5, 39, 85, 89, 7, 74, 38, 12, 46, 40, 25, 51, 2, 19, 8, 21, 62, 58, 29, 32, 77, 62, 9, 74, 98, 10, 55, 25, 62, 48, 48, 24, 21},
	}
	ans := []int{3, 3441}
	for i := 0; i < len(tests_a); i++ {
		if ans[i] != minAbsoluteSumDiff(tests_a[i], tests_b[i]) {
			fmt.Println("error expected", ans[i], "output", minAbsoluteSumDiff(tests_a[i], tests_b[i]))
		} else {
			fmt.Println("correct")
		}
	}
}

/*leetcode - 1818*/

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	if len(nums1) == 1 {
		return absDiff(nums1[0], nums2[0])
	}
	var i, sum, ans, s, e int
	search := make([]int, len(nums1))
	copy(search, nums1)
	sort.Ints(search)
	for ; i < len(nums1); i++ {
		sum += absDiff(nums1[i], nums2[i])
	}
	ans = sum
	for i = 0; i < len(nums1); i++ {
		if nums1[i] == nums2[i] {
			continue
		} else if nums2[i] <= search[0] {
			ans = min(ans, sum+(search[0]-nums1[i]))
		} else if nums2[i] >= search[len(nums1)-1] {
			ans = min(ans, sum+(nums1[i]-search[len(nums1)-1]))
		} else {
			s = 0
			e = len(nums1) - 1
			for s <= e {
				if search[(s+e)/2] > nums2[i] {
					e = (s+e)/2 - 1
				} else if search[(s+e)/2] < nums2[i] {
					s = (s+e)/2 + 1
				} else {
					e = (s + e) / 2
					break
				}
			}
			ans = min(ans, sum-absDiff(nums1[i], nums2[i])+min(nums2[i]-search[e], search[e+1]-nums2[i]))
		}
	}
	return ans % 1000000007
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
