package main

import "fmt"

func main(){
  //only interger co_efficients considered
 s := "x+y+w=1\ny=2\nz-6+y=5\n2z+2w-3=18-z\n"; n := 4
// s := "3a+2b-4c=3\n2a+3b+3c=15\n5a-3b+c=14\n"; n := 3
 fmt.Println(solveEqn(s, n))
}

//0x + y = 1; zero before variable in an equation is not considered

func solveEqn(s string, n int)map[byte]float64{
  rhs := make([]float64, n)
  //res := make([]int, n)
  co_eff := make([][]float64, n)
  m := make(map[byte]int)
  var k, e bool
  var i, j, t, c int
  p := true
  co_eff[0] = make([]float64,n)
  for ;i < len(s); i++{
    if s[i] == 10{ //new line; new equation signal
      if c != 0{
        if p{
          rhs[j] += float64(c)
        }else{
          rhs[j] -= float64(c)
        }
      }
      e = false
      p = true
      c = 0
      j += 1
      if j < n{
        co_eff[j] = make([]float64, n)
      }
    }else if s[i] > 47 && s[i] < 58{ //number
      c = c*10 + int(s[i]-48)
    }else if s[i] == 43{ //"+"
      if (e && !p) || (!e && p){
        rhs[j] -= float64(c)
      }else{
        rhs[j] += float64(c)
      }
      p = true
      c = 0
    }else if s[i] == 45{ //"-"
      if (e && !p) || (!e && p){
        rhs[j] -= float64(c)
      }else{
        rhs[j] += float64(c)
      }
      p = false
      c = 0
    }else if s[i] == 61{ //"="
      if (e && !p) || (!e && p){
        rhs[j] -= float64(c)
      }else{
        rhs[j] += float64(c)
      }
      e = true
      p = true
      c = 0
    }else{ //char_variable
      _, k = m[s[i]]
      if !k{
        m[s[i]] = len(m)
      }
      if c == 0{ //x + y = 1; coefficient of x == 1
        c = 1
      }
      //fmt.Println(j, s[i],m[s[i]])
      if (e && !p) || (!e && p){
        co_eff[j][m[s[i]]] += float64(c)
      }else{
        co_eff[j][m[s[i]]] -= float64(c)
      }
      c = 0
    }
  }
//  printmatrix(co_eff)
//  fmt.Println(rhs)
  /*
  Gauss elimination method, turn co_eff into an upper triangular matirx using
  row operations (also perform the same operations on rhs) and than solve for each
  variable starting from the end row(i == n-1)
  */

  for i = 0; i < n-1; i++{
    j = i+1
    c = i
    if co_eff[i][c] == 0{
      k = false
      //find row with non zero value at col == c if it doesnt exist than return
      for ;j < n; j++{
        if co_eff[j][c] != 0{
          k = true
          break
        }
      }
      if k{
        //swap i and j rows
        rhs[j], rhs[i] = rhs[i], rhs[j]
        for t = c; t < n; t++{ //row operation
          co_eff[i][t], co_eff[j][t] = co_eff[j][t], co_eff[i][t]
        }
        j = j+1
      }else{
        return map[byte]float64{} //infinite or no solution case
      }
    }

    for ;j < n; j++{
      if co_eff[j][c] == 0{
        continue
      }
      co_eff[j][c] /= co_eff[i][c]
      rhs[j] -= co_eff[j][c]*rhs[i]
      for t = n-1; t > c; t--{//row operation
        co_eff[j][t] -= co_eff[j][c]*co_eff[i][t]
      }
      co_eff[j][c] = 0
    }
  }
  fmt.Println("row echleon form")
  printmatrix(co_eff)
  fmt.Println("rhs", rhs)
  for i = n-1; i >= 0; i--{
    for j = i+1; j < n; j++{
      rhs[i] -= co_eff[i][j]*rhs[j]
    }
    rhs[i] /= co_eff[i][i]
  }
  res := make(map[byte]float64)
  for b, i := range m{
    res[b] = rhs[i]
  }
  return res
}

func printmatrix(m [][]float64){
  for i := 0; i < len(m); i++{
    fmt.Println(m[i])
  }
  return
}
