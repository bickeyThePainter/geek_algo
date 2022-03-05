package brush

import "testing"

func decodeString(s string) string {
	res:=""
	for i:=0;i<len(s);{
		if s[i]>='0' && s[i]<='9'{
			cur,curIndex:=decode(s,i,"")
			res+=cur
			i=curIndex+1
		}else{
			res+=s[i:i+1]
		}
	}
	return res
}
func decode(s string, index int,prev string) (string,int){
	if s[index]==']'{
		return prev,index
	}
	if s[index]>='a' && s[index]<='z'{
		return decode(s,index+1,prev+s[index:index+1])
	}
	if s[index]=='['{
		return decode(s,index+1,prev)
	}
	left:=0
	for index<len(s) && s[index]>='0' && s[index]<='9'{
		left*=10
		left+=int(s[index]-'0')
		index++
	}
	cur,curIndex:=decode(s,index,"")
	res:=""
	for i:=0;i<left;i++{
		res+=cur
	}
	return prev+res,curIndex
}

func TestDecode(t *testing.T){
	decodeString("2[abc]3[cd]")
}
