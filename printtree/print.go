package printtree

import "fmt"

type TreeNode struct{
	Val int
	Left, Right *TreeNode
}

func Print(root *TreeNode){
  if root == nil{
    fmt.Println()
    return
  }
   s := [][]*TreeNode{[]*TreeNode{root}}
   var max_not_reached bool
   var depth, i, j, k, offset, gap int
   for{
     s = append(s, []*TreeNode{})
     max_not_reached = false
     for i = 0; i < len(s[depth]); i++{
       if s[depth][i] == nil{
         s[depth+1] = append(s[depth+1], nil)
         s[depth+1] = append(s[depth+1], nil)
       }else{
         if s[depth][i].Left != nil || s[depth][i].Right != nil{
           max_not_reached = true
         }
         s[depth+1] = append(s[depth+1], s[depth][i].Left)
         s[depth+1] = append(s[depth+1], s[depth][i].Right)
       }
     }

     if !max_not_reached{
       s = s[:len(s)-1]
       break
     }
     depth += 1
   }
   offset = 1<<uint(depth) - 1
   depth += 1
   for i, _ = range s{
     for k = offset; k > 0; k--{
       fmt.Print(" ")
     }
     for j = 0; j < len(s[i])-1; j++{
       if s[i][j] == nil{
         fmt.Print(" ")
       }else{
         fmt.Print(s[i][j].Val)
       }
       for k = gap; k > 0; k--{
         fmt.Print(" ")
       }
     }
     gap = offset
     offset /= 2
     if s[i][len(s[i])-1] != nil{
       fmt.Print(s[i][len(s[i])-1].Val)
     }
     fmt.Println()
   }
   return
}
