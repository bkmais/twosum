package main

func twosum(nums []int, target int) []int {

	tot := 0
	var resarr = []int{0, 0}
	for i := 0; i < len(nums); i++ {

		for j := 1; j < len(nums); j++ {
			tot = nums[i] + nums[j]
			if target == tot {
				resarr[0] = i
				resarr[1] = j
				break
			} else {
				tot = 0
			}
		}
	}
	return resarr
}
