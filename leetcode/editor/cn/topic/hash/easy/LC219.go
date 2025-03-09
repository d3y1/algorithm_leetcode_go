// LC219 存在重复元素 II
// @author d3y1

package easy

func containsNearbyDuplicate(nums []int, k int) bool {
    return containsNearbyDuplicate1(nums, k)
    // return containsNearbyDuplicate2(nums, k)
}

// 哈希
func containsNearbyDuplicate1(nums []int, k int) bool {
    // num -> idx
    hash := map[int]int{}
    for i, num := range nums {
        if j, ok := hash[num]; ok && i-j <= k {
            return true
        }
        hash[num] = i
    }

    return false
}

// 滑动窗口
func containsNearbyDuplicate2(nums []int, k int) bool {
    set := map[int]struct{}{}
    for i, num := range nums {
        if i > k {
            delete(set, nums[i-k-1])
        }
        if _, ok := set[num]; ok {
            return true
        }
        set[num] = struct{}{}
    }

    return false
}