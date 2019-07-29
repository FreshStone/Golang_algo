package main

import "fmt"

func main(){
  s := "catbatatecatatebat"
  words := []string{"cat","ate","bat"}
  fmt.Println(startingindex(s, words))
}

func startingindex(s string, words []string)[]int{
  if len(s) == 0 || len(words) == 0 || len(s) < len(words[0]){
    return []int{}
  }

  var i, v, j int
  var ok bool
  l := len(words[0])
  var res []int
  store := make([][]int, l)
  m := make(map[string]int)

  for i = 0; i < len(words); i++{
    m[words[i]] = i
  }

  for i = 0; i < l && i+l < len(s)+1; i++ {
    store[i] = []int{}
    v, ok = m[s[i:i+l]]
    if ok{
      store[i] = append(store[i], v)
    }
  }

  for ;i+l < len(s)+1; i++{
    v, ok = m[s[i:i+l]]
    if ok{
      for j = 0; j < len(store[0]); j++{ //check for prior occurrence of the matched word
        if store[0][j] == v{             //  if words[] contain duplicates words than m = map[string][]int
          store[0] = store[0][j+1:]      //  if store[0][j] = v[0]{   //initialize cnt = 0
          break                          //          if cnt == 0{ firstOccurrence = j}; cnt += 1
        }                                //          if cnt == len(v){ store[0] = store[0][firstOccurrence+1:] }
      }                                  //   }
      store[0] = append(store[0], v)     //   store[0] = append(store[0], v[0])
      if len(store[0]) == len(words){
        res = append(res, i-(len(words)-1)*l)
        store[0] = store[0][1:]
      }
    }else{
      store[0] = []int{}
    }
    store = append(store, store[0])
    store = store[1:]
  }
  return res
}
