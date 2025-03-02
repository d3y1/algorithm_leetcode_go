// LC2085 统计出现过一次的公共字符串
// @author d3y1

package easy

func countWords(words1 []string, words2 []string) int {
    // return countWords1(words1, words2)
    return countWords2(words1, words2)
}

// 哈希 - 单
func countWords1(words1 []string, words2 []string) (ans int) {
    cnt := make(map[string]int)

    for _, word := range words1 {
        cnt[word]++
    }
    for _, word := range words2 {
        if cnt[word] > 1 {
            cnt[word] = -1
        }else{
            cnt[word]--
        }
    }

    for key := range cnt {
        if cnt[key] == 0 {
            ans++
        }
    }

    return
}

// 哈希 - 双
func countWords2(words1 []string, words2 []string) (ans int) {
    cnt1 := make(map[string]int)
    cnt2 := make(map[string]int)

    for _, word := range words1 {
        cnt1[word]++
    }
    for _, word := range words2 {
        cnt2[word]++
    }

    for key := range cnt1 {
        if cnt1[key]==1 && cnt2[key]==1 {
            ans++
        }
    }

    return
}