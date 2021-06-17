package Week10
// reverse -> reverse word end with ' '
func reverseWords(s string) string {
	_re := func(cs []rune, l, r int) {
		for l < r {
			cs[l], cs[r] = cs[r], cs[l]
			l++
			r--
		}
	}
	cs := []rune(s)
	l, n := 0, len(cs)
	_re(cs, 0, n-1)
	for ; l < n && cs[l] == ' '; l++ {

	}
	r := l
	cache := make([]rune, 0)
	for ; r < n; r++ {
		if cs[r] == ' ' {
			_re(cs, l, r-1)
			if l < r {
				cache = append(cache, cs[l:r]...)
				cache = append(cache, ' ')
			}
			l = r + 1
		}
	}
	if l < n {
		_re(cs, l, n-1)
		cache = append(cache, cs[l:n]...)
	}
	if cache[len(cache)-1] == ' ' {
		return string(cache[0 : len(cache)-1])
	}
	return string(cache)
}
