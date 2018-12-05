package main

import "fmt"

func main(){
	a := "absasdsam"
	b := "sam"
	j := 0; k := 0
	for i := 0; i<len(a) && j<len(b); i++{
		if a[i] == b[j]{
			j++
			if ((i<len(a)-1) && a[i+1]==b[k]){
				k++
			}else{
				k=0
			}
		}else{
			if k == 0 {
				j = 0
			}else if a[i] == b[k]{
				k=k
				j=k+1
			}else {
				k = k-1
				j = k+1
			}
		}
	}
	if j==len(b){
		fmt.Println("b exist in a")
	}
}
