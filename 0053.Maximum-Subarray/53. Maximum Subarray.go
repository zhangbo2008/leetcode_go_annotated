package leetcode

// 解法一 DP // 子数列和子序列问题都是按照dp[n]表示以n结尾的数组结果.做转移方程即可.
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		//if dp[i-1] > 0 {
		//	dp[i] = nums[i] + dp[i-1] //如果当前位置大于0,那么跟前面可以拼接更大.
		//} else {
		//	dp[i] = nums[i]  //如果当前小于0了.那么不使用前面即可.
		//}


		dp[i]=max(nums[i] + dp[i-1], nums[i] )
		res = max(res, dp[i])
	}
	return res
}

// 解法二 模拟
func maxSubArray1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, res, p := nums[0], 0, 0
	for p < len(nums) {
		res += nums[p]
		if res > maxSum {
			maxSum = res
		}
		if res < 0 {
			res = 0
		}
		p++
	}
	return maxSum
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}