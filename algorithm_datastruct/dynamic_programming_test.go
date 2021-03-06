package main

import (
	"fmt"
	"testing"
)

// nums 为列表，可正数，可负数，长度一定大于2
// 设置f(n)表示从下标0到下标n的最大不相邻数列之和，那么就可以有f(n) = max(f(n-1), f(n-2) + v(n))，其中v(n) = max(下标n的值，0)
func maxNotAdjacentNumsSum(nums []int) int {
	f := make([]int, len(nums))
	f[0] = max(0, nums[0])
	f[1] = max(f[0], nums[1])
	for i := 2; i < len(nums); i++ {
		v := max(0, nums[i])
		f[i] = max(f[i-1], f[i-2]+v)
	}
	return f[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// n表示游戏的数目，x表示总天数
// values表示通关游戏所获得的成就值，cost表示需要花费的天数
func zeroOnePackage(n int, x int, values []int, costs []int) int {
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, x+1)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= x; j++ {
			// 先初始化(假定不使用)，如果是只有一款就
			if i == 1 {
				f[i][j] = 0
			} else {
				f[i][j] = f[i-1][j]
			}
			// 比较使用的情况，要满足剩余的天数大于消耗的天数
			if j >= costs[i] && f[i][j] < f[i-1][j-costs[i]]+values[i] {
				f[i][j] = f[i-1][j-costs[i]] + values[i]
			}
		}
	}
	return f[n][x]
}

func TestZeroOnePackage(t *testing.T) {
	//nums := []int{7,8,-1,5,6,10,-3}
	//fmt.Println(maxNotAdjacentNumsSum(nums))

	n := 2
	x := 2
	values := []int{0, 10, 20}
	costs := []int{0, 1, 2}
	fmt.Println(zeroOnePackage(n, x, values, costs))
}

// leetcode 198
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2] + nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

// leetcode 1043
// dp(i) = max(dp(i), dp(j) + (i-j)*MAX(v[j...i]))，其中0<i-j<=k,也就是j是变化的，要迭代
func maxSumAfterPartitioning(arr []int, k int) int {
	dp := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		j := i - 1
		maxVal := 0
		for i - j <= k && j >= 0 {
			maxVal = max(maxVal, arr[j])
			dp[i] = max(dp[i], dp[j] + (i-j)*maxVal)
			j--
		}
	}

	return dp[len(arr)-1]
}

func zeroOnePackage1(n int, x int, values []int, costs []int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, x+1)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= x; j++ {
			// 初始化，假定不使用
			if i == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j]
			}
			// 进行比较
			if costs[i] <= j && dp[i][j] > dp[i-1][j-costs[i]] + values[i] {
				dp[i][j] =  dp[i-1][j-costs[i]] + values[i]
			}
		}
	}
	return dp[n][x]
}

func longestPalindrome1(s string) string {
	size := len(s)
	d := make([][]bool, size)
	for i := 0; i < size; i++ {
		d[i] = make([]bool, size)
	}

	// 针对长度做迭代
	output := ""
	curSize := 0
	for l := 1; l <= size; l++ {
		for i := 0; i <= size - l; i++ {
			j := i + l - 1
			if l == 1 {
				d[i][j] = true
			} else if l == 2 {
				d[i][j] = s[i] == s[j]
			} else {
				d[i][j] = d[i+1][j-1] && s[i] == s[j]
			}
			if d[i][j] == true && curSize < l {
				output = s[i:j+1]
			}
		}
	}
	return output
}