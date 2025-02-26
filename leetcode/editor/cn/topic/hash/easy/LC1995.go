// LC1995 统计特殊四元组
// @author d3y1

package easy

func countQuadruplets(nums []int) int {
    // return countQuadruplets1(nums)
    // return countQuadruplets2(nums)
    return countQuadruplets3(nums)
    // return countQuadruplets4(nums)
}

// 暴力循环 - 四重循环
func countQuadruplets1(nums []int) int {
    n := len(nums)
    ans := 0
    for a:=0; a<n; a++ {
        for b:=a+1; b<n; b++ {
            for c:=b+1; c<n; c++ {
                for d:=c+1; d<n; d++ {
                    if nums[a]+nums[b]+nums[c] == nums[d] {
                        ans++
                    }
                }
            }
        }
    }
    return ans
}

// 哈希 - 三重循环
func countQuadruplets2(nums []int) int {
    n := len(nums)
    // num -> cnt
    cnt := map[int]int{}
    ans := 0

    // 固定c
    for c:=n-2; c>=2; c-- {
        cnt[nums[c+1]]++
        for a:=0; a<c; a++ {
            for b:=a+1; b<c; b++ {
                if sum:=nums[a]+nums[b]+nums[c]; cnt[sum]>0 {
                    ans += cnt[sum]
                }
            }
        }
    }
    return ans
}

// 哈希 - 二重循环
func countQuadruplets3(nums []int) int {
    n := len(nums)
    cnt := map[int]int{}
    ans := 0
    // nums[a]+nums[b]+nums[c]=nums[d] => nums[a]+nums[b]=nums[d]-nums[c]
    // 固定b
    for b:=1; b<n-2; b++ {
        for a:=0; a<b; a++{
            cnt[nums[a]+nums[b]]++
        }
        for d:=b+2; d<n; d++ {
            ans += cnt[nums[d]-nums[b+1]]
        }
    }
    return ans
}

// 动态规划(背包)
func countQuadruplets4(nums []int) int {
    // 选择3个数(nums[a]+nums[b]+nums[c])
    n := 3
    // 背包容量(nums[d]<=100)
    v := 100
    ans := 0

    // d[i][j]表示选择i个数累加成j的方案数
    dp := make([][]int, n+1)
    for i:=0; i<=n; i++ {
        dp[i] = make([]int, v+1)
    }

    // 选择0个数累加成0的方案数
    dp[0][0] = 1
    // 枚举d
    for _, num := range nums {
        ans += dp[n][num]
        for i:=n; i>0; i-- {
            for j:=num; j<=v; j++ {
                dp[i][j] += dp[i-1][j-num]
            }
        }
    }

    return ans
}