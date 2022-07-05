package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main(){
	var n, l int
	var line string
        fmt.Scanf("%d", &n)
        r := bufio.NewReader(os.Stdin)
        for i := 0; i < n; i++{
		fmt.Scanf("%d", &l)

		var a [][]int
		for ;l > 0;{
			line, _ = r.ReadString('\n')
			a = append(a, converttointarr(line))
			l -= 1
		}
		fmt.Println(f(a))
	}
}

func converttointarr(s string) []int {
    var n []int
    for _, f := range strings.Fields(s) {
        i, err := strconv.Atoi(f)
        if err == nil {
            n = append(n, i)
        }
    }
    return n
}

func f(a [][]int)int{
	var mstore []int
	var maxquality, v int

	tmp := (len(a)-1)*a[len(a)-1][0]
	mstore = append(mstore, []int{-tmp, tmp, tmp, -tmp}...)
	for j := 1; j < len(a[len(a)-1]); j++{
		tmp = a[len(a)-1][j]*(len(a)-1)
		if mstore[0] + tmp < 0{
			mstore[0] = -tmp
			mstore[1] = tmp
		}

		if tmp > mstore[2]{
			mstore[2] = tmp
			mstore[3] = -tmp
		}
	}

	for i := len(a)-2; i > -1; i--{
		v = max(a[i][0]*(i+1)+mstore[0], -a[i][0]*(i+1)+mstore[1], -a[i][0]*(i+1)+mstore[2], a[i][0]*(i+1)+mstore[3])
		if v > maxquality{
			maxquality = v
		}
		tmp = a[i][1]*i
		mstore = append(mstore, []int{v-tmp, v+tmp, v+tmp, v-tmp}...)

		for j := 1; j < len(a[i]); j++{
			v = max(a[i][j]*(i+1)+mstore[0], -a[i][j]*(i+1)+mstore[1], -a[i][j]*(i+1)+mstore[2], a[i][j]*(i+1)+mstore[3])
			if v > maxquality{
				maxquality = v
			}

			tmp = a[i][(j+1)%len(a[i])]*i
			if v - tmp > mstore[4]{
				mstore[4] = v - tmp
				mstore[5] = v + tmp
			}

			if v + tmp > mstore[6]{
				mstore[6] = v + tmp
				mstore[7] = v - tmp
			}
		}
		mstore = mstore[4:]
	}
	return maxquality
}


func max(a...int)int{
	var m int
	for _, v := range a{
		if v > m{
			m = v
		}
	}
	return m
}
