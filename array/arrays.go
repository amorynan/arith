package array

import "sort"

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
