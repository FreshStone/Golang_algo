package main

import(
	"fmt"
	"strings"
	"sort"
)

func main(){
	tests := [][]int{
		{1,4,7,8,2,1},
		{1,2,5,4,4,3,2,2,1,2,3},
		{21,40,26,33,31},
		{1,1,2},
	}
	ans := []int{1,1,5,0}
	for i := 0 ; i < len(tests); i++{
		if ans[i] != lastStoneWeightII(tests[i]){
			fmt.Println("error", "expected-", ans[i], "returning-", lastStoneWeightII(tests[i]))
		}else{
			fmt.Println("correct answer")
		}
	}
}
/*
leetcode 1049
knapsack approach
divide the stones into 2 subsets such that the difference between the sum of elements
of each subset is minimum
*/

func lastStoneWeightII(stones []int)int{
	if len(stones) == 1{
		return stones[0]
	}
	if len(stones) == 2{
		if stones[0] < stones[1]{
			return stones[1]-stones[0]
		}
		return stones[0]-stones[1]
	}
	var i, j, id, ans, total int
	for ; i < len(stones);i++{
		total += stones[i]
	}
	dp := make([][]bool, 2)
	dp[0] = make([]bool, 1+(total/2)); dp[0][0] = true
	dp[1] = make([]bool, len(dp[0])); dp[1][0] = true
	for i = 0; i < len(stones); i++{
		id = (i+1)%2
		for j = 1; j < len(dp[0]); j++{
			dp[id][j] = dp[1-id][j]
			if stones[i] <= j{
				dp[id][j] = dp[id][j] || dp[1-id][j-stones[i]]
			}
			if dp[id][j]{
				ans = j
			}
		}
	}
	return total - 2*ans
}

func lastStoneWeightII_rec(stones []int)int{
	if len(stones) == 1{
		return stones[0]
	}
	if len(stones) == 2{
		return abs(stones[0],stones[1])
	}
	var ok bool
	var t, o byte
  var i int
  var b strings.Builder
  store := [][]int{} //store[i] = []int{stone_weight, feq}
  m := make(map[string]int)
  char := make([][]byte, 101)
  for t = 48; t < 58; t++{
    for o = 48; o < 58; o++{
      if t == 48{
        char[i] = []byte{o}
      }else{
        char[i] = []byte{t, o}
      }
      i++
    }
  }
  char[100] = []byte{49,48,48}
  sort.Ints(stones)
  store = append(store,[]int{stones[0],1})
  for i = 1; i < len(stones);i++{
  	if stones[i] == stones[i-1]{
  		store[len(store)-1][1] += 1
  	}else{
  		store = append(store, []int{stones[i],1})
  	}
  }

  var rec func(int,int,int)int
  rec = func(s int, e int, w int)int{
  	//fmt.Println(s,e,w,store)
  	if s == e{
  		if store[s][1]%2 == 0{
  			return w
  		}
  		return abs(store[s][0],w)
  	}
  	var ans, i, j, v int
  	var str string
  	b.Write(char[w])
  	for j = s; j < e+1; j++{
  		b.Write(char[store[j][0]])
  		if store[j][1] > 1{
  			b.Write(char[store[j][1]])
  		}
  	}
  	str = b.String()
  	b.Reset()
  	v, ok = m[str]
  	if ok {
  		return v
  	}
  	ans = 101
  	for i = s; i < e+1; i++{
  		if store[i][1] > 1{
  			store[i][1] -= 1
  			ans = min(ans, rec(s,e,abs(w,store[i][0])))
  			store[i][1] += 1
  		}else{
  			v = store[i][0]
  			for j = i; j < e; j++{
  				store[j][0] = store[j+1][0]
  				store[j][1] = store[j+1][1]
  			}
  			ans = min(ans, rec(s,e-1,abs(w,v)))
  			for j = e-1; j > i; j--{
  				store[j][0] = store[j-1][0]
  				store[j][1] = store[j-1][1]
  			}
  			store[i][0] = v
  			store[i][1] = 1
  		}
  	}
  	m[str] = ans
  	return ans
  }
  return rec(0,len(store)-1,0)
}

func min(a, b int)int{
	if a < b{
		return a
	}
	return b
}

func abs(a, b int)int{
	if a < b{
		return b-a
	}
	return a-b
}