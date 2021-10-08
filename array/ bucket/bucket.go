package _bucket

/**
	bucket sort 是非常重要的算法在 大数据领域 需要熟练掌握
 */

type pair struct {
	min , max int
}

func maximumGap(nums []int) int {

}


func min(a... int) int {
	res := a[0]
	for i := 1; i < len(a); i++{
		if a[i] < res {
			res = a[i]
		}
	}
	return res
}

func max(a... int) int {
	res := a[0]
	for i := 1; i < len(a); i++{
		if a[i] > res {
			res = a[i]
		}
	}
	return res
}
