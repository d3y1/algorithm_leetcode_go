// LC2099 找到和最大的长度为 K 的子序列
// @author d3y1

package easy

import "sort"

func maxSubsequence(nums []int, k int) []int {
    // return maxSubsequence1(nums, k)
    return maxSubsequence2(nums, k)
}

func maxSubsequence1(nums []int, k int) []int {
    n := len(nums)
    idx := make([]int, n)
    for i := range idx {
        idx[i] = i
    }

    sort.Slice(idx, func(i int, j int) bool {
        return nums[idx[i]] > nums[idx[j]]
    })
    sort.Ints(idx[:k])

    ans := make([]int, k)
    for i, j := range idx[:k] {
        ans[i] = nums[j]
    }

    return ans
}

func maxSubsequence2(nums []int, k int) []int {
    n := len(nums)
    nums1 := make([]int, n)

    copy(nums1, nums)
    sort.Ints(nums1)

    cnt := make(map[int]int)
    for _, num := range nums1[n-k:] {
        cnt[num]++
    }

    ans := make([]int, 0)
    for _, num := range nums {
        if cnt[num] > 0 {
            ans = append(ans, num)
            cnt[num]--
        }
    }

    return ans
}