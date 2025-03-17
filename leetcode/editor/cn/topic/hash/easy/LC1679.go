// LC1679 K 和数对的最大数目
// @author d3y1

package easy

func maxOperations(nums []int, k int) int {
    // return maxOperations1(nums, k)
    return maxOperations2(nums, k)
}

// 哈希
func maxOperations1(nums []int, k int) (ans int) {
    cnt := map[int]int{}
    for _, num := range nums {
        if val, ok := cnt[k-num]; ok {
            if val > 0 {
                ans++
                cnt[k-num]--
            }else{
                cnt[num]++
            }
        }else{
            cnt[num]++
        }
    }
    return
}

// 哈希 - 简化
func maxOperations2(nums []int, k int) (ans int) {
    cnt := map[int]int{}
    for _, num := range nums {
        if cnt[k-num] > 0 {
            ans++
            cnt[k-num]--
        }else{
            cnt[num]++
        }
    }
    return
}