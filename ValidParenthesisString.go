package main

import "fmt"

func main(){
	var s [14]string
	s[0] = "(*(()))((())())*(**(()))((*)()(()))*(())()(())(()" // false 
	s[1] = "****(((((((((******))" //false
	s[2] = "(())((())()()(*)(*()(())())())()()((()())((()))(*" // false
	s[3] = "*****(*****)" // true
	s[4] = "((*)" //true
	s[5] = "((*))((*" //false
	s[6] = "((*****((((*)" //false
	s[7] = "(((*****((*" //false
	s[8] = "(((*****((*********" //true
	s[9] = "(((*)((**" //false
	s[10] = "*()(())*()(()()((()(()()*)(*(())((((((((()*)(()(*)" //false
	s[11] = "(*)(((**(*" //false
	s[12] = "(((()*())))((()(((()(()))()**(*)())))())()()*" //true
	s[13] = "((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()" //true
	for _, v := range s{
		fmt.Println(validParenthesis(v))
	}
}

func validParenthesis(s string)bool{
	if len(s) == 0{
                return true
        }

        if s[0] == 41 || s[len(s)-1] == 40{
                return false
        }
        var prev_open, prev_open_old, next_close, next_open, prev_star, curr_star, diff int
	var star_started bool
        for i := 0; i < len(s); i++{
		if s[i] == 42{
			if star_started{
				curr_star += 1
				continue
			}

			if prev_open < next_close{
                                if prev_open >= next_close-prev_star-curr_star{
					prev_star += curr_star+prev_open-next_close
					prev_open = next_open
					prev_open_old = 0
					diff = 0
                                }else{
                                        return false
                                }
                        }else{
				if prev_open - prev_open_old > next_close+curr_star{
					diff += prev_open - prev_open_old-next_close-curr_star
			        }else{
					diff -= min(diff, next_close+curr_star - prev_open + prev_open_old)
				}
				prev_open_old = prev_open-next_close
                                prev_open += next_open-next_close
				prev_star += curr_star
			}
			curr_star = 1
			star_started = true
                        next_close = 0
                        next_open = 0
                }else{
                        if s[i] == 40{
                                next_open += 1
                        }else{
                                if next_open > 0{
                                        next_open -= 1
                                }else{
                                        next_close += 1
                                }
                        }
			star_started = false
                }
        }

	if star_started{
		if diff+prev_open-prev_open_old > curr_star{
			return false
		}
	}else{
		diff += prev_open - prev_open_old-next_close-curr_star
		if next_open > 0 || prev_open < next_close-prev_star-curr_star || diff > 0{
			return false
		}
	}

        return true
}

func min(a, b int)int{
	if a < b{
		return a
	}
	return b
}
