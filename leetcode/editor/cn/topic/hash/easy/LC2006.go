// LC2006 差的绝对值为 K 的数对数目
// @author d3y1

package easy

import "math"

func countKDifference(nums []int, k int) int {
    // return countKDifference1(nums, k)
    // return countKDifference2(nums, k)
    // return countKDifference3(nums, k)
    return countKDifference4(nums, k)
}

// 双重循环
func countKDifference1(nums []int, k int) int {
    n := len(nums)
    cnt := 0
    for i:=0; i<n; i++ {
        for j:=i+1; j<n; j++ {
            if int(math.Abs(float64(nums[i]-nums[j])))==k {
                cnt++
            }
        }
    }

    return cnt
}

// 双重循环
func countKDifference2(nums []int, k int) int {
    n := len(nums)
    cnt := 0
    for i:=0; i<n; i++ {
        for j:=i+1; j<n; j++ {
            if abs(nums[i]-nums[j])==k {
                cnt++
            }
        }
    }

    return cnt
}

// range
func countKDifference3(nums []int, k int) int {
    cnt := 0
    for i, x := range nums {
        for _, y := range nums[:i] {
            if abs(x-y) == k {
                cnt++
            }
        }
    }

    return cnt
}

// 绝对值
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

// 哈希
func countKDifference4(nums []int, k int) int {
    ans := 0
    cnt := map[int]int{}
    for _, num := range nums {
        ans += cnt[num-k]+cnt[num+k]
        cnt[num]++
    }

    return ans
}