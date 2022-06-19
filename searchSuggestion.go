package main

import (
	"fmt"
	"sort"
)

func main() {
	tests := [][]string{
		{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
		{"havana"},
		{"abc", "abcd", "dbc"},
	}
	words := []string{"mouse", "havan", "ab"}
	for i := 0; i < len(tests); i++ {
		fmt.Println(suggestedProducts(tests[i], words[i]))
	}
}

/*leetcode -1268
Linear search for finding start & end for letter searchword[i] can be replaced with Binary search
in the range(old_start, old_end) to find start & end occurences for each prefixes of the searchword.
  //for next iteration  (old_start=start;old_end=end)
*/

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	var found bool
	var i, start, j int
	end := len(products)
	ans := make([][]string, len(searchWord))
	for ; i < len(searchWord); i++ {
		found = false
		for j = start; j < end; j++ {
			if len(products[j]) > i && products[j][i] == searchWord[i] {
				if !found {
					found = true
					start = j
				}
				if len(ans[i]) < 3 {
					ans[i] = append(ans[i], products[j])
				}
			} else if found {
				break
			}
		}
		if j == end && !found {
			break
		}
		end = j
	}
	return ans
}
