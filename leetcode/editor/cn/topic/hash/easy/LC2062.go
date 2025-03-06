// LC2062 统计字符串中的元音子字符串
// @author d3y1

package easy

import "strings"

func countVowelSubstrings(word string) int {
    // return countVowelSubstrings1(word)
    return countVowelSubstrings2(word)
}

// 哈希
func countVowelSubstrings1(word string) int {
    // 是否元音
    isVowel := func(ch rune) bool {
        //if ch=='a' || ch=='e' || ch=='i' || ch=='o' || ch=='u' {
        //    return true
        //}

        //switch(ch){
        //    case 'a': return true
        //    case 'e': return true
        //    case 'i': return true
        //    case 'o': return true
        //    case 'u': return true
        //    default: return false
        //}

        switch(ch){
        case 'a','e','i','o','u': return true
        default: return false
        }
    }

    // 是否全部5种元音子字符串
    isVowelStr := func(cnt map[rune]int) bool {
        for _, ch := range "aeiou" {
            if cnt[ch] == 0 {
                return false
            }
        }
        return true
    }

    n := len(word)
    ans := 0
    for i:=0; i<n; i++ {
        chI := rune(word[i])
        if !isVowel(chI) {
            continue
        }

        cnt := map[rune]int{}
        for j:=i; j<n; j++ {
            chJ := rune(word[j])
            if !isVowel(chJ) {
                break
            }
            cnt[chJ]++
            if j-i+1 >= 5 {
                if isVowelStr(cnt) {
                    ans++
                }
            }
        }
    }

    return ans
}

// 哈希
func countVowelSubstrings2(word string) (ans int) {
    // 分割出仅包含元音的字符串
    for _, part := range strings.FieldsFunc(word, func(r rune) bool { return !strings.ContainsRune("aeiou", r) }) {
        // aeiou -> v(上限)
        cnt := ['v']int{}
        left := 0
        for _, ch := range part {
            cnt[ch]++
            // 左指针 当且仅当该元音个数不止一个时才移动左指针
            for cnt[part[left]] > 1 {
                cnt[part[left]]--
                left++
            }
            if cnt['a'] > 0 && cnt['e'] > 0 && cnt['i'] > 0 && cnt['o'] > 0 && cnt['u'] > 0 {
                ans += left+1
            }
        }
    }

    return
}