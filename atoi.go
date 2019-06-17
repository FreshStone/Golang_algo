package main

import "fmt"

func main(){
	str := "   +2147483647"
	fmt.Println(atoi(str))
}

func atoi(str string)int{ //32 bit int
	var num_started, positive bool
	var res, i int
	tmp := (1<<31 - 1)/10
	for ; i < len(str);{
		if num_started{
			if str[i] > 47 && str[i] < 58 {
				if positive{
					if res < tmp || (res == tmp && str[i]-48 < 8){
						res = res*10 + int(str[i]-48)
					}else{
						res = 1<<31 - 1
						break
					}
				}else{
					if res < tmp || (res == tmp && str[i]-48 < 9){
                                                res = res*10 + int(str[i]-48)
                                        }else{
                                                res = 1<<31
                                                break
                                        }
				}
				i += 1
			}else{
				break
			}
		}else{
			if string(str[i]) == "-"{
				num_started = true
				i += 1
			}else if string(str[i]) == "+"{
				num_started = true
                                i += 1
				positive = true
			}else if str[i] > 47 && str[i] < 58 {
				 num_started = true
				 positive = true
			}else{
				i += 1
			}
		}
	}
	if !positive{
		return -1*res
	}

	return res
}
