package main

import (
  "fmt"
  "strings"
)

func main(){
  s := "any carnal pleasure"
  s = base64encode(s)
  fmt.Println(base64decode(s))
}


//No Padding used
//["A"-"Z", "a"-"z", "0"-"9", "+", "/"]
func base64encode(s string)string{
  if len(s) == 0{
    return ""
  }
  var i int
  var b strings.Builder
  for i = 2 ;i < len(s);{
    b.WriteByte(charEncode(s[i-2]/4))
    b.WriteByte(charEncode(16*(s[i-2]%4) + s[i-1]/16))
    b.WriteByte(charEncode(4*(s[i-1]%16) + s[i]/64))
    b.WriteByte(charEncode(s[i]%64))
    i += 3
  }

  if i - len(s) == 1{
    i -= 2
    b.WriteByte(charEncode(s[i]/4))
    b.WriteByte(charEncode(16*(s[i]%4)))
  }else if i == len(s){
    i -= 1
    b.WriteByte(charEncode(s[i-1]/4))
    b.WriteByte(charEncode(16*(s[i-1]%4) + s[i]/16))
    b.WriteByte(charEncode(4*(s[i]%16)))
  }
  return b.String()
}

func base64decode(s string)string{
  if len(s) == 0{
    return ""
  }
  var i int
  var b strings.Builder
  for i = 3; i < len(s);{
    b.WriteByte(charDecode(s[i-3])*4 + charDecode(s[i-2])/16)
    b.WriteByte((charDecode(s[i-2])%16)*16 + charDecode(s[i-1])/4)
    b.WriteByte((charDecode(s[i-1])%4)*64 + charDecode(s[i]))
    i += 4
  }

  if i - len(s) == 1{
    i -= 2
    b.WriteByte(charDecode(s[i-1])*4 + charDecode(s[i])/16)
  }else if i == len(s){
    i -= 1
    b.WriteByte(charDecode(s[i-2])*4 + charDecode(s[i-1])/16)
    b.WriteByte((charDecode(s[i-1])%16)*16 + charDecode(s[i])/4)
  }
  return b.String()
}

func charEncode(b byte)byte{
  if b < 26{
    return 65+b
  }else if b < 52{
    return 71 + b          //97 + b - 26
  }else if b < 62{
    return b - 4          //48 + b - 52
  }else if b == 62{  //"+"
    return 43
  }else{          //"="
    return 47
  }
  return 0
}

func charDecode(b byte)byte{
  if b > 96{
    return b - 71 //b-97 + 26
  }else if b > 64{
    return b-65
  }else if b > 47{
    return b + 4 //b-48 + 52
  }else if b == 43{
    return 62
  }else{
    return 63
  }
  return 0
}
