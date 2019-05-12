package main

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
	"math"
)

func main(){
	var n int
	var l, r int64
	testarr := make([][]string, n)
	fmt.Scanf("%d", &n)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++{
		testarr = append(testarr, strings.Split(readline(reader), " "))
	}

	for i := 0; i < n; i++{
		l, _ = strconv.ParseInt(testarr[i][0], 10, 64)
		r, _ = strconv.ParseInt(testarr[i][1], 10, 64)
		fmt.Println(f(l, r))
	}
}

func readline(r *bufio.Reader)string{
	s, _, _ := r.ReadLine()
	return strings.TrimRight(string(s), "\r\n")
}

func f(l, r int64)int64{
    var cnt int64
    le := int64(math.Log10(float64(l)))
    ri := int64(math.Log10(float64(r)))
    tmp := int64(math.Pow(10, float64(le)))
    Lmsn := l/tmp
    Rmsn := r/int64(math.Pow(10, float64(ri)))
    if Rmsn == Lmsn && le == ri{
        if Rmsn%2 == 0{
            return r-l+1
        }
    }else{
        if Lmsn%2 == 0{
            cnt += (Lmsn+1)*tmp - l
        }
    }
    if ri == le{
        if Rmsn%2 == 0{
            return cnt + ((Rmsn-Lmsn-1)/2)*tmp + r%tmp + 1
        }else{
            return cnt + ((Rmsn-Lmsn)/2)*tmp
        }
    }else{
        cnt += ((9-Lmsn)/2)*tmp
    }
    tmp = tmp*10
    for le = le+1; le < ri; le++{
        cnt += 4*tmp
	tmp = tmp*10
    }
    
    cnt += ((Rmsn-1)/2)*tmp
    if Rmsn%2 == 0{
        cnt += r - Rmsn*tmp + 1
    }
    return cnt
}
