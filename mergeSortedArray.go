package main

import (
	"fmt"
	"sort"
)

// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
// Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
// The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int){

	copy(nums1[m:],nums2)

	sort.Ints(nums1)
}

func main(){
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2,5,6}
	n := 3

	MergeSortedArray(nums1, m, nums2, n)

	fmt.Println(nums1)
}