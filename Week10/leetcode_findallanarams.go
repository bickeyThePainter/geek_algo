package Week10
// 找到所有s中的p的异位词
func findAnagrams(s string, p string) []int {
	check := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		check[p[i]]++
	}
	t := len(p)
	r := 0
	res := make([]int, 0)
	for ; r < len(s); r++ {
		if _, ok := check[s[r]]; ok {
			check[s[r]]--
		}
		if r >= t-1 {
			if r-t >= 0 {
				if _, ok := check[s[r-t]]; ok {
					check[s[r-t]]++
				}
			}

			found := true
			for _, v := range check {
				if v != 0 {
					found = false
				}
			}
			if found {
				res = append(res, r-t+1)
			}

		}

	}
	return res
}
