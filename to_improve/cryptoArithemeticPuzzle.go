package main

import (
  "fmt"
  //"math"
)

func main(){
  test := [][]string{
    {"DB", "CD", "EF"},
    {"WFAE", "BXYZ", "FPPP"},
    {"XWGF", "CBA", "ADEM"},
    {"BA", "GHFEDC", "GHFCHZ"},
    {"ABCD", "AFGHIJ", "AFHKAB"},
    {"SEND", "MORE", "MONEY"},
  }
  for _, v := range test{
    r := decrypt(v[0], v[1], v[2])
    f(r)
  }
}

func decrypt(a, b, c string)[]byte{
  if len(c) < len(a) || len(c) < len(b){
    return []byte{}
  }
  letters := make([]int, 26)
  digits := make([]byte, 10)
  rec(0, a, b, c, letters, digits)
  return digits
}

func rec(carry int, a, b, c string, letters []int, digits []byte)int{
  if len(a) == 0 && len(b) == 0 && len(c) == 0{
    if carry == 0{
      return 1
    }
    return -1
  }
  var i int
  if len(a) > 0 && letters[a[len(a)-1]-65] == 0{
    for i = 0; i < 10; i++{
      if digits[i] == 0{
        digits[i] = a[len(a)-1]
        letters[a[len(a)-1]-65] = i+1
        if rec(carry, a, b, c, letters, digits) == 1{
          return 1
        }
        digits[i] = 0
        letters[a[len(a)-1]-65] = 0
      }
    }
    return -1
  }

  if len(b) > 0 && letters[b[len(b)-1]-65] == 0{
    for i = 0; i < 10; i++{
      if digits[i] == 0{
        digits[i] = b[len(b)-1]
        letters[b[len(b)-1]-65] = i+1
        if rec(carry, a, b, c, letters, digits) == 1{
          return 1
        }
        digits[i] = 0
        letters[b[len(b)-1]-65] = 0
      }
    }
    return -1
  }

  i = carry
  if len(a) > 0{
    i += letters[a[len(a)-1]-65]-1
    a = a[:len(a)-1]
  }
  if len(b) > 0{
    i += letters[b[len(b)-1]-65]-1
    b = b[:len(b)-1]
  }
  carry = i/10
  i = i%10

  if letters[c[len(c)-1]-65] == 0{
    if digits[i] != 0{
      return -1
    }
    letters[c[len(c)-1]-65] = i+1
    digits[i] = c[len(c)-1]
    if rec(carry, a, b, c[:len(c)-1], letters, digits) == -1{
      letters[c[len(c)-1]-65] = 0
      digits[i] = 0
      return -1
    }
    return 1
  }else{
    if i == letters[c[len(c)-1]-65]-1{
      return rec(carry, a, b, c[:len(c)-1], letters, digits)
    }
  }
  return -1
}

func f(b []byte){
    for i := 0; i < 10; i++{
      if b[i] > 0{
        fmt.Println(string(b[i]), i)
      }
    }
    fmt.Println()
}
