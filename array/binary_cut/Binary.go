package binary_cut

/**
 二分法 指 在数组中 每次 都 从中间切分 ， 然后按照情况抛掉其中一半 ，以加速 查找的过程
 */


/*
已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,3,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,3]
若旋转 7 次，则可以得到 [0,1,3,4,5,6,7]
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
数组中无重复, 它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
 */

/**
查找  的关键是 寻找极小值  的过程
 */
func findMin(arr []int) int {
	s, b := 0, len(arr)-1
	for s < b {
		if arr[s] < arr[b] {
			return arr[s]
		}
		m := s + (b  - s) >> 1
		if arr[m] > arr[s] {
			s = m
		}else {
			b = m
		}
	}
	return arr[s]
}

/**
数组中有重复, 它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
 */
func findMinWithoutDump(arr []int) int {

}

/**
峰值元素是指其值严格大于左右相邻值的元素。

给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
你可以假设 nums[-1] = nums[n] = -∞ 。
你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
输入：nums = [1,2,1,3,5,6,4]
输出：1 或 5
解释：你的函数可以返回索引 1，其峰值元素为 2；
     或者返回索引 5， 其峰值元素为 6。
对于所有有效的 i 都有 nums[i] != nums[i + 1]

 */
func findPeakElement(nums []int) int {

}