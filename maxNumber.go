package main

import "fmt"

func main(){
  nums1 := [][]int{{3,4,6,5},    {6,7},  {3,9}, {3,9}, {2,4,6,6,1}}
  nums2 := [][]int{{9,1,2,5,8,3},{6,0,4},{8,9}, {8,9}, {3,6,6}}
  k := []int{5,5,3,4,5}
  //[9, 8, 6, 5, 3], [6,7,6,0,4], [9,8,9], [8,9,3,9], [6,6,6,6,1]
  for i, _ := range nums1{
    fmt.Println(MaxNumber(nums1[i], nums2[i], k[i]))
  }
}

func MaxNumber(nums1, nums2 []int, k int)[]int{
  if k == 0{
    return []int{}
  }
  var cm, nm, l, i, m int
  var res []int
  var q [][]int //q[i] = []int{max_val, itr of nums1, itr of nums2 }
  v := make([][][]bool, len(nums1)+1)
  a2 := make([][]int, len(nums2)+1)
  a1 := make([][]int, len(nums1)+1)
  v[len(nums1)] = make([][]bool, len(nums2)+1)
  a2[len(nums2)] = []int{-1,-1,-1,-1,-1,-1,-1,-1,-1,-1}
  a1[len(nums1)] = []int{-1,-1,-1,-1,-1,-1,-1,-1,-1,-1}
  for i = len(nums1)-1 ;i >= 0 ; i--{
    v[i] = make([][]bool, len(nums2)+1)
    a1[i] = make([]int, 10)
    copy(a1[i], a1[i+1])
    a1[i][nums1[i]] = i
  }

  for i = len(nums2)-1 ;i >= 0; i--{
    a2[i] = make([]int, 10)
    copy(a2[i], a2[i+1])
    a2[i][nums2[i]] = i
  }

  q = append(q, []int{0, 0, 0})
  v[0][0] = make([]bool, k+1)
  f := func(a bool, x, y, e int){
    if a{
      // finds max element in the range [x,e) in constant (O(10)) time
      for m = 9; m >= 0; m--{
        if a1[x][m] != a1[e][m]{
          break
        }
      }
      x = a1[x][m]+1
    }else{
      for m = 9; m >= 0; m--{
        if a2[y][m] != a2[e][m]{
          break
        }
      }
      y = a2[y][m]+1
    }

    if m >= nm{
      nm = m
      if len(v[x][y]) == 0{
        v[x][y] = make([]bool, k+1)
      }
      q = append(q, []int{m, x, y})
    }
  }
  for ;k > 0; k-- {
    res = append(res, cm)
    l = len(q)
    for i = 0; i < l; i++{
      if !v[q[i][1]][q[i][2]][k] && q[i][0] == cm{
        if q[i][1] == len(nums1){
          f(false, q[i][1], q[i][2], len(nums2)-k+1)
        }else if q[i][2] == len(nums2){
          f(true, q[i][1], q[i][2], len(nums1)-k+1)
        }else{
          if len(nums2) > q[i][2] + k - 2 {//len(nums2)-q[i][2] >= k-1
            f(true, q[i][1], q[i][2], len(nums1))
          }else{
            f(true, q[i][1], q[i][2], len(nums1)-k+len(nums2)-q[i][2]+1)
          }

          if len(nums1) > q[i][1] + k - 2 {//len(nums1)-q[i][1] >= k-1
            f(false, q[i][1], q[i][2], len(nums2))
          }else{
            f(false, q[i][1], q[i][2], len(nums1)-k+len(nums2)-q[i][1]+1)
          }
        }
        v[q[i][1]][q[i][2]][k] = true
      }
    }
    q = q[l:]
    cm = nm
    nm = -1
  //  fmt.Println(res, q)
  }
  return append(res, cm)[1:]
}
