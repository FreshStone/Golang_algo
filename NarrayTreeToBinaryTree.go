package main

import (
  "fmt"
)

type tree struct{
  val int
  child []*tree
}
/*
                    1
-----------------------------------------------
|         |              |      |             |
2         3              4      5             6
|         |              |                    |
7      8 9 10 11       12 13           14 15 16 17 18
         |
       19 20 21
*/

func main(){
  var{
    root := &tree{val:1}
    root.child = append(root.child, &tree{val: 2})
    root.child = append(root.child, &tree{val: 3})
    root.child = append(root.child, &tree{val: 4})
    root.child = append(root.child, &tree{val: 5})
    root.child = append(root.child, &tree{val: 6})
    root.child[0].child = append(root.child[0].child, &tree{val: 7})
    root.child[1].child = append(root.child[1].child, &tree{val: 8})
    root.child[1].child = append(root.child[1].child, &tree{val: 9})
    root.child[1].child = append(root.child[1].child, &tree{val: 10})
    root.child[1].child = append(root.child[1].child, &tree{val: 11})
    root.child[2].child = append(root.child[2].child, &tree{val: 12})
    root.child[2].child = append(root.child[2].child, &tree{val: 13})
    root.child[4].child = append(root.child[4].child, &tree{val: 14})
    root.child[4].child = append(root.child[4].child, &tree{val: 15})
    root.child[4].child = append(root.child[4].child, &tree{val: 16})
    root.child[4].child = append(root.child[4].child, &tree{val: 17})
    root.child[4].child = append(root.child[4].child, &tree{val: 18})
    root.child[1].child[1].child = append(root.child[1].child[1].child, &tree{val: 19})
    root.child[1].child[1].child = append(root.child[1].child[1].child, &tree{val: 20})
    root.child[1].child[1].child = append(root.child[1].child[1].child, &tree{val: 21})
  }
  printlevel(root) //level_order_Narray_Tree
  root = encode(root)
  printlevel(root) //level_order_bt
  printinorder(root) //inorder_bt
  fmt.Println()
  root = decode(root)
  printlevel(root) //level_order_Narray_Tree
}

func encode(root *tree)*tree{
  if root == nil || len(root.child) == 0{
    return root
  }
  root.child = []*tree{nil, rec_encode(root.child)}
  return root
}
/* encoding
     root
   /      \
siblings    child
*/
func rec_encode(siblings []*tree)*tree{
  if len(siblings) == 0{
    return nil
  }
  x := rec_encode(siblings[1:]) //siblings
  y := rec_encode(siblings[0].child) //child
  siblings[0].child = siblings[0].child[:0]
  siblings[0].child = append(siblings[0].child, x)
  siblings[0].child = append(siblings[0].child, y)
  return siblings[0]
}

func decode(root * tree)*tree{
  if root == nil || len(root.child) == 0{
    return root
  }
  root.child = rec_decode(root.child[1], root.child[:0])
  return root
}

func rec_decode(root *tree, siblings []*tree)[]*tree{
  if root == nil{
    return siblings
  }
  x := root.child[0]
  root.child = rec_decode(root.child[1], root.child[:0])
  siblings = append(siblings, root)
  return rec_decode(x, siblings)
}

func printinorder(root *tree){ //call only for bt
  if root == nil{
    return
  }
  printinorder(root.child[0])
  fmt.Print(root.val, " ")
  printinorder(root.child[1])
  return
}
func printlevel(root *tree){
  arr := []*tree{root}
  var i, l int
  for ;len(arr) > 0;{
    l = len(arr)
    for i = 0; i < l; i++{
      if arr[i] != nil{
        fmt.Print(arr[i].val, " ")
        arr = append(arr, arr[i].child...)
      }
    }
    arr = arr[l:]
    fmt.Println()
  }
  return
}
