package Week01
// 1. brute force rotate n times
// 2. reverse 3 times
// 3. follow up index rules check is n%k==0 start=start+1
// 4. copy to a new array

func rotate1(nums []int, k int)  {
	n:=len(nums)
	k=k%n
	if k==0{
		return
	}
	for i:=0;i<k;i++{
		last:=nums[n-1]
		for j:=n-2;j>=0;j--{
			nums[j+1]=nums[j]
		}
		nums[0]=last
	}
}

func rotate2(nums []int, k int)  {
	reverse:=func(nums []int,l,r int){
		for l<r{
			nums[l],nums[r]=nums[r],nums[l]
			l++
			r--
		}
	}
	n:=len(nums)
	k=k%n
	reverse(nums,0,n-1)
	reverse(nums,0,k-1)
	reverse(nums,k,n-1)

}

func rotate3(nums []int, k int)  {
	n:=len(nums)
	k=k%n
	cnt:=0
	for start:=0;cnt<n;start++{
		current:=start
		prev:=nums[start]
		for {
			nextIdx:=(current+k)%n
			nums[nextIdx],prev,current=prev,nums[nextIdx],nextIdx
			cnt++
			if current==start{
				break
			}
		}
	}

}



func rotate4(nums []int, k int)  {
	n:=len(nums)
	k=k%n
	if k==0{
		return
	}
	newNums:=make([]int,n,n)
	for i:=0;i<n;i++{
		newNums[(i+k)%n]=nums[i]
	}
	copy(nums,newNums)

}
