package main

func reverse(root *node){
	if (*root).next == nil {
		r.next = root
		return
	}
	n := (*root).next
	reverse(n)
	if root != &r{
		(*n).next = root
		(*root).next = nil
	}
}

