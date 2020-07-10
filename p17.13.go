package main

func respace(dictionary []string, sentence string) int {
	head := &Trie{
		next: [26]*Trie{},
	}
	for _, str := range dictionary {
		node := head
		for i := len([]byte(str)) - 1; i >= 0; i -- {
			b := []byte(str)[i]
			if node.next[b - 'a'] == nil {
				node.next[b - 'a'] = &Trie{
					next: [26]*Trie{},
				}
			}
			node = node.next[b - 'a']
		}
		node.isEnd = true
	}
	length := len(sentence)
	dp := make([]int, length + 1)
	dp[0] = 0
	for i := 1; i <= length; i ++ {
		dp[i] = length
	}

	for i := 1; i <= len(sentence); i ++ {
		dp[i] = dp[i - 1] + 1
		node := head
		for j := i; j > 0; j -- {
			if node.next[sentence[j - 1] - 'a'] != nil {
				if node.next[sentence[j - 1] - 'a'].isEnd {
					dp[i] = min(dp[j - 1], dp[i])
				}
				node = node.next[sentence[j - 1] - 'a']
			} else {
				break
			}
		}
	}
	return dp[length]
}

type Trie struct {
	next [26]*Trie
	isEnd bool
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}