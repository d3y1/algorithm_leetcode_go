// LC2053 数组中第 K 个独一无二的字符串
// @author d3y1

package easy

func kthDistinct(arr []string, k int) (ans string) {
    // cnt := make(map[string]int)
    cnt := map[string]int{}
    for _, str := range arr {
        cnt[str]++
    }

    seq := 0
    for _, str := range arr {
        if cnt[str] == 1 {
            seq++
        }
        if seq == k {
            ans = str
            break
        }
    }

    return
}