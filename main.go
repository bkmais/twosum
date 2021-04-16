package main

func main() {

	// nums := []int{1, 4, 6, 8}
	// target := 9

	nums := []int{3, 2, 4}
	target := 6

	var res []int
	res = twosum(nums, target)

	for i := 0; i < len(res); i++ {
		println(res[i])
	}

	// strings := []string{"hello", "world"}
	// for i := range strings {
	// 	str1 := strings[i]
	// 	fmt.Println(str1)
	// }
}
