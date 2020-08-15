package main
// subarray
// leetcode 974
// 暴力直接找, 接近O(n^2)
//func check(A []int, start int, K int, result *int) {
////	size := len(A)
////	var sum int
////	for i := start; i < size; i++ {
////		sum += A[i]
////		if sum%K == 0 {
////			result++
////		}
////	}
////}
//
//func subarraysDivByK1(A []int, K int) int {
//	var result int
//
//	for i := 0; i < len(A); i++ {
//		check(A, i, K, &result)
//	}
//	return result
//}

//通过规律, O(n)
// % is remainder operator in C++ (and not a proper modulus). To get a positive number, we have to add by the base: a mod b = ((a % b) + b) % b.
// Look for anytime S[j] % k == S[i-1] % k and you know everything in between [i...j] must be divisible by K!!
// 上面的公式无法适用于模为0的时候，模为0的时候，自身就是一个结果了，所以需要额外需要判断有s[j] % K == 0，但是为了统一算法，我们可以定下s[-1] = 0，这时是可以符合条件的
// https://leetcode.com/problems/subarray-sums-divisible-by-k/discuss/584722/C%2B%2B-O(N)-Explained
func subarraysDivByK(A []int, K int) int {
	remainMap := make(map[int]int, 0)
	var sum, remain, result int
	remainMap[0] = 1 //s[-1] = 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
		remain = ((sum % K) + K) % K
		if val, ok := remainMap[remain]; ok {
			result += val
		}
		remainMap[remain] = remainMap[remain] + 1
	}
	return result
}

// 在数组的处理中，一般都需要用于map来缩小运算时间
//func majorityElement(nums []int) int {
//	count := make(map[int]int, 0)
//	majority := len(nums) / 2
//	for _, val := range nums {
//		count[val] = count[val] + 1
//		if count[val] > n {
//			return v
//		}
//	}
//	return -1
//}

// 输出数组里面频繁前k的数字
// https://leetcode-cn.com/problems/top-k-frequent-elements/submissions/
// 这个的算法是利用了两个map，互为反向，数字到个数是用于查询当前数字的个数，个数到数字是为了输出答案
func topKFrequent(nums []int, k int) []int {
	val2Freq := make(map[int]int, 0)
	freq2Vals := make(map[int][]int,0)
	maxFreq := 0
	freq := 0
	for i := 0;i < len(nums); i++ {
		val2Freq[nums[i]] += 1
		freq = val2Freq[nums[i]]
		if freq2Vals[freq] == nil {
			freq2Vals[freq] = make([]int, 0)
		}
		freq2Vals[freq] = append(freq2Vals[freq], nums[i])
		if freq > maxFreq {
			maxFreq = freq
		}
	}

	result := make([]int, 0)
	last := 0
	for i := 0; i < k; {
		last = len(freq2Vals[maxFreq]) - 1
		if val2Freq[freq2Vals[maxFreq][last]] == maxFreq {
			result = append(result, freq2Vals[maxFreq][last])
			i += 1
		}
		freq2Vals[maxFreq] = freq2Vals[maxFreq][:last]
		for len(freq2Vals[maxFreq]) == 0 && maxFreq >= 0{
			maxFreq -= 1
		}
	}
	return result
}