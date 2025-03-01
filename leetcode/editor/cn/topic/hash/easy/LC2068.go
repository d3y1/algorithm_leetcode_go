// LC2068 检查两个字符串是否几乎相等
// @author d3y1

package easy

func checkAlmostEquivalent(word1 string, word2 string) bool {
    // return checkAlmostEquivalent1(word1, word2)
    // return checkAlmostEquivalent11(word1, word2)
    return checkAlmostEquivalent111(word1, word2)
    // return checkAlmostEquivalent2(word1, word2)
    // return checkAlmostEquivalent22(word1, word2)
    // return checkAlmostEquivalent3(word1, word2)
}

func checkAlmostEquivalent1(word1 string, word2 string) bool {
    cnt1 := map[int]int{}
    cnt2 := map[int]int{}

    for _, letter := range word1 {
        cnt1[int(letter)-'a']++
    }
    for _, letter := range word2 {
        cnt2[int(letter)-'a']++
    }

    abs := func(num1 int, num2 int) (ans int) {
        if num1 <= num2 {
            ans = num2-num1
        }else{
            ans = num1-num2
        }
        return
    }

    for i:=0; i<26; i++ {
        if abs(cnt1[i], cnt2[i]) > 3 {
            return false
        }
    }

    return true
}

func checkAlmostEquivalent11(word1 string, word2 string) bool {
    diff := map[int]int{}

    for _, letter := range word1 {
        diff[int(letter)-'a']++
    }
    for _, letter := range word2 {
        diff[int(letter)-'a']--
    }

    abs := func(num int) (ans int) {
        if num >= 0 {
            ans = num
        }else{
            ans = -num
        }
        return
    }

    for i:=0; i<26; i++ {
        if abs(diff[i]) > 3 {
            return false
        }
    }

    return true
}

func checkAlmostEquivalent111(word1 string, word2 string) bool {
    diff := [26]int{}

    for _, letter := range word1 {
        diff[letter-'a']++
    }
    for _, letter := range word2 {
        diff[letter-'a']--
    }

    for _, val := range diff {
        if val < -3 || 3 < val {
            return false
        }
    }

    return true
}

func checkAlmostEquivalent2(word1 string, word2 string) bool {
    cnt1 := map[rune]int{}
    cnt2 := map[rune]int{}
    set := map[rune]bool{}

    for _, letter := range word1 {
        cnt1[letter]++
        set[letter] = true
    }
    for _, letter := range word2 {
        cnt2[letter]++
        set[letter] = true
    }

    abs := func(num1 int, num2 int) (ans int) {
        if num1 <= num2 {
            ans = num2-num1
        }else{
            ans = num1-num2
        }
        return
    }

    for key := range set {
        if abs(cnt1[key], cnt2[key]) > 3 {
            return false
        }
    }

    return true
}

func checkAlmostEquivalent22(word1 string, word2 string) bool {
    cnt := map[rune]int{}
    set := map[rune]bool{}

    for _, letter := range word1 {
        cnt[letter]++
        set[letter] = true
    }
    for _, letter := range word2 {
        cnt[letter]--
        set[letter] = true
    }

    for key := range set {
        if cnt[key] < -3 || 3 < cnt[key] {
            return false
        }
    }

    return true
}

func checkAlmostEquivalent3(word1 string, word2 string) bool {
    diff := make(map[byte]int)

    for i:=0; i<len(word1); i++ {
        diff[word1[i]]++
    }
    for i:=0; i<len(word2); i++ {
        diff[word2[i]]--
    }

    for _, val := range diff {
        if val < -3 || 3 < val {
            return false
        }
    }
    return true
}