// Count all permuatations of BST having height(h) and containing n nodes(val- [1,n])

package main

import (
	"fmt"
	"math"
)

func main(){
	n := 5
	h := 3
	fmt.Println(countperms(n, h))
}

func countperms(n, max_height int)int{
	type e struct{
		l int
		u int
		cnt []int
	}

	dp := make([][]e, n)
	i := 0
	j := 0
	var h, k, c_tmp, le, ri int
	for ;j < len(dp);{
		for ;j < len(dp);{
			if len(dp[i]) == 0{
				dp[i] = make([]e, len(dp))
			}
			dp[i][j] = e{l: int(math.Log2(float64(j-i+1))), u: min(j-i, max_height), cnt: make([]int, max_height+2)}// height = i-1
			if dp[i][j].l > max_height{
				i += 1
				j += 1
				continue
			}

			for h = dp[i][j].l; h <= dp[i][j].u; h++{
				c_tmp = 0
				if i == j{
					c_tmp = 1
				}else{
					if dp[i+1][j].cnt[h] > 0{
						c_tmp += dp[i+1][j].cnt[h] - dp[i+1][j].cnt[h-1]
					}

					if dp[i][j-1].cnt[h] > 0{
						c_tmp += dp[i][j-1].cnt[h] - dp[i][j-1].cnt[h-1]
					}
				}

				for k = i+1; k < j; k++{
					if dp[i][k-1].cnt[h] > 0{
						le = dp[i][k-1].cnt[h] - dp[i][k-1].cnt[h-1]
					}else{
						le = 0
					}

					if dp[k+1][j].cnt[h] > 0{
						ri = dp[k+1][j].cnt[h] - dp[k+1][j].cnt[h-1]
					}else{
						ri = 0
					}

					c_tmp += combination(j-i, k-i)*(dp[i][k-1].cnt[min(h-1, dp[i][k-1].u+1)]*ri + le*dp[k+1][j].cnt[min(h-1, dp[k+1][j].u+1)] + le*ri)
				}
				dp[i][j].cnt[h+1] = dp[i][j].cnt[h]+c_tmp
			}
			i += 1
			j += 1
		}
		j = len(dp)-i+1
		i = 0
	}
	if dp[0][len(dp)-1].cnt[max_height+1] > 0{
		return dp[0][len(dp)-1].cnt[max_height+1]-dp[0][len(dp)-1].cnt[max_height]
	}

	return 0
}

func min(a, b int)int{
	if a > b{
		return b
	}
	return a
}

func combination(a, b int)int{
	i := a-b+1
	j := 1
	num, den := 1, 1
	for ;i <= a; i++{
		if j <= b{
			den = den*j
			j += 1
		}
		num = num*i
	}

	return num/den
}

