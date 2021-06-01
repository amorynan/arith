package array

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

var twoSumCases = []struct {
	nums []int
	target int
	verify map[int]struct{}
}{
	{ []int{3,2,4}, 6, map[int]struct{}{1: {}, 2: {}}},
}

func TestTwoSum(t *testing.T) {
	for _, val := range twoSumCases {
		res := twoSum(val.nums, val.target)
		resMap := make(map[int]struct{}, len(res))
		for _, v := range res {
			resMap[v] = struct{}{}
		}
		assert.Equal(t, val.verify, resMap)
	}
}

var threeSumCases = []struct{
	params []int
	res [][]int
}{
	{params: []int{0,0,0,0}, res: [][]int{{0,0,0}}},
	{params: []int{-1,-1,2,2,2}, res: [][]int{{-1,-1,2}}},
	{params: []int{-1,0,1,2,-1,-4,-2,-3,3,0,4}, res: [][]int{{-4,0,4},{-4,1,3},{-3,-1,4},{-3,0,3},{-3,1,2},{-2,-1,3},{-2,0,2},{-1,-1,2},{-1,0,1}}},
	{params: []int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}, res: [][]int{{-4,-2,6},{-4,0,4},{-4,1,3},{-4,2,2},{-2,-2,4},{-2,0,2}}},
}

func TestThreeSum(t *testing.T){
	for _, ca := range threeSumCases {
		res := threeSum_Best(ca.params)
		for _, r := range res {
			sort.Ints(r)
		}
		assert.Equal(t, ca.res, res)
	}
}


var threeSumClosetCases = []struct{
	params []int
	target int
	res int
}{
	{params: []int{-1,2,1,-4}, target: 1, res: 2},
}

func TestThreeSumCloset (t *testing.T) {
	for _, ca := range threeSumClosetCases {
		res := threeSumClosest(ca.params, ca.target)
		assert.Equal(t, ca.res, res)
	}
}

var fourSumCases = []struct{
	params []int
	target int
	res [][]int
}{
	{params: []int{-3,-2,-1,0,0,1,2,3}, target: 0, res: [][]int{{-3,-2,2,3},{-3,-1,1,3},{-3,0,0,3},{-3,0,1,2},{-2,-1,0,3},{-2,-1,1,2},{-2,0,0,2},{-1,0,0,1}}},
}

func TestFourSum (t *testing.T) {
	for _, ca := range fourSumCases{
		res := fourSum(ca.params, ca.target)
		assert.Equal(t, ca.res, res)
	}
}

var repeatedCases = []struct{
	params []int
	len int
	res []int
}{
	//{params: []int{1,1,2}, res: 2},
	{params: []int{0,0,1,1,1,2,2,3,3,4}, len: 5, res: []int{0,1,2,3,4}},
}

func TestReapted(t *testing.T) {
	for _, ca := range repeatedCases{
		incs,res := removeDuplicates(ca.params)
		assert.Equal(t, ca.len, incs)
		assert.Equal(t, ca.res, res)
	}
}

var nextPermutationArr = []struct {
	ca []int
	res []int
}{
	{ca: []int{1,2,3}, res: []int{1,3,2}},
	{ca: []int{4,3,2,1}, res: []int{1,2,3,4}},
	{ca: []int{2,3,1}, res:[]int{3,1,2}},
	{ca: []int{1,5,1}, res:[]int{5,1,1}},
}

func TestNextPermutation(t *testing.T)  {
	for _, ca := range nextPermutationArr {
		res := nextPermutation(ca.ca)
		assert.Equal(t, ca.res, res)
	}
}

var findMedianSortedCases = []struct {
	nums1 []int
	nums2 []int
	target float64
}{
	{ []int{1,3}, []int{2}, 2.00000},
	{ []int{1,2}, []int{3,4}, 2.50000},
	{ []int{2}, []int{1,3,4}, 2.50000},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for _, ca := range findMedianSortedCases{
		res := findMedianSortedArrays(ca.nums1, ca.nums2)
		assert.Equal(t, ca.target, res)
	}
}

var combinationTestCases = []struct{
	target int
	nums   []int
	res [][]int
}{
	//{target: 8, nums: []int{2,3,5}, res: [][]int{{2,2,2,2},{2,3,3},{3,5}}},
	{target: 8, nums: []int{3,1,3,5,1,1},res: [][]int{{1,1,1,5},{1,1,3,3},{3,5}}},
	//{target: 5, nums: []int{2, 5, 2, 1, 2}, res: [][]int{{1,2,2}, {5}}},
}
func TestCombinationSum(t *testing.T)  {
	for _, ca := range combinationTestCases{
		res := combinationSum(ca.nums, ca.target)
		assert.Equal(t, ca.res, res)
	}
}
func TestCombinationSum2(t *testing.T)  {
	for _, ca := range combinationTestCases{
		res := combinationSum2(ca.nums, ca.target)
		assert.Equal(t, ca.res, res)
	}
}


func TestCombinationMultiplication(t *testing.T) {
	res := combinationMultiplication(8)
	for _, v := range res {
		t.Logf("%+v", v)
	}
}


var firstMissPositiveCases = []struct{
	ca []int
	res int
}{
	{ca: []int{-1, 4, 2, 1, 9, 10}, res: 3},
}
func TestFirstMissingPositive(t *testing.T) {
	for _, c := range firstMissPositiveCases {
		res := firstMissingPositive(c.ca)
		assert.Equal(t, c.res, res)
	}
}
