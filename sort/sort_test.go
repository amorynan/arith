package sort

import "testing"

func TestDistributionAndMerge (t *testing.T) {
	target := []int {7,6,2,4,1,9,3,8,0,5}
	res := DistributionAndMerge(target)
	t.Logf("res :%+v", res)
}
