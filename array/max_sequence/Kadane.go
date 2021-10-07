package max_sequence

/**
Kadane 算法  是指  每次都以 数组 中 i 结尾的计算比较，从而可以求解 各种变式数组中最大连续子序列的方式
 */

/**
求数组中 最大连续子序列的和
 */
func maxSum(arr []int) int {
	dpi, res := arr[0], arr[0]
	for i := 0; i < len(arr); i ++ {
		dpi = max(arr[i], arr[i]+dpi)
		res = max(res, dpi)
	}
	return res
}

/**
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
示例 1:
输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
 */
func maxProduct(arr []int) int {
	dpimin, dpimax, res := arr[0], arr[0], arr[0]
	for i := 0; i < len(arr); i ++{
		dpimax, dpimin = max(arr[i], max(dpimax*arr[i], dpimin*arr[i])), min(arr[i], min(dpimin*arr[i], dpimax*arr[i]))
		res = max(res, dpimax)
	}
	return res
}

/*
给你一个整数数组，返回它的某个 非空 子数组（连续元素）在执行一次可选的删除操作后，所能得到的最大元素总和。
换句话说，你可以从原数组中选出一个子数组，并可以决定要不要从中删除一个元素（只能删一次哦），（删除后）子数组中至少应当有一个元素，然后该子数组（剩下）的元素总和是所有子数组之中最大的。
注意，删除一个元素后，子数组 不能为空。
输入：arr = [1,-2,0,3]
输出：4
解释：我们可以选出 [1, -2, 0, 3]，然后删掉 -2，这样得到 [1, 0, 3]，和最大。
 */
func maximumSum(arr []int) int {
	dpi, dpiDeled, res := arr[0], arr[0], arr[0]
	for i := 0; i < len(arr); i ++ {
		dpi, dpiDeled = max(arr[i], dpi+arr[i]), max(dpi, arr[i]+dpiDeled)
		res = max(res, max(dpi, dpiDeled))
	}
	return res
}

/**
给你一个整数数组 arr 和一个整数 k。
首先，我们要对该数组进行修改，即把原数组 arr 重复 k 次。
举个例子，如果 arr = [1, 2] 且 k = 3，那么修改后的数组就是 [1, 2, 1, 2, 1, 2]。
然后，请你返回修改后的数组中的最大的子数组之和。
注意，子数组长度可以是 0，在这种情况下它的总和也是 0。
由于 结果可能会很大，所以需要 模（mod） 10^9 + 7 后再返回。 
示例 1：
输入：arr = [1,2], k = 3
输出：9
 */
func kConcatenationMaxSum(arr []int, k int) int {
	dpi, res, singleSum := arr[0], arr[0], arr[0]
	for i := 1; i < len(arr)*min(k, 2); i ++ {
		dpi = max(arr[i%len(arr)], arr[i%len(arr)] + dpi)
		res = max(res , dpi)
		if i < len(arr){
			singleSum += arr[i]
		}
	}

	for singleSum > 0 && k > 2 {
		res = (res + singleSum) % 1000000007
		k --
	}
	return res %  1000000007

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
