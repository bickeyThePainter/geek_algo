package Week04

// 1. 动态规划
// 2. greedy 贪心

func jump1(nums []int) int {
	dp := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j < len(nums) && (dp[i+j] == 0 || dp[i+j] > dp[i]+1) {
				dp[i+j] = dp[i] + 1
			}
		}
	}
	return dp[len(nums)-1]
}

func jump2(nums []int) int {
	farthest, jumpEnd, jumps := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		n := nums[i]
		if i+n > farthest {
			farthest = i + n
		}
		if jumpEnd == i {
			jumps++
			jumpEnd = farthest
		}
	}
	return jumps

}

