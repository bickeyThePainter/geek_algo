package Week06_07

import "container/list"

// 1. bfs
// 2. bi bfs

func ladderLength1(beginWord string, endWord string, wordList []string) int {
	check := make(map[string]struct{})
	for _, w := range wordList {
		check[w] = struct{}{}
	}
	if _, ok := check[endWord]; !ok {
		return 0
	}
	visited := make(map[string]struct{})
	q := list.New()
	q.PushBack(beginWord)
	visited[beginWord] = struct{}{}
	level := 1
	for q.Len() > 0 {
		length := q.Len()
		for i := 0; i < length; i++ {
			cur := q.Remove(q.Front()).(string)
			cs := []rune(cur)
			for i := 0; i < len(cs); i++ {
				origin := cs[i]
				for a := 'a'; a <= 'z'; a++ {
					if a == origin {
						continue
					}
					cs[i] = a
					tmp := string(cs)
					if tmp == endWord {
						return level + 1
					}
					if _, ok := check[tmp]; ok {
						if _, v := visited[tmp]; v {
							continue
						}
						// delete(check,tmp) 也可
						visited[tmp] = struct{}{}
						q.PushBack(tmp)
					}
				}
				cs[i] = origin
			}

		}
		level++
	}
	return 0
}

func ladderLength2(beginWord string, endWord string, wordList []string) int {
	check := make(map[string]struct{})
	for _, w := range wordList {
		check[w] = struct{}{}
	}
	if _, ok := check[endWord]; !ok {
		return 0
	}
	beginSet := make(map[string]struct{})
	endSet := make(map[string]struct{})
	beginSet[beginWord] = struct{}{}
	endSet[endWord] = struct{}{}
	level := 1
	for len(beginSet) > 0 && len(endSet) > 0 {
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}
		tmp := make(map[string]struct{})
		for w, _ := range beginSet {
			cs := []rune(w)
			for i := 0; i < len(cs); i++ {
				origin := cs[i]
				for a := 'a'; a <= 'z'; a++ {
					if origin == a {
						continue
					}
					cs[i] = a
					cur := string(cs)
					if _, end := endSet[cur]; end {
						return level + 1
					}
					if _, ok := check[cur]; ok {
						tmp[cur] = struct{}{}
						delete(check, cur)
					}
				}
				cs[i] = origin
			}
		}
		beginSet = tmp
		level++
	}
	return 0
}
