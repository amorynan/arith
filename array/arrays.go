package array

import "sort"

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