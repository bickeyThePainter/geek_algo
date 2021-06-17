package Week10

import "container/list"

// two pointers
func reverseOnlyLetters(s string) string {
	l, r := 0, len(s)-1
	cs := []rune(s)
	for l < r {
		for l < r && !(cs[l] >= 'a' && cs[l] <= 'z' || cs[l] >= 'A' && cs[l] <= 'Z') {
			l++
		}
		for l < r && !(cs[r] >= 'a' && cs[r] <= 'z' || cs[r] >= 'A' && cs[r] <= 'Z') {
			r--
		}
		if l < r {
			cs[l], cs[r] = cs[r], cs[l]
			l++
			r--
		}
	}
	return string(cs)
}

// stack
func reverseOnlyLetters1(s string) string {
	q:=list.New()
	for _,c:=range s{
		if c>='a' && c<='z' || c>='A' && c<='Z'{
			q.PushFront(c)
		}
	}
	cs:=make([]rune,len(s))
	for i,c:=range s{
		if c>='a' && c<='z' || c>='A' && c<='Z'{
			cs[i]=q.Remove(q.Front()).(rune)
		}else{
			cs[i]=c
		}
	}
	return string(cs)

}