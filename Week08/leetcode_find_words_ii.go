package Week08
//1. 前缀树查找
func findWords(board [][]byte, words []string) []string {
	if len(words) == 0 {
		return []string{}
	}
	n, m := len(board), len(board[0])
	root := &Trie{children: make([]*Trie, 26, 26)}
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	res := make([]string, 0)
	for _, w := range words {
		root.Insert(w)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			t := board[i][j]
			if root.children[t-byte('a')] != nil {
				dfs(board, i, j, root.children[t-byte('a')], dirs, &res)
			}
		}
	}
	return res
}
func dfs(board [][]byte, i, j int, root *Trie, dirs [][]int, res *[]string) {
	if root.word != "" {
		*res = append(*res, root.word)
		root.word = ""
	}
	origin := board[i][j]
	board[i][j] = byte('0')
	for _, d := range dirs {
		x, y := i+d[0], j+d[1]
		if x < 0 || x >= len(board) || y < 0 || y >= len(board[i]) || board[x][y] == byte('0') || root.children[board[x][y]-byte('a')] == nil {
			continue
		}
		dfs(board, x, y, root.children[board[x][y]-byte('a')], dirs, res)
	}
	board[i][j] = origin

}

type Trie struct {
	present  byte
	children []*Trie
	word     string
}

func (this *Trie) Insert(w string) {
	cur := this
	for _, c := range w {
		index := byte(c) - byte('a')
		if cur.children[index] == nil {
			cur.children[index] = &Trie{byte(c), make([]*Trie, 26, 26), ""}
		}
		cur = cur.children[index]
	}
	cur.word = w
}
