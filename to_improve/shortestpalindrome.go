package main

import (
  "fmt"
  "strings"
)

func main(){
  s := []string{
    "aaaa",
    "abcde",
    "AABAACAABAA",
    "AAACAAAAAC",
    "AAABAAA",
    "ababbabacbae",
    "aacecaaa",
    "abcd",
    "babbbabbaba",
  }
  for _, v := range s{
    fmt.Println(shortestPalindrome(v))
    //shortestPalindrome(v)
  }
}

func shortestPalindrome(s string)string {
    if len(s) < 2{
      return s
    }
    var i, j int
    var b strings.Builder
    jump := make([]int, len(s))
    p := make([]int, len(s))
    for i = 1; i < len(s)/2;{
  //    fmt.Println(i, j)
      if s[i] == s[j]{
        j += 1
        jump[i] = j
        i += 1
      }else{
        if j == 0{
          i += 1
        }else{
          j = jump[j-1]
        }
      }
    }
    fmt.Println(jump)
    j = 0
    for i = len(s)-1; i >= 0 && i > j;{
    //  fmt.Println(i, j)
      if j > 0{
        if s[i] == s[j]{
          p[i] = j
          j += 1
          i -= 1
        }else{
          j = jump[j-1]
        }
      }else{
        if s[i] == s[0]{
          j = 1
          p[i] = 0
        }
        i -= 1
      }
    }
    //fmt.Println("i-", i, "j-", j)
    if i == j{
      j *= 2
    }else{
      j = 2*j -1
    }

    for i = len(s)-1; i > j; i--{
      b.WriteByte(s[i])
    }
    b.WriteString(s)
    return b.String()
}
