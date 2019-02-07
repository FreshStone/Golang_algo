package main

import "fmt"

type ListNode struct {
	Val int
        Next *ListNode
}

var r ListNode

func main(){
	a := []int{1,2,3,4,5}
	for _, v := range a{
		insert(&r, v)
	}
	r := *reverseKGroup(&r, 2)
	printlist(r.Next)
}

func insert(root *ListNode, a int){
        var n ListNode
        if (*root).Next == nil {
                (*root).Next = &n
                n.Val = a
                return
        }
        insert((*root).Next, a)
}

func printlist(root *ListNode){
	if root.Next == nil{
		fmt.Printf("%v\n", root.Val)
		return
	}
	fmt.Printf("%v->", root.Val)
	printlist(root.Next)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1{
		return head
	}
	v := k
	start := head
	first := head.Next
	curr := first
	var tmp, previous *ListNode
	for curr != nil{
		if v == 1{
			start.Next = curr
			first.Next = nil
			start = first
			first = curr.Next
			curr.Next = previous
			curr = first
			previous = nil
			v = k
		}else{
			if curr.Next != nil{
				tmp = curr.Next
				curr.Next = previous
				previous = curr
				curr = tmp
				v -= 1
			}else{
				start.Next = first
				previous, curr = curr, previous
				for curr != nil{
					tmp = curr.Next
					curr.Next = previous
					previous = curr
					curr = tmp
				}
			}
		}
	}
	return head
}
