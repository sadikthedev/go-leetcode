package main

import(
	"fmt"
)

func main(){
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2,5,6}
	n := 3

	MergeSortedArray(nums1, m, nums2, n)
	fmt.Println(nums1)
	fmt.Println(RemoveElements(nums1, m)) 
	fmt.Println(RemoveDuplicates(nums1))	
}