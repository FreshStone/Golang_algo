package main

import "fmt"

type WordFilter struct {
	pre     map[byte]*WordFilter
	suf     map[byte]*WordFilter
	indexes []int
}

func main() {
	words := []string{"apple", "attitude"}
	word_filter := Constructor(words)
	fmt.Println(word_filter.F("a", "e"))
}

/* leetcode - 745
optimisation to be added:
L = len(words[i]); L <= 10
N = len(words); N <=15000
since L <11; all possible combinations of prefix's and suffix's of a word can be added to
single trie like for eg- "apple" -{{"apple#a","apple#ap".."apple#apple"},
									{"appl#a","appl#ap".."appl#apple"},
									.
									.
									{"a#a","a#ap".."a#apple"}}
above approach will increase time complexity of the constructor func from O(N*L) to O(N*L*L)
but will greatly reduce tc of F() from O(N+L) to O(L)
*/
func Constructor(words []string) WordFilter {
	var i, j int
	var tmp, tmp_pre, tmp_suf *WordFilter
	var ok bool
	trie_node := WordFilter{pre: map[byte]*WordFilter{}, suf: map[byte]*WordFilter{}}
	for i = len(words) - 1; i >= 0; i-- {
		tmp_pre = &trie_node
		tmp_suf = &trie_node
		for j = 0; j < len(words[i]); j++ {
			tmp, ok = tmp_pre.pre[words[i][j]]
			if !ok {
				tmp = &WordFilter{pre: map[byte]*WordFilter{}}
				tmp_pre.pre[words[i][j]] = tmp
			}
			tmp_pre = tmp
			tmp_pre.indexes = append(tmp_pre.indexes, i)
			tmp, ok = tmp_suf.suf[words[i][len(words[i])-1-j]]
			if !ok {
				tmp = &WordFilter{suf: map[byte]*WordFilter{}}
				tmp_suf.suf[words[i][len(words[i])-1-j]] = tmp
			}
			tmp_suf = tmp
			tmp_suf.indexes = append(tmp_suf.indexes, i)
		}
	}
	return trie_node
}

func (this *WordFilter) F(prefix string, suffix string) int {
	var tmp_pre, tmp_suf *WordFilter
	var ok bool
	var i, j int
	tmp_pre = this
	tmp_suf = this
	for ; i < len(prefix); i++ {
		tmp_pre, ok = tmp_pre.pre[prefix[i]]
		if !ok {
			return -1
		}
	}
	for j = len(suffix) - 1; j >= 0; j-- {
		tmp_suf, ok = tmp_suf.suf[suffix[j]]
		if !ok {
			return -1
		}
	}
	for i, j = 0, 0; i < len(tmp_pre.indexes) && j < len(tmp_suf.indexes); {
		if tmp_pre.indexes[i] > tmp_suf.indexes[j] {
			i += 1
		} else if tmp_pre.indexes[i] < tmp_suf.indexes[j] {
			j += 1
		} else {
			return tmp_pre.indexes[i]
		}
	}
	return -1
}
