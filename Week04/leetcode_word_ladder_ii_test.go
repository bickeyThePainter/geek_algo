package Week04

import (
	"container/list"
	"testing"
)

// bfs + dfs
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	set := make(map[string]struct{})
	for _, w := range wordList {
		set[w] = struct{}{}
	}
	if _, ok := set[endWord]; !ok {
		return [][]string{}
	}
	check := make(map[string][]string)
	q := list.New()
	q.PushBack(beginWord)
	found := false
	dis := make(map[string]int)
	dis[beginWord] = 0
	for q.Len() > 0 {
		count := q.Len()
		for i := 0; i < count; i++ {
			cur := q.Remove(q.Front()).(string)
			if _, ok := check[cur]; ok {
				continue
			}
			words := findWords(set, cur)
			check[cur] = words
			for _, w := range words {
				if _, ok := dis[w]; ok {
					continue
				}
				dis[w] = dis[cur] + 1
				if w == endWord {
					found = true
				} else {
					q.PushBack(w)
				}
			}
		}
		if found {
			break
		}
	}
	res := make([][]string, 0)
	dfs(check, &res, make([]string, 0), beginWord, endWord, dis)
	return res

}
func findWords(set map[string]struct{}, word string) []string {
	cs := []rune(word)
	res := make([]string, 0)
	for i := 0; i < len(cs); i++ {
		origin := cs[i]
		for a := 'a'; a <= 'z'; a++ {
			if origin == a {
				continue
			}
			cs[i] = a
			tmp := string(cs)
			if _, ok := set[tmp]; !ok {
				continue
			}
			res = append(res, tmp)
		}
		cs[i] = origin
	}
	return res
}

func dfs(check map[string][]string, res *[][]string, prev []string, cur string, endWord string, dis map[string]int) {
	prev = append(prev, cur)
	if cur == endWord {
		_res := make([]string, len(prev), len(prev))
		copy(_res, prev)
		*res = append(*res, _res)
		return
	}
	for _, c := range check[cur] {
		if dis[c] == dis[cur]+1 {
			dfs(check, res, prev, c, endWord, dis)
		}
	}
	prev = prev[:len(prev)-1]

}

func TestWords(t *testing.T){
	findLadders("cet",
	"ism",
	[]string{"kid","tag","pup","ail","tun","woo","erg","luz","brr","gay","sip","kay","per","val","mes","ohs","now","boa","cet","pal","bar","die","war","hay","eco","pub","lob","rue","fry","lit","rex","jan","cot","bid","ali","pay","col","gum","ger","row","won","dan","rum","fad","tut","sag","yip","sui","ark","has","zip","fez","own","ump","dis","ads","max","jaw","out","btu","ana","gap","cry","led","abe","box","ore","pig","fie","toy","fat","cal","lie","noh","sew","ono","tam","flu","mgm","ply","awe","pry","tit","tie","yet","too","tax","jim","san","pan","map","ski","ova","wed","non","wac","nut","why","bye","lye","oct","old","fin","feb","chi","sap","owl","log","tod","dot","bow","fob","for","joe","ivy","fan","age","fax","hip","jib","mel","hus","sob","ifs","tab","ara","dab","jag","jar","arm","lot","tom","sax","tex","yum","pei","wen","wry","ire","irk","far","mew","wit","doe","gas","rte","ian","pot","ask","wag","hag","amy","nag","ron","soy","gin","don","tug","fay","vic","boo","nam","ave","buy","sop","but","orb","fen","paw","his","sub","bob","yea","oft","inn","rod","yam","pew","web","hod","hun","gyp","wei","wis","rob","gad","pie","mon","dog","bib","rub","ere","dig","era","cat","fox","bee","mod","day","apr","vie","nev","jam","pam","new","aye","ani","and","ibm","yap","can","pyx","tar","kin","fog","hum","pip","cup","dye","lyx","jog","nun","par","wan","fey","bus","oak","bad","ats","set","qom","vat","eat","pus","rev","axe","ion","six","ila","lao","mom","mas","pro","few","opt","poe","art","ash","oar","cap","lop","may","shy","rid","bat","sum","rim","fee","bmw","sky","maj","hue","thy","ava","rap","den","fla","auk","cox","ibo","hey","saw","vim","sec","ltd","you","its","tat","dew","eva","tog","ram","let","see","zit","maw","nix","ate","gig","rep","owe","ind","hog","eve","sam","zoo","any","dow","cod","bed","vet","ham","sis","hex","via","fir","nod","mao","aug","mum","hoe","bah","hal","keg","hew","zed","tow","gog","ass","dem","who","bet","gos","son","ear","spy","kit","boy","due","sen","oaf","mix","hep","fur","ada","bin","nil","mia","ewe","hit","fix","sad","rib","eye","hop","haw","wax","mid","tad","ken","wad","rye","pap","bog","gut","ito","woe","our","ado","sin","mad","ray","hon","roy","dip","hen","iva","lug","asp","hui","yak","bay","poi","yep","bun","try","lad","elm","nat","wyo","gym","dug","toe","dee","wig","sly","rip","geo","cog","pas","zen","odd","nan","lay","pod","fit","hem","joy","bum","rio","yon","dec","leg","put","sue","dim","pet","yaw","nub","bit","bur","sid","sun","oil","red","doc","moe","caw","eel","dix","cub","end","gem","off","yew","hug","pop","tub","sgt","lid","pun","ton","sol","din","yup","jab","pea","bug","gag","mil","jig","hub","low","did","tin","get","gte","sox","lei","mig","fig","lon","use","ban","flo","nov","jut","bag","mir","sty","lap","two","ins","con","ant","net","tux","ode","stu","mug","cad","nap","gun","fop","tot","sow","sal","sic","ted","wot","del","imp","cob","way","ann","tan","mci","job","wet","ism","err","him","all","pad","hah","hie","aim","ike","jed","ego","mac","baa","min","com","ill","was","cab","ago","ina","big","ilk","gal","tap","duh","ola","ran","lab","top","gob","hot","ora","tia","kip","han","met","hut","she","sac","fed","goo","tee","ell","not","act","gil","rut","ala","ape","rig","cid","god","duo","lin","aid","gel","awl","lag","elf","liz","ref","aha","fib","oho","tho","her","nor","ace","adz","fun","ned","coo","win","tao","coy","van","man","pit","guy","foe","hid","mai","sup","jay","hob","mow","jot","are","pol","arc","lax","aft","alb","len","air","pug","pox","vow","got","meg","zoe","amp","ale","bud","gee","pin","dun","pat","ten","mob"})

}
