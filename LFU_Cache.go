package main

import "fmt"

type node struct{
	k string
	val, frq int
	left *node
	right *node
}

type c struct{
	la, ra *node
	li, ri int
}

var adr = make(map[string]*node)
var cnt []c
var head = &node{}
var tail = &node{val: -1}
var l int  //current len of cache
var n = 3 //max len of cache; n > 0

func main(){
	//before setting a value check with get() whether it already exist or not
	set("a", 1)
	set("b", 2)
	set("c", 3)
	get("a")
	get("a")
	get("c")
	get("b")
	set("d", 4)
	get("b")
	get("d")
	get("d")
	set("e", 5)
	get("e")
	get("e")
	printlist(head.right)
}

func set(key string, v int){
	_, ok1 := adr[key]
	var a, tmp *node
	if ok1{   //assuming no already present val is inserted via set()
		return
	}else{
		a = &node{k: key, val: v, frq: 1}
		if l == 0{
			head.right = a
			tail.left = a
			a.right = tail
			a.left = head
			cnt = append(cnt, c{la: a, ra: a, li: -1, ri: -1})
			l += 1
			adr[key] = a
			return
		}

		if l == n{
			// delete last from map, remove last from linkedlist, update l, update cnt
			delete(adr, tail.left.k)
			tmp = tail.left
			tmp.left.right = tail
			tail.left = tmp.left
			l -= 1
			if cnt[tmp.frq-1].la == cnt[tmp.frq-1].ra{
				cnt[tmp.frq-1].la = nil
				cnt[tmp.frq-1].ra = nil
			}else{
                                cnt[tmp.frq-1].ra = tmp.left
			}
		}

		if cnt[0].la == nil{
			cnt[0].la = a
			cnt[0].ra = a
			tail.left.right = a
			a.left = tail.left
			tail.left = a
			a.right = tail
		}else{
			tmp = cnt[0].la
			tmp.left.right = a
			a.left = tmp.left
			tmp.left = a
			a.right = tmp
			cnt[0].la = a
		}
		l += 1
		adr[key] = a
	}
}

func get(key string)int{
	a, ok1 := adr[key]
	var tmp_l, tmp_r *node
	if ok1{
		if cnt[a.frq-1].ri != -1 && cnt[cnt[a.frq-1].ri].la != nil{
			tmp_l = cnt[cnt[a.frq-1].ri].la
			tmp_r = cnt[cnt[a.frq-1].ri].ra

			if cnt[a.frq-1].ri > a.frq{
				cnt[a.frq].la = a
	                        cnt[a.frq].ra = a
	                        cnt[a.frq].ri = cnt[a.frq-1].ri
	                        cnt[cnt[a.frq-1].ri].li = a.frq

	                        if cnt[a.frq-1].la != cnt[a.frq-1].ra{
	                                cnt[a.frq].li = a.frq-1
	                                cnt[a.frq-1].ri = a.frq
	                                if cnt[a.frq-1].la == a{
						cnt[a.frq-1].la = a.right
	                                }

	                                if cnt[a.frq-1].ra == a{
	                                        cnt[a.frq-1].ra = a.left
	                                }
	                        }else{
	                                cnt[a.frq-1].la = nil
	                                cnt[a.frq-1].ra = nil
	                                if a.frq != 1{
	                                        cnt[a.frq].li = cnt[a.frq-1].li
	                                        cnt[cnt[a.frq-1].li].ri = a.frq
	                                }else{
	                                        cnt[1].li = 0
	                                        cnt[0].ri = 1
	                                }
		                }

				if tmp_r.right != a{
					a.left.right = a.right
		                        a.right.left = a.left
	                                tmp_r.right.left = a
	                                a.right = tmp_r.right
	                                tmp_r.right = a
	                                a.left = tmp_r
				}

			}else{
				cnt[a.frq].la = a
				if cnt[a.frq-1].la != cnt[a.frq-1].ra{
	                                if cnt[a.frq-1].la == a{
	                                        cnt[a.frq-1].la = a.right
	                                }

	                                if cnt[a.frq-1].ra == a{
	                                        cnt[a.frq-1].ra = a.left
	                                }
	                        }else{
	                                cnt[a.frq-1].la = nil
	                                cnt[a.frq-1].ra = nil
	                                if a.frq != 1{
	                                        cnt[a.frq].li = cnt[a.frq-1].li
	                                        cnt[cnt[a.frq-1].li].ri = a.frq
	                                }
	                        }

				a.left.right = a.right
	                        a.right.left = a.left
				tmp_l.left.right = a
				a.left = tmp_l.left
				tmp_l.left = a
				a.right = tmp_l
			}
			a.frq += 1
		}else{
			if len(cnt) < a.frq+1{
				cnt = append(cnt, c{la: a, ra: a, li: a.frq-1, ri: -1})
			}else{
				cnt[a.frq].la = a
				cnt[a.frq].ra = a
				cnt[a.frq].li = a.frq-1
				cnt[a.frq].ri = -1
			}

			tmp_l = cnt[a.frq-1].la
			if cnt[a.frq-1].la != cnt[a.frq-1].ra{
                                if cnt[a.frq-1].la == a{
					cnt[a.frq-1].la = a.right
                                }

                                if cnt[a.frq-1].ra == a{
					cnt[a.frq-1].ra = a.left
                                }
				cnt[a.frq-1].ri = a.frq
                         }else{
                                cnt[a.frq-1].la = nil
                                cnt[a.frq-1].ra = nil
                                if a.frq != 1{
                                         cnt[a.frq].li = cnt[a.frq-1].li
                                         cnt[cnt[a.frq-1].li].ri = a.frq
                                }else{
					cnt[0].ri = 1
				}
                         }

			 if head.right != a{
				 a.left.right = a.right
	                         a.right.left = a.left
				 tmp_l.left.right = a
				 a.left = tmp_l.left
				 tmp_l.left = a
				 a.right = tmp_l
			 }
			 a.frq += 1
		}
		return a.val
	 }
	 return -1
}

func printlist(n *node){
	if n.val == -1{
		fmt.Println()
		return
	}
	fmt.Print(n.val, "->")
	printlist(n.right)
}
