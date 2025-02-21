// @author d3y1

package algorighm

func twoSum(nums []int, target int) []int {
	// return solution1(nums, target)
	return solution2(nums, target)
}

// 双重循环
func solution1(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 哈希
func solution2(nums []int, target int) []int {
	// map: val -> idx
	hash := map[int]int{}
	for i, num := range nums {
		if j, ok := hash[target-num]; ok {
			return []int{i, j}
		}
		hash[num] = i
	}
	return nil
}
