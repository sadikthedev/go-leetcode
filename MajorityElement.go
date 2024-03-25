package main

import(

)

func MajorityElements(nums []int) int{
	k := 0
	count := 0

	for _, num := range nums {
		if count == 0{
			k = num
		}
		if num == k{
			count++
		}else {
			count--
		}
	}
	return k
}