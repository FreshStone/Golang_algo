package main

import (
  "fmt"
  "sort"
)

func main(){
  rectangles := [][]int{
    {1,1,3,3},
    {3,1,4,2},
    {3,2,4,4},
    {1,3,2,4},
    {2,3,3,4},
  }
  fmt.Println(isRectangleCover(rectangles))
}
 //neither a point nor a line is considered a rectangle
func isRectangleCover(rectangles [][]int) bool {
    if len(rectangles) == 1{
      return true
    }
    var y [][]int
    var i, prev_h, curr_h, prev_x, l, r int
    x := make([][]int, 2*len(rectangles))
    for ;i < len(rectangles); i++{
      x[2*i] = []int{rectangles[i][0], rectangles[i][1], rectangles[i][3], 0} //start of rectangle
      x[2*i+1] = []int{rectangles[i][2], rectangles[i][1], rectangles[i][3], 1} //end
    }

    sort.Slice(x, func(i int, j int) bool {
      return (x[i][0] < x[j][0]) || (x[i][0] == x[j][0] && (x[i][3] > x[j][3] || (x[i][3] == x[j][3] && x[i][1] < x[j][1])))
    })
    prev_x = x[0][0]
    y = append(y, []int{x[0][1], x[0][2]})
    for i = 1; i < len(x) && x[i][0] == x[i-1][0]; i++{
      if x[i][1] == y[len(y)-1][1]{
        y = append(y, []int{x[i][1], x[i][2]})
      }else{
        return false
      }
    }
    prev_h = y[len(y)-1][1]-y[0][0]
    curr_h = prev_h
    for ;i < len(x); i++{
      if x[i][0] == prev_x || curr_h == prev_h{
        l = 0
        r = len(y)-1
        if x[i][3] == 0{ // add edge
          for ;l <= r;{
            if x[i][1] < y[(l+r)/2][0]{
              r = (l+r)/2 -1
            }else if x[i][1] > y[(l+r)/2][0]{
              l = (l+r)/2 + 1
            }else{
              return false
            }
          }

          if l < len(y) && x[i][2] > y[l][0]{
            return false
          }
          curr_h += x[i][2]-x[i][1]
          y = append(y[:l], append([][]int{{x[i][1],x[i][2]}}, y[l:]...)...)
        }else{ //remove edge
          for ;l <= r;{
            if x[i][1] < y[(l+r)/2][0]{
              r = (l+r)/2 -1
            }else if x[i][1] > y[(l+r)/2][0]{
              l = (l+r)/2 + 1
            }else{
              l = (l+r)/2
              break
            }
          }
          curr_h -= x[i][2]-x[i][1]
          y = append(y[:l], y[l+1:]...)
        }
        prev_x = x[i][0]
      }else{
        return false
      }
    }
    return true
}
