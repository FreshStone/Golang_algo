package main

func addtwonumbers(a1 *node, a2 *node) *node {
    var a int
    c := 0
    var r node
    l := &r
    for ;(a1.next != nil) || (a2.next != nil) || c == 1;{
	switch {
	case (*a1).next == nil && (*a1).next == nil :
	a1 = &node{0, nil}
	a2 = &node{0, nil}
	case (*a1).next == nil :
	a1 = &node{0, nil}
	a2 = (*a2).next
	case (*a2).next == nil :
        a2 = &node{0, nil}
	a1 = (*a1).next
	default :
	a1 = (*a1).next
	a2 = (*a2).next
	}
	/* "or"
	a1 = (*a1).next
        a2 = (*a2).next
	if a1 == nil {
		 a1 = &node{0, nil}
	}
	if a2 == nil {
		 a2 = &node{0, nil}
	 }
	 */

        a = (*a1).val + (*a2).val + c
	if a > 9 {
            a -=10
            c = 1
        }else{
		c = 0
	}
        var n node
        (*l).next = &n
        n.val = a
        l = &n
   }
   return &r
}

