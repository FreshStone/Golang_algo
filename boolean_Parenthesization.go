package main

import "fmt"

func main(){
	a := "T|F&T^T"
	fmt.Println(f(a))
}

func f(a string) int{
	if len(a) == 1{
		if a == "T"{
			return 1
		}
		return 0
	}

	type n struct{
		n_true, n_total int
	}

	dp := make([][]n, (len(a)+1)/2)
	var j, tr, to int
	for l := 1; l <= len(dp); l++{
		dp[l-1] = make([]n, (len(a)+1)/2)
		j = 2*(l-1)
		for ;j < len(a);{
			if l == 1{
				if string(a[j]) == "T"{
					dp[0][j/2] = n{n_true: 1, n_total : 1}
				}else{
					dp[0][j/2] = n{n_true: 0, n_total : 1}
				}
			}else{
				tr = 0
				to = 0
				for k := 1; k < l; k++{
					if string(a[j-2*k+1]) == "|"{ //or operator
						tr += dp[k-1][j/2].n_true*dp[l-k-1][(j/2)-k].n_total + (dp[k-1][j/2].n_total-dp[k-1][j/2].n_true)*dp[l-k-1][(j/2)-k].n_true
					}else if string(a[j-2*k+1]) == "&"{ //and operator
						tr += dp[k-1][j/2].n_true*dp[l-k-1][(j/2)-k].n_true
					}else{ //xor operator
						tr += dp[k-1][j/2].n_true*(dp[l-k-1][(j/2)-k].n_total-dp[l-k-1][(j/2)-k].n_true) + (dp[k-1][j/2].n_total-dp[k-1][j/2].n_true)*dp[l-k-1][(j/2)-k].n_true
					}

					to += dp[l-k-1][(j/2)-k].n_total*dp[k-1][j/2].n_total
				}
				dp[l-1][j/2] = n{n_true: tr, n_total: to}
			}
			j += 2
		}
	}
	return dp[len(dp)-1][len(dp)-1].n_true
}







