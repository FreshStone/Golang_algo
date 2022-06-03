package main

import "fmt"

func main(){
		m := map[int]bool{}
  closures(m) //maps are passed by reference in golang
  fmt.Println(m)
}

func closures(m map[int]bool){
  type list struct{
    next *list
    val int
  }
  head := list{}
  head.next = &list{val:6}
  head.next.next = &list{val:7}
  head.next.next.next = &list{val:8}
  var tmp *list
  pt := func (h list){ //pt has to be declared early for it to be called inside a
    tmp = head.next
    for ; tmp != nil;{
      fmt.Print(tmp.val, "->")
      tmp = tmp.next
    }
    fmt.Println()
  }
  pt(head)
  var a func(i int)
  a = func (i int){
    if i > 5{
      return
    }
    tmp = &head
    for ;tmp.next != nil; {
      tmp = tmp.next
    }
    tmp.next = &list{val:i}
    pt(head)
    m[i] = true
    a(i+1) // a had to be declared earlier for it to be called inside a
  }
  a(0)
}

