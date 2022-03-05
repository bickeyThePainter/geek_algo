package brush

import "testing"

func reversePairs(nums []int) int {
	return reverse(nums,0,len(nums)-1)
}

func reverse(nums []int,l,r int) int{
	if l>=r{
		return 0
	}
	mid:=l+(r-l)>>1
	return reverse(nums,l,mid)+reverse(nums,mid+1,r)+merge(nums,l,r)
}

func merge(nums[]int,l,r int) int{
	if l>=r{
		return 0
	}
	mid:=l+(r-l)>>1
	res:=0
	tmp:=make([]int,r-l+1)
	index:=0
	i,j:=l,mid+1
	for i<=mid{
		for j<=r && float64(nums[i])/2.0>float64(nums[j]){
			j++
		}
		res+=j-mid-1
		i++
	}
	i,j=l,mid+1
	for i<=mid && j<=r{
		if nums[i]<=nums[j]{
			tmp[index]=nums[i]
			i++
			index++
		}else{
			tmp[index]=nums[j]
			j++
			index++
		}
	}
	for i<=mid{
		tmp[index]=nums[i]
		i++
		index++
	}

	for j<=r{
		tmp[index]=nums[j]
		j++
		index++
	}
	for i:=l;i<=r;i++{
		nums[i]=tmp[i-l]
	}
	return res

}

func TestReversePairs(t *testing.T){
	reversePairs([]int{1,3,2,3,1})
}
