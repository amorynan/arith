package array

import (
	"fmt"
	"sort"
)

// ======================== recursive ===============================
func uniquePathsWithRecursion(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	lastRow, lastCol = m-1, n-1
	res := 0
	seekTarget(0, 0, &res)
	return res
}

func seekTarget(row, col int, res *int) {
	// find finish
	if row == lastRow || col == lastCol {
		*res ++
		return
	}

	if col <= lastCol && row <= lastRow {
		seekTarget(row, col+1, res)
		seekTarget(row+1, col, res)
	}

}

func uniquePaths(m int, n int) int {
	// recursion method
	//return uniquePathsWithRecursion(m int, n int)


	// dp function 1

	//return uniquePathsWithDpOne(m, n)
	// m is small
	if m > n {
		m, n = n, m
	}

	dp := make([]int, m, m)


	for b := 0; b < n; b ++ {
		for s := 0; s < m; s ++ {
			if b == 0 || s == 0 {
				dp[s] = 1
				continue
			}

			dp[s] = dp[s] + dp[s-1]
		}
	}
	return dp[m-1]
}


func uniquePathsWithDpOne(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for row := 0; row < m;  row++{
		for col := 0; col < n; col++ {
			if row == 0 || col == 0{
				dp[row][col] = 1
				continue
			}

			dp[row][col] = dp[row-1][col] + dp[row][col-1]
		}
	}
	return dp[m-1][n-1]
}



//========================== 加和 =======================================
func twoSum(nums []int, target int) []int {
	// only go through once, because first just can sum with the left in last
	leftMap := make(map[int]int, len(nums))
	for idx, item := range nums {
		if val, exist := leftMap[target - nums[idx]]; exist {
			return []int{idx, val}
		}
		leftMap[item] = idx
	}
	return nil
}

/**
Given array nums = [-1, 0, 1, 2, -1, -4],
A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
 */
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	if len(nums) < 3 {
		return nil
	}
	res := make([][]int, 0, 0)
	alreadyRight := make(map[int]map[int]struct{})
	leftMap := map[int][][]int{nums[0]+nums[1]:{{nums[0], nums[1]}}}
	for i := 2; i < len(nums); i++{
		f1, already := alreadyRight[nums[i]]
		if v, exist := leftMap[-nums[i]]; exist {
			for _, f := range v{
				if  _, e := f1[f[0]]; e && already {
					continue
				}
				res = append(res, []int{f[0],f[1], nums[i]})
				if f1 == nil {
					f1 = make(map[int]struct{})
				}
				f1[f[0]] = struct{}{}
				alreadyRight[nums[i]] = f1
			}
		}
		j:=0
		loop:
			for; j<i; j++ {
				if v, exist := leftMap[nums[i]+nums[j]];exist {
					for _, f := range v{
						if f[0] == nums[j] {
							j++
							goto loop
						}
					}
				}
				leftMap[nums[i]+nums[j]] = append(leftMap[nums[i]+nums[j]], []int{nums[j], nums[i]})
			}
		}
	return res
}
// sort and three pointer
// 排序对于去重 是很重要的手段
// 主要需要考虑的细节：1。去重判断，减少重复的计算，确定每个index 为三元组最小值之后，确定第二个最小的和最大那个值的即可，
//去重的场景：最小的值重复了可以直接跳过；在最小值确定的情况下， 第二小的也如果重复已经计算过的，也可以直接跳过
func threeSum_Best(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)

	res, index := make([][]int, 0, 0), 0

	// index 止步于倒数第三位
	for ; index < len(nums)-2; index ++ {
		start, end := index+1, len(nums)-1
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}

		for start < end {
			if start > index+1 && nums[start] == nums[start-1] {
				start ++
				continue
			}

			if end < len(nums)-1 && nums[end] == nums[end+1] {
				end --
				continue
			}
			sum := nums[index] + nums[start] + nums[end]
			if sum == 0 {
				res = append(res , []int{nums[index], nums[start], nums[end]})
				start ++
				end --
				continue
			}

			if sum < 0 {
				start ++
				continue
			}

			if sum > 0 {
				end --
				continue
			}

		}
	}
	return res
}

/**
Given array nums = [-1, 2, 1, -4], and target = 1.
The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 */
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	res, minDiff, index := 0, 0, 0
	for ; index < len(nums)-2; index ++ {
		if index > 1 && nums[index] == nums[index-1] {
			continue
		}
		start, end := index + 1, len(nums) - 1

		for start < end {
			if start > index + 1 && nums[start] == nums[start-1] {
				start ++
				continue
			}
			if end < len(nums)-1 && nums[end] == nums[end+1] {
				end --
				continue
			}
			sum := nums[index] + nums[start] + nums[end]
			diff := sum - target
			if diff == 0 {
				return target
			}
			if diff > 0 {
				diff = -diff
			}
			if minDiff == 0 || diff > minDiff {
				res = sum
				minDiff = diff
			}
			if sum < target {
				start ++
				continue
			}
			if sum > target {
				end --
				continue
			}
		}
	}
	return res
}

/**
 Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
 */
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	sort.Ints(nums)

	res, idx1 := make([][]int, 0), 0

	for ; idx1 < len(nums)-3; idx1 ++ {
		if idx1 > 0 && nums[idx1] == nums[idx1-1] {
			continue
		}
		for idx2 := idx1+1; idx2 < len(nums)-2 ; idx2 ++ {
			if idx2 > idx1+1 && nums[idx2] == nums[idx2-1] {
				continue
			}
			start , end := idx2 + 1, len(nums) -1
			for start < end {
				if start > idx2+1 && nums[start] == nums[start-1] {
					start ++
					continue
				}
				if end < len(nums)-1 && nums[end] == nums[end+1] {
					end --
					continue
				}

				sum := nums[idx1] +  nums[idx2] + nums[start] + nums[end]
				if sum == target {
					res = append(res, []int{nums[idx1] ,nums[idx2], nums[start], nums[end]})
					start ++
					end --
					continue
				}
				if sum < target {
					start ++
					continue
				}
				end --
			}
		}
	}
	return res

}

/**
 限定总和，求所有的加和方式
Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]
 */
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	resItem, res := make([]int, 0), make([][]int, 0)
	combinationSumLoop(candidates, resItem, 0, target, &res)
	return res
}

/**
$S 组合排序 = 回溯+记录
有几个需要注意的地方：1.注意每一个解都需要重新copy到res 中 2. 注意从每一个解在回来的时候需要重重置
 */
func combinationSumLoop(candidates, resItem []int, idx, target int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		// resItem sums is target , so make it into res
		// #### here is import copy , 如果直接将resItem 放进res 里面的的话，后面如果在回溯的过程中，有新的试错元素加入当前的resItem ，就会被覆盖
		cp := make([]int, len(resItem))
		copy(cp, resItem)
		*res = append(*res, cp)
		return
	}

	for i := idx; i < len(candidates); i ++ {
		if candidates[i] > target {
			// no need to combination
			break
		}
		resItem = append(resItem, candidates[i])
		combinationSumLoop(candidates, resItem, i, target-candidates[i], res)
		// 需要去掉这一次的经历
		resItem = resItem[:len(resItem)-1]
	}
}

/**
组合求和， 要求不能有重复的集合和元素
注意和上面不一样的地方除了下一个解的开始地方，以及消重的地方
 */
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	if len(candidates) == 0 {
		return [][]int{}
	}
	resItem, res := []int{}, [][]int{}
	combinationSum2Loop(candidates, resItem, 0, target, &res)
	return res
}

func combinationSum2Loop(candidates, resItem []int, idx, target int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		cp := make([]int, len(resItem))
		copy(cp, resItem)
		*res = append(*res, cp)
		return
	}

	for i := idx; i<len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		if  i > idx && candidates[i] == candidates[i-1]{
			continue
		}
		resItem = append(resItem, candidates[i])
		combinationSum2Loop(candidates, resItem, i+1, target-candidates[i], res)
		resItem = resItem[:len(resItem)-1]
	}
}

/**
 给定一个乘积 ，找到所有积因子组合
 */
func combinationMultiplication(multiplication int) [][]int {
	if multiplication == 0 {
		return [][]int{}
	}

	resItem, res, candidates := []int{}, [][]int{}, []int{}
	for i := 2; i < multiplication ; i ++ {
		candidates = append(candidates, i)
	}
	combinationMultiplicationLoop(candidates, resItem, multiplication, 0, &res)

	for _, v := range res {
		v1 := append(v,  1)
		res = append(res, v1)
	}
	res = append(res, []int{1, multiplication})
	return res
}

func combinationMultiplicationLoop(candidates, resItem []int, target, idx int, res *[][]int){
	if target == 1 {
		cp := make([]int, len(resItem))
		copy(cp, resItem)
		*res = append(*res, cp)
		return
	}

	// i 从0 开始 会让 结果集出现重复 和 不从0开始，让结果集不重复
	for i := idx; i < len(candidates); i++ {
		// 这里和上面的剪枝部分还不太一样
		if target % candidates[i] != 0 {
			continue
		}
		resItem = append(resItem, candidates[i])
		combinationMultiplicationLoop(candidates, resItem, target / candidates[i], i, res)
		resItem = resItem[:len(resItem)-1]
	}
}

/**
Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
 */
func removeDuplicates(nums []int) (int, []int) {
	i := 0
	for i< len(nums) {
		rep := i+1
		for rep < len(nums) && nums[rep] == nums[i] {
			nums = append(nums[:rep], nums[rep+1:]...)
		}
		i = rep
	}
	return len(nums), nums
}

/**
 采用复制解法的好处
 1. 不用在每次比较的时候改变原数组，思路清晰，
 2. 同时，降低每次改变元素都会有的性能开销
 */
func removeDuplicates_Best(nums []int) (int, []int) {
	if len(nums) <= 1{
		return len(nums), nums
	}
	i,r := 0,1
	for r < len(nums) {
		if nums[i] == nums[r] {
			r ++
			continue
		}
		if r - i > 1 {
			nums[i+1] = nums[r]
		}
		i++
		r++
	}
	nums = nums[:i+1]

	return len(nums), nums
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mid := (len(nums1) + len(nums2)) / 2
	midLeft := -1
	// odd ?
	if (len(nums1) + len(nums2)) % 2 == 0 {
		midLeft = mid-1
	}
	totalNum := make([]int, 0, len(nums1)+len(nums2))
	idx1, idx2 := 0, 0
	for idx:=0; idx< cap(totalNum); idx++ {
		if idx1 >= len(nums1) {
			totalNum = append(totalNum, nums2[idx2:]...)
			break
		}

		if idx2 >= len(nums2) {
			totalNum = append(totalNum, nums1[idx1:]...)
			break
		}

		if nums1[idx1] < nums2[idx2]{
			totalNum = append(totalNum, nums1[idx1])
			idx1 ++
		}else {
			totalNum = append(totalNum, nums2[idx2])
			idx2 ++
		}
	}

	if midLeft < 0 {
		return float64(totalNum[mid])
	}else {
		return float64(totalNum[midLeft] + totalNum[mid]) / 2.0
	}
}

/**
 依旧采用复制+双指针的方式，但是要格外注意一下 结束的条件，要让最后的i的值也需要比较，然后就是 空值需要 直接返回
 */
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return len(nums)
	}
	i, r := 0, len(nums)-1
	for i != r+1 {
		if nums[r] == val {
			r --
			continue
		}
		if nums[i] == val {
			nums[i] = nums[r]
			r --
		}

		i++
	}

	nums = nums[:i]
	return len(nums)
}

/**
 #s 找到当前排列int 数组的下一个大的数组
 */
func nextPermutation(nums []int)  []int{
	if len(nums) <= 1 {
		return nums
	}

	f, l := len(nums)-2, len(nums)-1

	for f>=0 && l < len(nums) && nums[f] >= nums[l] {
		for l < len(nums) {
			if nums[f] >= nums[l] {
				l ++
			}else {
				nums = switchArr(nums, f, l)
				return nums
			}
		}
		i, j := f, f+1
		for ; j < len(nums); j ++  {
			nums = switchArr(nums, i ,j)
			i ++
		}
		f--
		l = f+1
	}
	if f >= 0 {
		nums = switchArr(nums, f, l)
	}
	return nums
}

func switchArr(nums []int, idx1, idx2 int) []int{
	temp := nums[idx1]
	nums[idx1] = nums[idx2]
	nums[idx2] = temp
	return nums
}

/**
 #S. 旋转排序数组找到目标值，时间复杂度：o(logn)
还是二分，最重要的思想上判断有序的部分在那部分
 */
func search(nums []int, target int) int {
	if target < nums[0] && target > nums[len(nums)-1] {
		return -1
	}
	if len(nums) == 1 && nums[0] == target{
		return 0
	}else if len(nums) == 1 && nums[0] != target {
		return -1
	}

	l, r := 0, len(nums)-1
	//detail 1. l <= r to make sure mid is (l == r)
	for l <= r {
		m := (l + r) >> 1
		if nums[m] == target {
			return m
		}

		if nums[l] <= nums[m] {
			// order in left part
			if target >= nums[l] && target < nums[m] {
				r = m - 1
			}else {
			  	l = m + 1
			}
		}else {
			// order in right part
			// detail 2. there should be m+1 ,so use >= ,not >
			if target <= nums[r] && target >= nums[m+1] {
				l = m + 1
			}else {
				r = m - 1
			}
		}

	}
	return -1
}

//================================= 数组hash ========================
/**
 找到数组中缺失的第一个正整数
 	主要掌握两个解决方案
 */
func swap(nums *[]int, i, j int) {
	temp := (*nums)[i]
	(*nums)[i] = (*nums)[j]
	(*nums)[j] = temp
}
func firstMissingPositive(nums []int) int {
	// 置换 = 将正序的数据都先摆正到他的位置上
	//return firstMissingPositiveSwap(nums)
	return firstMissPositiveIndex(nums)
}

func firstMissPositiveIndex(nums []int) int {
	// trick 的方式: 由于符合问题的取值范围只能是[1, len(nums)+1], 而数组的下标取值范围是[0, len(nums)-1], 如果将后面的数据关联前面的，那就是数组
	// 下标方式作为一个简单的hash
	// 那么问题就在于如何关联------------- 如果原数组中的数字在[1, len(nums)] 中，那么就在该数字-1 作为数组的index 就可以标记一下，代表这个数字存在，
	// 标记的方式 就很 trick 了，把当前的数据取负一下即可， 那么如果本来就是负数的，就需要区分是不是我们手动变过来的，还是本来的，因为如果本来就是负数，也
	// 理解成我们变过来的，就相当于，改负数index 的值也被存在我们可以检索存不存在的hash map 中了，但是实际人家是不存在的, 所以我们预先处理一下，将原本的
	// 负数 都变成0
	// pre deal : 将0 和 负数 都当作缺失的值，变成len(nums)+1 这个最大的值
	for i, v := range nums {
		if v < 0 {
			nums[i] = len(nums)+1
		}
	}
	// 继续处理接下来的
	for _, v := range nums {
		// 现在还是负数的一定是手动变来的
		if v < 0 {
			v = -v
		}

		if v > 0 && v <= len(nums) && nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}
	}

	min := 1
	for idx, v := range nums{
		if v == 0 {
			continue
		}
		if v < 0 && min == idx+1{
			min ++
		}
	}
	return min
}

func firstMissingPositiveSwap(nums []int) int {
	// 置换 = 将正序的数据都先摆正到他的位置上
	for i := 0; i < len(nums) ; i ++ {
		if nums[i] < 1 {
			continue
		}
		//***** 这里要格外注意，并且，有大于 小于符号的 放在前面，顺序也不能错
		for nums[i] >= 1 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1]  {
			swap(&nums, i, nums[i]-1)
		}
	}
	min := 1
	for _, v := range nums{
		if v < 1 {
			continue
		}
		if v == min {
			min ++
		}
	}
	return min
}

/**
执行效果在go
 */
func maxArea(height []int) int {
	maxRecord := 0

	head, tail := 0, len(height)-1
	for head < tail {
		maxRecord = max(maxRecord, (tail-head) * min(height[head], height[tail]))
		if (height[head] <= height[tail]) {
			head ++
		} else {
			tail --
		}
	}
	return maxRecord
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

/**
 #####sssss 经典的接🌧️雨水 题目

 */
func trap(height []int) int {

	// find a cal area decrease and up
	l, area := 0, 0
	for l < len(height)-1 {
		if l < len(height)-2 && height[l] == height[l+1] {
			l ++
			continue
		}
		t := l + 1
		// find decrease pointer
		if height[t] >= height[l] {
			l ++
			continue
		}
		decrease := 0
		// find up pointer
		for t < len(height) && height[t] < height[l]{
			if height[t] <= height[t-1] {
				decrease ++
			}
			t ++
		}
		if t == len(height) && decrease == t-l-1 {
			// there is no need to cal,always decrease
			break
		}
		min := height[l]
		if t == len(height){
			// from right -> left
			break
		}
		// cal the area
		mid := l + 1
		for mid < t {
			area += min-height[mid]
			mid ++
		}
		l = t
	}
	r := len(height)-1
	for r > l {
		if r > 0 && height[r] == height[r-1] {
			r --
			continue
		}
		t := r - 1
			if height[t] >= height[t+1] {
				r --
				continue
			}
			for t > l && height[t] < height[r] {
				t --
			}
			min := height[r]
			mid := r - 1
			for mid > t {
				area += min - height[mid]
				mid --
			}
			r = t
	}
	return area
}

func trap_best(height []int) int {
	l, r , firstMaxLeft, firstMaxRight, res := 0, len(height)-1, 0, 0, 0
	for l < r {
		if height[l] <= height[r] {
			// 左边比右边矮，所以左边是瓶颈，作为可以接雨水的标准
			if height[l] >= firstMaxLeft {
				firstMaxLeft = height[l]
			}else {
				res += firstMaxLeft - height[l]
			}
			l ++
		}else {
			// 同右边
			if height[r] >= firstMaxRight {
				firstMaxRight = height[r]
			} else {
				res += firstMaxRight - height[r]
			}
			r --
		}
	}
	return res
}

/*
 数组中的数字代表是可以jump的步数， 从 index 为0 开始 jump 到 最后一个位置，中间jump 的次数最少是多少
[2,3,1,1,4] => 2->3->4 => res:2
⚠️难点主要是细节注意
 */
func jump(nums []int) int {
	jumps , sIdx := 0, 0
	for sIdx < len(nums) {
		if sIdx >= len(nums)-1 {
			break
		}
		maxNextIdx := 0
		nIdx := 0
		for n:=sIdx+1; n <= sIdx+nums[sIdx] && n <= len(nums)-1 ; n++{
			if n == len(nums)-1 {
				// already jump here
				nIdx = n
				break
			}
			if nums[n]+n >= maxNextIdx{
				maxNextIdx = nums[n]+n
				nIdx = n
			}
		}

		sIdx = nIdx
		jumps ++
	}
	return jumps
}

/**
可以跳跃到最后一步的判断
和上面的进行对比，由上面的方式可以做出来，只是判断最后不能跳动到最后的判断条件需要大道至简一下
其次是还有一种可以看起来更简洁的做法
 */
func canJump(nums []int) bool {
	curIdx := 0
	for curIdx < len(nums) && curIdx + nums[curIdx] < len(nums){
		if nums[curIdx] == 0 {
			return false
		}
		maxIdx , maxNextIdx := 0, 0
		for i:=curIdx; i <= curIdx+nums[curIdx] && i < len(nums); i++{
			tmpNextIdx := i + nums[i]
			if tmpNextIdx >= len(nums)-1 {
				return true
			}
			if tmpNextIdx >= maxNextIdx {
				maxIdx = i
				maxNextIdx = tmpNextIdx
			}
		}
		curIdx = maxIdx
	}
	return true
}

/**
 只需要维护一个最远的Index，即每一个nums 的数字都加上当前的index，为可出现的所有的最远的Index
 如果发现遍历到的某一个idx 下，当前的idx 已经大于 最远能跳到的Index，那就说明已经跳不到最远了
 */
func canJump_simple(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}

	maxJumpIdx := 0
	for i, v := range nums {
		if maxJumpIdx == len(nums)-1 {
			return true
		}
		if i > maxJumpIdx {
			return false
		}
		if maxJumpIdx < i+v {
			maxJumpIdx = i+v
		}
	}
	return true
}

// ===================================== 矩阵 matrix  ====================================
/**
 🚩旋转矩阵, 还是采用矩阵的对称性
 */
func rotate(matrix [][]int) {
	for i :=0 ;i < len(matrix) ; i ++ {
		for j := i + 1; j < len(matrix); j ++ {
			(matrix)[i][j], (matrix)[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < len(matrix) ; i ++ {
		for j := 0; j < len(matrix) / 2 ; j ++ {
			matrix[i][j], matrix[i][len(matrix)-j-1] = matrix[i][len(matrix)-j-1], matrix[i][j]
		}
	}
}

/**
 ⚠️一次通过还需要注意很多细节
	 小技巧就是判断完成的条件是确定的
 */
func spiralOrder(matrix [][]int) []int {
	edgeRowFrom, edgeRowTo, edgeColLeft, edgeColRight, row, col, res := 0, len(matrix), 0, len(matrix[0]), 0, 0, make([]int, 0, len(matrix)*len(matrix[0]))

	for cap(res) != len(res) {
		// col from left to right
		for cap(res) != len(res) && col < edgeColRight {
			res = append(res, matrix[row][col])
			col ++
		}
		col --
		// row from top to down
		for cap(res) != len(res) && row < edgeRowTo-1 {
			row ++
			res = append(res, matrix[row][col])
		}

		// col from right to left
		for cap(res) != len(res) && col > edgeColLeft {
			col --
			res = append(res, matrix[row][col])
		}
		// row  from down to up
		row --
		for cap(res) != len(res) && row > edgeRowFrom {
			res = append(res, matrix[row][col])
			row --
		}

		edgeRowFrom++
		edgeRowTo--
		edgeColLeft++
		edgeColRight--
		row++
		col++
	}
	return res
}

func generateMatrix(n int) [][]int {

	// go through the matrix and put val to it
	row, col, rowEdgeMin, rowEdgeMax, colEdgeMin, colEdgeMax :=0,0, 0, n, 0, n
	res := make([][]int, n, n)
	for i,_ := range res {
		res[i] = make([]int, n, n)
	}
	cnt := 1
	lastCnt := n*n

	for row <= n / 2 {
		// from left to right
		for cnt < lastCnt && col < colEdgeMax-1 {
			res[row][col] = cnt
			cnt++
			col ++
		}
		// from top to down
		for cnt < lastCnt && row < rowEdgeMax-1 {
			res[row][col] = cnt
			cnt++
			row ++
		}
		// from right to left
		for cnt < lastCnt && col > colEdgeMin  {
			res[row][col] = cnt
			cnt++
			col --
		}
		// from down to top
		for cnt <= lastCnt && row > colEdgeMin {
			res[row][col] = cnt
			cnt++
			row --
		}

		row++
		col++
		rowEdgeMin++
		rowEdgeMax--
		colEdgeMin++
		colEdgeMax--
	}
	if n % 2 != 0 {
		res[n/2][n/2] = cnt
	}
	return res
}

/**
 合并区间
 给定一些区间，尽可能合并一些交叉的区间，最后得到所有区间
	最重要的其实是前面区间的排序过程, quick sort or other may choose
 */
func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	// sort the matrix
	m := make(map[int][]int, 0)
	keys := make([]int, 0)
	for _, item := range intervals {
		if _, exist := m[item[0]]; !exist {
			m[item[0]] = make([]int, 0)
			keys = append(keys, item[0])
		}
		m[item[0]] = append(m[item[0]], item[1])
	}
	sort.Ints(keys)

	intervals = make([][]int, 0)
	for _, k := range keys {
		for _, v := range m[k] {
			intervals = append(intervals, []int{k, v})
		}
	}

	// merge the intervals
	res := make([][]int, 0)
	for _, item := range intervals {
		if len(res) == 0 {
			res = append(res, item)
			continue
		}
		lastItem := res[len(res)-1]
		if item[0] <= lastItem[1] {
			// merge
			if lastItem[1] < item[1] {
				res[len(res)-1][1] = item[1]
			}
			continue
		}
		res = append(res, item)
	}
	return res
}

var lastRow, lastCol = 0, 0


/**
set zero
需要空间复杂度是常量级别的
 */
func setZeroes(matrix [][]int)  {
	rowsToBeSet := make([]int,0,len(matrix))
	colsToBeSet := make([]int,0, len(matrix[0]))

	for r := 0; r < len(matrix); r++{
		for c := 0; c < len(matrix[0]); c ++ {
			if matrix[r][c] == 0 {
				if len(colsToBeSet) == 0 || (len(colsToBeSet)>0 && colsToBeSet[len(colsToBeSet)-1] != c) {
					colsToBeSet = append(colsToBeSet, c)
				}
				if  len(rowsToBeSet) == 0 || len(rowsToBeSet)>0 && rowsToBeSet[len(rowsToBeSet)-1] != r {
					rowsToBeSet = append(rowsToBeSet, r)
				}
			}
		}
	}

	for _, r := range rowsToBeSet {
		for c := 0; c < len(matrix[0]); c ++ {
			*&matrix[r][c] = 0
		}
	}

	for _, c := range colsToBeSet {
		for r := 0; r < len(matrix); r ++ {
			*&matrix[r][c] = 0
		}
	}
}
/**
	在矩阵中查找值
重要点是 如何 映射指针下标到 数组下标
 */
func searchMatrixWithTwiceBinarySearch(matrix [][]int, target int) bool {
	sr, er, hr := 0, len(matrix)-1, 0
	for sr <= er {
		hr = (sr + er) >> 1
		if matrix[hr][len(matrix[0])-1] == target {
			return true
		}
		if matrix[hr][len(matrix[0])-1] < target {
			sr = hr+1
		}else {
			er = hr-1
		}
	}
	if sr > len(matrix)-1 {
		return false
	}

	sc, ec, hc := 0, len(matrix[0])-1, 0
	for sc <= ec {
		hc = (sc + ec ) >> 1
		if matrix[sr][hc] == target{
			return true
		}
		if matrix[sr][hc] < target {
			sc = hc+1
		}else {
			ec = hc-1
		}
	}
	return false
}

/**
 只需要一次二分查找
 */
func searchMatrix(matrix [][]int, target int) bool {
	// search with twice binary search
	//return searchMatrixWithTwiceBinarySearch(matrix, target)
	s, e, m := 0, len(matrix) * len(matrix[0]) - 1, 0

	for s <= e {
		m = (s + e) >> 1

		if matrix[m/len(matrix[0])][m-m/len(matrix[0])*len(matrix[0])] == target {
			return true
		}
		if matrix[m/len(matrix[0])][m%len(matrix[0])] < target {
			s = m + 1
		}else {
			e = m-1
		}
	}
	return false
}

/*
 数独判断，
 重点是如何在一次循环中判断完成所有不重复的0-9
 特别是在矩阵中辨识 数独中的 格子的
 */
func isValidSudoku(board [][]byte) bool {
	// 暴力破解
	//isValidSudokuRough(board)
	// for every col , row and box add a 9 index * 9
	rowCache, colCache, boxCache := make([][]bool, 9) , make([][]bool, 9), make([][]bool, 9)
	for i, _ := range rowCache {
		rowCache[i] = make([]bool, 9)
	}
	for i, _ := range colCache {
		colCache[i] = make([]bool, 9)
	}
	for i, _ := range boxCache {
		boxCache[i] = make([]bool, 9)
	}
	for r:=0; r < 9 ; r ++ {
		for c:=0; c < 9; c++{
			if board[r][c] == '.' {
				continue
			}
			num := board[r][c] - '0' - byte(1)
			if rowCache[r][num] || colCache[c][num] || boxCache[r/3*3+c/3][num] {
				return false
			}
			rowCache[r][num] = true
			colCache[c][num] = true
			boxCache[r/3*3+c/3][num] = true
		}
	}
	return true
}
func isValidSudokuRough(board [][]byte) bool {

	checkMap := make(map[byte]struct{}, 0)
	for r:=0; r<9; r++{
		for c:=0; c<9; c++ {
			if _, exist := checkMap[board[r][c]]; exist {
				return false
			}
			if board[r][c] != '.' {
				checkMap[board[r][c]] = struct{}{}
			}
		}
		checkMap = make(map[byte]struct{}, 0)
	}
	checkMap = make(map[byte]struct{}, 0)
	for c:=0; c<9; c++{
		for r:=0; r<9; r++ {
			if _, exist := checkMap[board[r][c]]; exist {
				return false
			}
			if board[r][c] != '.' {
				checkMap[board[r][c]] = struct{}{}
			}
		}
		checkMap = make(map[byte]struct{}, 0)
	}


	for r:=0; r <= 6; r = r + 3 {
		for c:=0; c <= 6; c = c+3{
			checkMap = make(map[byte]struct{}, 0)
			for ri := r ; ri < r+3; ri++ {
				for ci := c ; ci < c+3; ci++ {
					if _, exist := checkMap[board[ri][ci]]; exist {
						return false
					}
					if board[ri][ci] != '.' {
						checkMap[board[ri][ci]] = struct{}{}
					}
				}
			}
		}
	}

	return true
}

/**
 🏁🏁 解数独
	还是需要注意dfs 中 定义的辅助数据的变化, 以及归来的过程这些辅助数据的还原
 */
type EmptyPostion struct {
	row int
	col int
}
func solveSudoku(board [][]byte)  {
	rowCache, colCache, boxCache, emptyPos := make([][]bool, 9), make([][]bool,9), make([][]bool, 9), make([]EmptyPostion, 0)
	for i := range rowCache {
		rowCache[i] = make([]bool, 9)
	}
	for i := range rowCache {
		colCache[i] = make([]bool, 9)
	}
	for i := range rowCache {
		boxCache[i] = make([]bool, 9)
	}
	// first cache all dimension that value already show up
	for r:=0; r<9; r++ {
		for c:=0; c<9; c++{
			if board[r][c] == '.' {
				emptyPos = append(emptyPos, EmptyPostion{row: r, col: c})
				continue
			}
			num := board[r][c] - '0' - byte(1)
			rowCache[r][num] = true
			colCache[c][num] = true
			boxCache[r/3*3+c/3][num] = true
		}
	}

	// recursive for put value
	ok := putValueRecursive(emptyPos, 0 ,&board, &rowCache, &colCache, &boxCache)
	//var putValueRecursive func(int) bool
	//putValueRecursive = func(idx int) bool {
	//	if idx == len(emptyPos) {
	//		// end
	//		return true
	//	}
	//
	//
	//	for i:=0 ; i < 9 ; i++ {
	//		if rowCache[emptyPos[idx].row][i] || colCache[emptyPos[idx].col][i] || boxCache[emptyPos[idx].row/3*3+emptyPos[idx].col/3][i] {
	//			// exist
	//			continue
	//		}
	//		board[emptyPos[idx].row][emptyPos[idx].col] = byte(i+1)+'0'
	//		rowCache[emptyPos[idx].row][i] = true
	//		colCache[emptyPos[idx].col][i] = true
	//		boxCache[emptyPos[idx].row/3*3+emptyPos[idx].col/3][i] = true
	//
	//		ok := putValueRecursive(idx+1)
	//		if ok {
	//			return true
	//		}
	//		// continue with next possible value
	//		board[emptyPos[idx].row][emptyPos[idx].col] = '.'
	//		rowCache[emptyPos[idx].row][i] = false
	//		colCache[emptyPos[idx].col][i] = false
	//		boxCache[emptyPos[idx].row/3*3+emptyPos[idx].col/3][i] = false
	//	}
	//	return false
	//}
	//ok := putValueRecursive(0)
	fmt.Println(ok)
}
//
func putValueRecursive(emptyPos []EmptyPostion, idx int,board *[][]byte, rowCache, colCache, boxCache *[][]bool) bool {
	if idx == len(emptyPos) {
		// end
		return true
	}


	for i:=0 ; i < 9 ; i++ {
		if (*rowCache)[emptyPos[idx].row][i] || (*colCache)[emptyPos[idx].col][i] || (*boxCache)[emptyPos[idx].row/3*3+emptyPos[idx].col/3][i] {
			// exist
			continue
		}
		(*board)[emptyPos[idx].row][emptyPos[idx].col] = byte(i+1)+'0'
		(*rowCache)[emptyPos[idx].row][i] = true
		(*colCache)[emptyPos[idx].col][i] = true
		(*boxCache)[emptyPos[idx].row/3*3+emptyPos[idx].col/3][i] = true

		ok := putValueRecursive(emptyPos, idx+1, board, rowCache, colCache, boxCache)
		if ok {
			// already find , and do not continue
			return true
		}
		// continue with next possible value
		(*board)[emptyPos[idx].row][emptyPos[idx].col] = '.'
		(*rowCache)[emptyPos[idx].row][i] = false
		(*colCache)[emptyPos[idx].col][i] = false
		(*boxCache)[emptyPos[idx].row/3*3+emptyPos[idx].col/3][i] = false
	}
	return false
}