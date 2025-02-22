// LC1 两数之和
// @author d3y1

package easy

func twoSum(nums []int, target int) []int {
    // return twoSum1(nums, target)
    return twoSum2(nums, target)
}

// 双重循环
func twoSum1(nums []int, target int) []int {
    for i,num := range nums {
        for j:=i+1; j<len(nums); j++ {
            if num+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}

// 哈希
func twoSum2(nums []int, target int) []int {
    // map: val -> idx
    hash := map[int]int{}
    for i,num := range nums {
        if j,ok := hash[target-num]; ok {
            return []int{i, j}
        }
        hash[num] = i
    }
    return nil
}
