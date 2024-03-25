package main

import(
	"fmt"
)

var p = fmt.Println

func main(){
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2,5,6,6,6}
	n := 3

	MergeSortedArray(nums1, m, nums2, n)
	p(nums1)
	p(RemoveElements(nums1, m)) 
	p(RemoveDuplicates(nums1))
	p(MajorityElements(nums2))
}