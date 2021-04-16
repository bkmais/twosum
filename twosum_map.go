package main

func twosum(nums []int, target int) []int {
	var twosummap = make(map[int]int)
	res := []int{0, 0}
	for i := 0; i < len(nums); i++ {

		iComp := target - nums[i]

		if val, ok := twosummap[iComp]; ok {
			res = []int{i, val}
			return res
		}
		twosummap[nums[i]] = i
	}
	return res
}
