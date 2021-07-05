package array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestArrayMake(t *testing.T)  {
	ta := make([]int, 0, 1)
	assert.NotNil(t, ta)
	ta = append(ta, 1)
	assert.Equal(t, len(ta),1)

	tasss :=  make([]int, 0)
	assert.NotNil(t, ta)
	tasss = append(tasss, 1)
	assert.Equal(t, len(tasss),1)

	tas := make([]int, 35, 35)
	tas = append(tas, 1)
	assert.Equal(t, len(tas),36)
	assert.Equal(t, 1, tas[35])

	tass := make([]int, 35)
	tass = append(tass, 1)
	assert.Equal(t, len(tass),36)
	assert.Equal(t, 1, tas[35])

	tt := make([]int, 0, 35)
	tt = append(tt, 1)
	assert.Equal(t, len(tt), 1)
	assert.Equal(t, 1, tt[0])

}

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


var trapTest = []struct{
	ca []int
	res int
}{
	{ca: []int{0,1,0,2,1,0,1,3,2,1,2,1}, res: 6},
	{ca: []int{4,2,0,3,2,5}, res: 9},
	{ca: []int{5,5,1,7,1,1,5,2,7,6}, res: 23},
}
func TestTrap(t *testing.T)  {
	for _, ca := range trapTest {
		res := trap_best(ca.ca)
		assert.Equal(t, ca.res, res)
	}
}

var jumpTest = []struct{
	ca []int
	res int
}{
	{ca: []int{2,3,1}, res: 1},
	{ca: []int{2, 3, 1, 1, 4}, res: 2},
	{ca: []int{1, 1, 1, 1, 1}, res: 4},
	{ca: []int{1}, res: 0},
	{ca: []int{1,2}, res: 1},
	{ca: []int{1,2,1,1,1}, res: 3},
	{ca: []int{10,9,8,7,6,5,4,3,2,1,1,0}, res: 2},
}

func TestJump(t *testing.T) {
	for _, ca := range jumpTest {
		jumps := jump(ca.ca)
		assert.Equal(t, ca.res, jumps)
	}
}

var canJumpCases = []struct{
	ca []int
	res bool
}{
	{ca: []int{1,1,2,2,0,1,1}, res: true},
	{ca: []int{2,3,1,1,4}, res: true},
	{ca: []int{3,2,1,0,4}, res: false},
	{ca: []int{1,1,1,0}, res: true},
	{ca: []int{1,1,0,1}, res: false},
}

func TestCanJump(t *testing.T) {
	for _, ca := range canJumpCases {
		res := canJump(ca.ca)
		assert.Equal(t, ca.res, res)
	}
}

var rotateMatrix = []struct{
	matrix [][]int
	res    [][]int
} {
	//{matrix: [][]int{{1}}, res: [][]int{{1}}},
	{matrix: [][]int{{1,2,3},{4,5,6}, {7,8,9}}, res: [][]int{{7,4,1}, {8,5,2}, {9,6,3}}},
}

func TestRotate(t *testing.T) {
	for _, ca := range rotateMatrix {
		rotate(ca.matrix)
		assert.Equal(t, ca.res, ca.matrix)
	}
}

var spiralOrderCase = []struct{
	matrix [][]int
	res []int
}{
	{matrix: [][]int{{1,2,3}, {4,5,6}, {7,8,9}}, res: []int{1,2,3,6,9,8,7,4,5}},
	{matrix: [][]int{{1,2,3,4}, {5,6,7,8}, {9,10,11,12}}, res: []int{1,2,3,4,8,12,11,10,9,5,6,7}},
}
func TestSpiralOrder(t *testing.T) {
	for _, ca := range spiralOrderCase {
		res := spiralOrder(ca.matrix)
		assert.Equal(t, ca.res, res)
	}
}

var generateMatrixCase = []struct{
	length int
	res [][]int
}{
	{length: 1, res: [][]int{{1}}},
	{length: 2, res: [][]int{{1,2}, {4,3}}},
	{length: 3, res: [][]int{{1,2,3},{8,9,4}, {7,6,5}}},
	{length: 4, res: [][]int{{1,2,3,4},{12,13,14,5},{11,16,15,6},{10,9,8,7}}},
}

func TestGenerateMatrix(t *testing.T)  {
	for _, ca := range generateMatrixCase {
		res := generateMatrix(ca.length)
		assert.Equal(t, ca.res, res)
	}
}

var pathCases = []struct {
	m   int
	n   int
	res int
}{
	{m: 3, n:2, res: 3},
	{m: 3, n:7, res: 28},
	{m: 51, n:9, res: 1916797311},
}

func TestUniquePath(t *testing.T)  {
	for _, ca := range pathCases {
		res := uniquePaths(ca.m, ca.n)
		assert.Equal(t, ca.res, res)
	}
}

var setCase = []struct{
	ma [][]int
	res [][]int
}{
	{ma: [][]int{{1,0,3}}, res: [][]int{{0,0,0}}},
	{ma: [][]int{{1,1,1}, {1,0,1}, {1,1,1}}, res: [][]int{{1,0,1}, {0,0,0}, {1,0,1}}},
	{ma: [][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}}, res: [][]int{{0,0,0,0},{0,4,5,0},{0,3,1,0}}},
}
func TestSetZeros(t *testing.T)  {
	for _, ca := range setCase {
		//setZeroes(ca.ma)
		setZeroesWithBit(ca.ma)
		assert.Equal(t, ca.res, ca.ma)
	}
}

var searchMatrixCase = []struct {
	matrix [][]int
	target int
	find  bool
}{
	{matrix: [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}, target: 9, find: false},
}
func TestSearchMatrix(t *testing.T) {
	for _, ca := range searchMatrixCase {
		res := searchMatrix(ca.matrix, ca.target)
		assert.Equal(t, ca.find, res)
	}

}
var sudoCase = []struct {
	matrix [][]byte
	find  bool
}{
	{matrix: [][]byte{{'5','3','.','.','7','.','.','.','.'},{'6','.','.','1','9','5','.','.','.'},{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},{'4','.','.','8','.','3','.','.','1'},{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'}, {'.','.','.','4','1','9','.','.','5'},{'.','.','.','.','8','.','.','7','9'}},  find: true},
}
func TestSudoku(t *testing.T){
	for _, ca := range sudoCase {
		res := isValidSudoku(ca.matrix)
		assert.Equal(t, ca.find, res)

	}
}

func TestSolveSudoku(t *testing.T) {
	for _, ca := range sudoCase {
		solveSudoku(ca.matrix)
		fmt.Printf("%+v\v", ca.matrix)
	}
}

func TestOp(t *testing.T) {
	r := 0
	s := 0
	s ^= 1 << r
	t.Log(s)
	s ^= 1 << r
	t.Log(s)
	s ^= 1 << r
	t.Log(s)
	//test := []int{1,3,4}
	//test = append(test[:2], test[3:]...)
	//t.Log(test)
	//a := uint(1)
	//b := ^a
	//c := b^a
	//t.Logf("b :%v, c:%v", b, c)
	//t.Log(bits.TrailingZeros8(uint8(10)))
}

var permutationCases = []struct{
	nums []int
}{
	//{nums: []int{1,2,3}},
	//{nums: []int{1, 2}},
	{nums: []int{1,1,2,2,2}},
}

func TestPermutationCase(t *testing.T) {
	for _, ca := range permutationCases {
		res := Permutation(ca.nums)
		resUnique := PermuteUnique(ca.nums)
		t.Logf("not unique:%v", res)
		t.Logf("unique:%v", resUnique)

	}
}

var queueCase = []struct {
	n int
	res [][]string
}{
	{n: 4, res: [][]string{{".Q..","...Q","Q...","..Q."}, {"..Q.","Q...", "...Q", ".Q.." }}},
}

func TestSolveQueue(t *testing.T)  {
	for _, ca := range queueCase {
		res := solveNQueens(ca.n)
		assert.Equal(t, res, ca.res)
	}
}

var setColorsCases = []struct {
	n []int
	res []int
}{
	{n:[]int{2,0,2,1,1,0}, res: []int{0,0,1,1,2,2}},
}

func TestSetColors(t *testing.T) {
	for _, ca := range setColorsCases {
		sortColors(ca.n)
		assert.Equal(t,ca.res, ca.n)
	}
}

