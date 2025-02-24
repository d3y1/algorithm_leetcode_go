// LC1941 检查是否所有字符出现次数相同
// @author d3y1

package easy

func areOccurrencesEqual(s string) bool {
    // return areOccurrencesEqual1(s)
    // return areOccurrencesEqual2(s)
    return areOccurrencesEqual3(s)
    // return areOccurrencesEqual4(s)
}

func areOccurrencesEqual1(s string) bool {
    hash := map[rune]int{}

    for _,x := range s {
        hash[x]++
    }

    times := 0
    for _,val := range hash {
        if times == 0 {
            times = val
        }else if val != times {
            return false
        }
    }

    return true
}

func areOccurrencesEqual2(s string) bool {
    hash := map[rune]int{}

    for _,x := range s {
        hash[x]++
    }

    times := hash[rune(s[0])];
    for _,val := range hash {
        if val != times {
            return false
        }
    }

    return true
}

func areOccurrencesEqual3(s string) bool {
    hash := map[byte]int{}

    for i:=0; i<len(s); i++ {
        hash[s[i]]++
    }

    times := hash[s[0]]
    for _,val := range hash {
        if val != times {
            return false
        }
    }

    return true
}

func areOccurrencesEqual4(s string) bool {
    hash := [26]int{}

    for _,x := range s {
        hash[x-'a']++
    }

    times := 0
    for _,val := range hash {
        if val == 0 {
            continue
        }
        if times == 0 {
            times = val
        }else if val != times {
            return false
        }
    }

    return true
}