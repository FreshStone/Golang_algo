package main

import "fmt"

type node struct{
  val, height int
  left, right *node
}

func main(){
  var root *node
  /*
  root := insert(nil, 5)
  root = insert(root, 2) //no rotation
  root = insert(root, 7) //no rotation
  root = insert(root, 1) //no rotation
  root = insert(root, 4) //no rotation
  root = insert(root, 6) //no rotation
  root = insert(root, 9) //no rotation
  root = insert(root, 3) //no rotation
  root = insert(root, 16) //no rotation
  root = insert(root, 15) //Right-Left rotation*/
  root = insert(root, 9)
  root = insert(root, 5)
  root = insert(root, 10)
  root = insert(root, 0)
  root = insert(root, 6)
  root = insert(root, 11)
  root = insert(root, -1)
  root = insert(root, 1)
  root = insert(root, 2)//L-R rotation
  root = delete(root, 11) //R rotation
  printpreorder(root)
  fmt.Println()
  printlevelorder(root)
}

func insert(root *node, v int)*node{
  if root == nil{
    return &node{val: v}
  }

  if v > root.val{
    root.right = insert(root.right, v)
  }else if v < root.val{
    root.left = insert(root.left, v)
  }else{
    return root
  }

  hl := height(root.left)
  hr := height(root.right)
  if hl > hr{
    if hl-hr > 1{
      return rebalance(root, true)
    }
    root.height = 1+hl
  }else if hr > hl{
    if hr-hl > 1{
      return rebalance(root, false)
    }
    root.height = 1+hr
  }
  return root
}

func delete(root *node, v int)*node{
  if root == nil{
    return nil
  }

  if v < root.val{
    root.left = delete(root.left, v)
  }else if v > root.val{
    root.right = delete(root.right, v)
  }else if root.right == nil{
    return root.left
  }else{
    root.val = findinorder(root.right)
    root.right = delete(root.right, root.val)
  }

  hl := height(root.left)
  hr := height(root.right)
  if hl > hr{
    if hl-hr > 1{
      return rebalance(root, true)
    }
    root.height = 1+hl
  }else if hr > hl{
    if hr-hl > 1{
      return rebalance(root, false)
    }
    root.height = 1+hr
  }else{
    root.height = 1+hr
  }
  return root
}

func findinorder(root *node)int{
  if root.left == nil{
    return root.val
  }
  return findinorder(root.left)
}

func rebalance(root *node, left bool)*node{ //left subtree height is more
  if left{
    // inserted_val > root.left.val
    if root.left.left == nil || (root.left.right != nil && root.left.left.height < root.left.right.height){
      root.left = leftrotate(root.left)
    }
    root = rightrotate(root)
  }else{
    // inserted_val < root.right.val
    if root.right.right == nil || (root.right.left != nil && root.right.left.height > root.right.right.height){
      root.right = rightrotate(root.right)
    }
    root = leftrotate(root)
  }

  return root
}

func leftrotate(root *node)*node{
  tmp := root.right
  root.right = tmp.left
  tmp.left = root
  root.height = height(root)
  tmp.height = height(tmp)
  return tmp
}

func rightrotate(root *node)*node{
  tmp := root.left
  root.left = tmp.right
  tmp.right = root
  root.height = height(root)
  tmp.height = height(tmp)
  return tmp
}

func height(root *node)int{
  if root == nil{
    return -1
  }
  hl := -1
  hr := -1
  if root.left != nil{
    hl = root.left.height
  }

  if root.right != nil{
    hr = root.right.height
  }
  if hl > hr{
    return 1+hl
  }

  return 1+hr
}
func printpreorder(root * node){
  if root == nil{
    return
  }
  printpreorder(root.left)
  fmt.Print(root.val, " ")
  printpreorder(root.right)
  return
}

func printlevelorder(root *node){
  if root == nil{
    return
  }
  var i, l int
  q := []*node{root}
  for ;len(q) > 0;{
    l = len(q)
    for i = 0; i < l; i++{
      fmt.Print(q[i], " ")
      if q[i].left != nil{
        q = append(q, q[i].left)
      }
      if q[i].right != nil{
        q = append(q, q[i].right)
      }
    }
    fmt.Println()
    q = q[l:]
  }
  return
}
