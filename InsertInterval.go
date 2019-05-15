package main

import "fmt"

func main(){
	a := [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}
	b := []int{4,8}
	fmt.Println(insert(a, b))
}

func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0{
        return append(intervals, newInterval)
    }
    var tmp [][]int
    l := bs(intervals, newInterval[0]) -1
    r := bs(intervals, newInterval[1])
    if l == -1 || newInterval[0] > intervals[l][1]{
        tmp = intervals[r:]
        if r == 0{
            tmp = append([][]int{newInterval}, tmp...)
        }else{
            tmp = append([][]int{{newInterval[0], max(intervals[r-1][1], newInterval[1])}}, tmp...)
        }
        intervals = append(intervals[:l+1], tmp...)
        return intervals
    }
    tmp = intervals[r:]
    intervals = append(intervals[:l], []int{intervals[l][0], max(intervals[r-1][1], newInterval[1])})
    intervals = append(intervals, tmp...)
    return intervals
}

func bs(intervals [][]int, v int)int{
    l := 0
    r := len(intervals)-1
    for l <= r{
        if v < intervals[(l+r)/2][0]{
            r = (l+r)/2 - 1
        }else if v > intervals[(l+r)/2][0]{
            l = (l+r)/2 + 1
        }else{
            return (l+r)/2 + 1
        }
    }
    return l
}

func max(i, j int)int{
    if i > j{
        return i
    }
    return j
}
