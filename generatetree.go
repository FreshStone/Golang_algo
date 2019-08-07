package main
import "fmt"

type node struct{
  val int
  left, right *node
}
func main(){
  a := []int{1,3,2,6,7,8,9,5,4,11,12,13,15,17,16,21,20,19,18,14,10}
  printinorder(postorderbst(a))
  fmt.Println()
}

func postorderbst(tree []int)*node{
  if len(tree) == 0{
    return nil
  }
  min := -1<<31
  max := 1<<31 -1
  type a struct{
    root *node
    lower, upper int
    writetoleft bool
  }
  var tmp *node
  var arr []a
  arr = append(arr, a{root: &node{val: tree[len(tree)-1]}, lower: min, upper: max})
  for i := len(tree)-2; i > -1;{
    if arr[len(arr)-1].root.val < tree[i]{ //right branch
      arr = append(arr, a{root: &node{val: tree[i]}, lower: arr[len(arr)-1].root.val+1, upper: arr[len(arr)-1].upper})
      i -= 1
    }else{
      if arr[len(arr)-1].lower <= tree[i]{ //left branch
        arr[len(arr)-1].writetoleft = true
        arr = append(arr, a{root: &node{val: tree[i]}, lower: arr[len(arr)-1].lower, upper: arr[len(arr)-1].root.val})
        i -= 1
      }else{ //go to parent
        tmp = arr[len(arr)-1].root
        arr = arr[:len(arr)-1]
        if arr[len(arr)-1].writetoleft{
            arr[len(arr)-1].root.left = tmp
        }else{
            arr[len(arr)-1].root.right = tmp
        }
      }
    }
  }
  for i := len(arr)-1; i > 0; i--{
    if arr[i-1].writetoleft{
        arr[i-1].root.left = arr[i].root
    }else{
        arr[i-1].root.right = arr[i].root
    }
  }

  return arr[0].root
}

func printinorder(root *node){
  if root == nil{
    return
  }
  printinorder(root.left)
  fmt.Print(" ", root.val)
  printinorder(root.right)
}
