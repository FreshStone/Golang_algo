package main

import (
  "fmt"
  "math"
)
/*
Strings A and B are K-similar (for some non-negative integer K) if we can swap the positions of two letters in A exactly K times so that the resulting string equals B.

Given two anagrams A and B, return the smallest K for which A and B are K-similar.

Example 1:
Input: A = "ab", B = "ba"
Output: 1

Example 2:
Input: A = "abc", B = "bca"
Output: 2
*/

func main(){
  tests := [][]string{
    {"ab","ba"},
    {"abc","bca"},
    {"abac","baca"},
    {"aabc","abca"},
    {"abc","abc"},
    {"abc","acb"},
    {"abc","bac"},
    {"abc","cab"},
    {"abc","cba"},
    {"abcd","abcd"},
    {"abcd","abdc"},
    {"abcd","acbd"},
    {"abcd","acdb"},
    {"abcd","adcb"},
    {"abcd","adbc"},
    {"abcd","bacd"},
    {"abcd","badc"},
    {"abcd","bcad"},
    {"abcd","bcda"},
    {"abcd","bdac"},
    {"abcd","bdca"},
    {"abacbfcdedbcafbd","acdbcfdafbabcdeb"},
  }
  var i, j int
  results := []int{1,2,2,2,0,1,1,2,1,0,1,1,2,1,2,1,2,2,3,3,2,9}
  for i, _ = range tests{
    j = kSimilarity(tests[i][0], tests[i][1])
    if results[i] == j{
      fmt.Println("Correct")
    }else{
      fmt.Println(j, "is the Wrong answer", tests[i][0], tests[i][1])
    }
  }
}

func kSimilarity(A string, B string) int {
    if len(A) == 1{
        return 0
    }
    var i, j, l, c, res int
    b := make([]int, 6)
    bfs := [][]int{make([]int, 9)}
    for ;i < len(A); i++{
        if A[i] == B[i]{
            bfs[0][8] |= 1<<uint(i)
        }else {
            bfs[0][int(A[i]-97)] |= 1<<uint(i)
            b[int(B[i]-97)] |= 1<<uint(i)
        }
    }
    //fmt.Println(bfs[0], b)
    for{
      l = len(bfs)
      for i = 0; i < l; i++{
        if A[bfs[i][6]] == B[bfs[i][7]]{
          bfs[i][int(A[bfs[i][6]]-97)] ^= 1<<uint(bfs[i][6])
          bfs[i][8] |= 1<<uint(bfs[i][7])
          bfs[i][6] += 1
          for ;bfs[i][6] < len(A);{
            if (bfs[i][8] & (1<<uint(bfs[i][6]))) == 0{
              break
            }
            bfs[i][6] += 1
          }

          if bfs[i][6] == len(A){
            return res
          }
          bfs[i][7] = bfs[i][6]
        }
        bfs[i][8] |= 1<<uint(bfs[i][7])
        if bfs[i][int(B[bfs[i][7]]-97)] & b[int(A[bfs[i][6]]-97)] > 0{
          j = bfs[i][int(B[bfs[i][7]]-97)] & b[int(A[bfs[i][6]]-97)]
          j = int(math.Log2(float64(j^(j-1)))) //finding index of least significant bit
          bfs = append(bfs, []int{bfs[i][0],bfs[i][1],bfs[i][2],bfs[i][3],bfs[i][4],bfs[i][5],bfs[i][6],j,bfs[i][8]})
          bfs[len(bfs)-1][int(B[bfs[i][7]]-97)] ^= 1<<uint(j)
        }else{
          c = int(B[bfs[i][7]]-97)
          for j = bfs[i][6]+1; j < len(A); j++{
            if bfs[i][c] & (1<<uint(j)) > 0{
              bfs = append(bfs, []int{bfs[i][0],bfs[i][1],bfs[i][2],bfs[i][3],bfs[i][4],bfs[i][5],bfs[i][6],j,bfs[i][8]})
              bfs[len(bfs)-1][c] ^= 1<<uint(j)
            }
          }
        }
      }
      bfs = bfs[l:]
      res += 1
      //fmt.Println(bfs)
    }
    return res
}
