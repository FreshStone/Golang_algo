package main

import "fmt"

type AllOne struct{
  m map[string]*dll
  g map[int][]*dll
  min, max int
}

type dll struct{
  k string
  v int
  prv, nxt *dll
}
/*
  1. double linked list
  2. (m)map1 contains address of every node
  2. (g)map2 contains addresses(start & end) of each group of nodes(strings) with same values
*/

func Constructor() AllOne {
    return AllOne{m: make(map[string]*dll), g: make(map[int][]*dll)}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string)  {
    adr := this.m[key]
    if adr == nil{
      adr = &dll{k: key, v: 1}
      this.m[key] = adr
      if this.min == 0{
        this.min = 1
        this.max = 1
        this.g[1] = []*dll{adr, adr}
        return
      }
      adr.v = 0
      adr.nxt = this.g[this.min][0]
      this.g[this.min][0].prv = adr
      this.min = 1
    }else if this.g[adr.v][0] == this.g[adr.v][1]{
      if this.min == adr.v{
        this.min = adr.v + 1
      }
      delete(this.g, adr.v)
    }else if this.g[adr.v][1] == adr{
      this.g[adr.v][1] = adr.prv
    }else {
      adr.nxt.prv = adr.prv
      if this.g[adr.v][0] == adr{
        this.g[adr.v][0] = adr.nxt
      }
      if adr.prv != nil{
        adr.prv.nxt = adr.nxt
      }

      adr.nxt = this.g[adr.v][1].nxt
      adr.prv = this.g[adr.v][1]
      this.g[adr.v][1].nxt = adr
      if adr.nxt != nil{
        adr.nxt.prv = adr
      }
    }

    adr.v += 1
    if adr.v == this.max + 1{
      this.max += 1
      this.g[adr.v] = []*dll{adr, adr}
    }else if len(this.g[adr.v]) != 0{
      this.g[adr.v][0] = adr
    }else{
      this.g[adr.v] = []*dll{adr, adr}
    }
    return
}


/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string)  {
  if this.m[key] == nil{
    return
  }
  adr := this.m[key]
  if adr.v == 1{
    delete(this.m, key)
    if this.g[1][0] == this.g[1][1]{
      delete(this.g, 1)
      if this.max == 1{
        this.max = 0
        this.min = 0
      }else{
        this.min = adr.nxt.v
        adr.nxt.prv = nil
      }
    }else{
      if this.g[1][0] == adr{
        this.g[1][0] = adr.nxt
      }
      if this.g[1][1] == adr{
        this.g[1][1] = adr.prv
      }
      if adr.nxt != nil{
        adr.nxt.prv = adr.prv
      }
      if adr.prv != nil{
        adr.prv.nxt = adr.nxt
      }
    }
    return
  }


  if this.g[adr.v][0] == this.g[adr.v][1]{
    if this.max == adr.v{
      this.max = adr.v - 1
    }
    delete(this.g, adr.v)
  }else if this.g[adr.v][0] == adr{
    this.g[adr.v][0] = adr.nxt
  }else {
    adr.prv.nxt = adr.nxt
    if this.g[adr.v][1] == adr{
      this.g[adr.v][1] = adr.prv
    }
    if adr.nxt != nil{
      adr.nxt.prv = adr.prv
    }

    adr.nxt = this.g[adr.v][0]
    adr.prv = this.g[adr.v][0].prv
    this.g[adr.v][0].prv = adr
    if adr.prv != nil{
      adr.prv.nxt = adr
    }
  }

  adr.v -= 1
  if adr.v == this.min - 1{
    this.min -= 1
    this.g[adr.v] = []*dll{adr, adr}
  }else if len(this.g[adr.v]) != 0{
    this.g[adr.v][1] = adr
  }else{
    this.g[adr.v] = []*dll{adr, adr}
  }
  return
}


/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
    if len(this.m) == 0{
      return ""
    }
    return this.g[this.max][1].k
}


/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
  if len(this.m) == 0{
    return ""
  }
  return this.g[this.min][0].k
}

func main(){
  obj := Constructor()
  obj.Inc("aa")
  obj.Inc("bb")
  obj.Inc("gg")
  obj.Inc("aa")
  obj.Dec("gg")
  //fmt.Println(obj.m, obj.g)
  fmt.Println(obj.GetMinKey())
  obj.Dec("aa")
  obj.Inc("bb")
  //fmt.Println(obj.m, obj.g)
  fmt.Println(obj.GetMaxKey())
  fmt.Println(obj.GetMinKey())
}
