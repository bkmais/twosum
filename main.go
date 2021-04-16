package main

func main() {

	nums := []int{1, 4, 6, 8}
	target := 9

	var res []int
	res = twosum(nums, target)

	for i := 0; i < len(res); i++ {
		println(res[i])
	}
}
