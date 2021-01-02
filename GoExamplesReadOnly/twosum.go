package GoExamplesReadOnly

import "fmt"

func twoSum(nums []int, target int) []int {
	var s []int
	for i, n := range nums {
		for j, m := range nums {
			if n != m && target == nums[i]+nums[j] {
				s = append(s, i)
				s = append(s, j)
				fmt.Println(s, target, nums[i], nums[j])
				return s
			} else {
				fmt.Println("no solutons")
			}
		}
	}
	return s
}

//func mainfalse() {
//	nums := []int{2, 7, 9, 11}
//	target := 9
//	twoSum(nums, target)
//}
