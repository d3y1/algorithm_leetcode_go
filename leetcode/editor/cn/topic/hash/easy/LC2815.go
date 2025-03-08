// LC2815 数组中的最大数对和
// @author d3y1

package easy

import "math"

func maxSum(nums []int) int {
    ans := -1

    // digit -> max num
    hash := [10]int{}
    for i := range hash {
        // 不存在 最大数
        hash[i] = math.MinInt
    }

    max := func(a int, b int) int {
        if a < b {
            return b
        }
        return a
    }

    for _, num := range nums {
        // 当前数 数位上最大的数字
        digit := 0
        for val:=num; val>0; val/=10 {
            digit = max(digit, val%10)
        }
        // 当前数 + 最大数(对应)
        ans = max(ans, num+hash[digit])
        hash[digit] = max(hash[digit], num)
    }

    return ans
}