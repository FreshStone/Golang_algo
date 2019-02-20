package main

import "fmt"

type node struct{
	m map[string]*node
	previousend bool
}

func main(){
	a := []string{"dog", "dream", "dreamer", "dr", "dreat"}
	root := &node{m: make(map[string]*node)}
	for _, v := range a{
		insert(root, v)
	}
	fmt.Println(find(root, "drea", 0))
}

func insert(r *node, a string){
	if len(a) == 0{
                r.previousend = true
                return
        }
	_, ok := r.m[string(a[0])]
	if !ok{
		r.m[string(a[0])] = &node{m: make(map[string]*node)}
	}
	insert(r.m[string(a[0])], a[1:])
	return
}

func find(r *node, a string, i int)[]string{
	var rtr []string
	if i < len(a){
		add, ok := r.m[string(a[i])]
		if !ok{
			return rtr
		}else{
			return find(add, a, i+1)
		}
	}

	if r.previousend{
		rtr = append(rtr, a[:i])
	}

	for k, v := range r.m{
		rtr = append(rtr, find(v, a+k, len(a)+1 )...)
	}

	return rtr
}
