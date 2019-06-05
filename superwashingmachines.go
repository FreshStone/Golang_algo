package main

import "fmt"

func main(){
	a := [][]int{{1,0,5},{0,0,4,0,0,3,0,7,4},{0,0,4,0,3,0,5,4},{0,0,6,8,2,6,1,5,2,5,1,0},{0,0,0,3,0,5,0,4,15,3},{2,10,5,1,2},{0,0,4,10,0,4},{2,1,3,6,7,2,7},{5,2,8,4,1,0,9,5,0,0,4,1,0}}
	for _, v := range a{
		fmt.Println(findMinMoves(v))
	}
}

func findMinMoves(machines []int)int{
	if len(machines) == 1{
		return 0
	}

	var avg int
	for _, v := range machines{
		avg += v
	}

	if avg%len(machines) != 0{
		return -1
	}

	avg = avg/len(machines)
	var shifts, sum_deficit, sum_surplus, consecutive_surplus, consecutive_deficit, tmp_moves int
	var deficit, surplus, moves []int
	if machines[0] < avg{
		deficit = append(deficit, avg-machines[0])
		sum_deficit = deficit[0]
	}else{
		surplus = append(surplus, machines[0]-avg)
                sum_surplus = surplus[0]
	}

	i := 1
	for ;i < len(machines);{
		if len(deficit) > 0{
			if machines[i] <= avg{
				deficit = append(deficit, avg-machines[i])
				sum_deficit += avg-machines[i]
			}else{
				consecutive_surplus = 0
				tmp_moves = 0
				for{				//find consecutive surplus that are equal/greater than the total deficit
					if machines[i] < avg{
						i -= 1
						break
					}
					consecutive_surplus += machines[i]-avg
					if consecutive_surplus >= sum_deficit{
						break
					}
					i += 1
				}
				sum_deficit = max(sum_deficit - consecutive_surplus, 0)

				for len(deficit) > 0{		//removing consecutive surplus from deficit array
					if deficit[len(deficit)-1]  == -1{
						tmp_moves = max(moves[len(moves)-1], tmp_moves)
						moves = moves[:len(moves)-1]
						deficit = deficit[:len(deficit)-1]
					}

					if deficit[len(deficit)-1] > consecutive_surplus{
						tmp_moves += consecutive_surplus
						deficit[len(deficit)-1] -= consecutive_surplus
						deficit = append(deficit, -1)
						moves = append(moves, tmp_moves)
						break
					}else{
						tmp_moves += deficit[len(deficit)-1]
						consecutive_surplus -= deficit[len(deficit)-1]
						deficit = deficit[:len(deficit)-1]
					}
				}

				if len(deficit) == 0{
					shifts = max(shifts, tmp_moves)
					sum_surplus = consecutive_surplus
					surplus = []int{consecutive_surplus, -1}
					moves = []int{machines[i]-avg-consecutive_surplus}
				}
			}
		}else{
			if machines[i] >= avg{
                                surplus = append(surplus, machines[i]-avg)
				sum_surplus += machines[i]-avg
                        }else{
                                consecutive_deficit = 0
                                tmp_moves = 0
                                for{
                                        if machines[i] > avg{
                                                i -= 1
                                                break
                                        }
                                        consecutive_deficit += avg - machines[i]
                                        if consecutive_deficit >= sum_surplus{
                                                break
                                        }
                                        i += 1
                                }
				sum_surplus = max(sum_surplus-consecutive_deficit, 0)

                                for len(surplus) > 0{
                                        if surplus[len(surplus)-1]  == -1{
                                                tmp_moves = max(moves[len(moves)-1], tmp_moves)
                                                moves = moves[:len(moves)-1]
                                                surplus = surplus[:len(surplus)-1]
                                        }

                                        if surplus[len(surplus)-1] > consecutive_deficit{
                                                tmp_moves += consecutive_deficit
                                                surplus[len(surplus)-1] -= consecutive_deficit
                                                surplus = append(surplus, -1)
                                                moves = append(moves, tmp_moves)
                                                break
                                        }else{
                                                tmp_moves += surplus[len(surplus)-1]
                                                consecutive_deficit -= surplus[len(surplus)-1]
                                                surplus = surplus[:len(surplus)-1]
                                        }
                                }

				if len(surplus) == 0{
                                        shifts = max(shifts, tmp_moves)
					sum_deficit = consecutive_deficit
                                        deficit = []int{consecutive_deficit}
                                }
                        }
		}
		i += 1
	}
	return shifts
}

func max(i, j int)int{
	if i > j{
		return i
	}
	return j
}
