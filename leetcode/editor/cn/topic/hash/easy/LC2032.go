// LC2032 至少在两个数组中出现的值
// @author d3y1

package easy

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    // return twoOutOfThree1(nums1, nums2, nums3)
    // return twoOutOfThree11(nums1, nums2, nums3)
    // return twoOutOfThree111(nums1, nums2, nums3)
    // return twoOutOfThree2(nums1, nums2, nums3)
    return twoOutOfThree3(nums1, nums2, nums3)
}

// 哈希
func twoOutOfThree1(nums1 []int, nums2 []int, nums3 []int) []int {
    return union(union(intersection(nums1,nums2),intersection(nums1,nums3)),intersection(nums2,nums3))
}

// 交集
func intersection(nums1 []int, nums2 []int) []int {
    set := map[int]bool{}
    set1 := map[int]bool{}

    for _, num := range nums1 {
        set1[num] = true
    }
    for _, num := range nums2 {
        if set1[num] {
            set[num] = true
        }
    }

    ans := make([]int, 0, len(set))
    for num := range set {
        ans = append(ans, num)
    }

    return ans
}

// 并集
func union(nums1 []int, nums2 []int) []int {
    set := make(map[int]bool)

    for _, num := range nums1 {
        set[num] = true
    }
    for _, num := range nums2 {
        set[num] = true
    }

    ans := make([]int, 0, len(set))
    for num := range set {
        ans = append(ans, num)
    }

    return ans
}

// 哈希 - set
func twoOutOfThree11(nums1 []int, nums2 []int, nums3 []int) (ans []int) {
    set1 := map[int]bool{}
    set2 := map[int]bool{}
    set3 := map[int]bool{}
    ansSet := map[int]struct{}{}

    for _, num := range nums1 {
        set1[num] = true
    }
    for _, num := range nums2 {
        set2[num] = true
    }
    for _, num := range nums3 {
        set3[num] = true
    }

    for num := range set1 {
        if set2[num] || set3[num] {
            ansSet[num] = struct{}{}
        }
    }
    for num := range set2 {
        if set1[num] || set3[num] {
            ansSet[num] = struct{}{}
        }
    }
    for num := range set3 {
        if set1[num] || set2[num] {
            ansSet[num] = struct{}{}
        }
    }

    for num:= range ansSet {
        ans = append(ans, num)
    }

    return
}

// 哈希 - set - 简化
func twoOutOfThree111(nums1 []int, nums2 []int, nums3 []int) (ans []int) {
    hash := func(nums []int) map[int]bool {
        set := map[int]bool{}
        for _, num := range nums {
            set[num] = true
        }
        return set
    }

    set1, set2, set3 := hash(nums1), hash(nums2), hash(nums3)

    ansSet := map[int]struct{}{}

    check := func(x map[int]bool, y map[int]bool, z map[int]bool) {
        for num := range x {
            if y[num] || z[num] {
                ansSet[num] = struct{}{}
            }
        }
    }

    check(set1, set2, set3)
    check(set2, set1, set3)
    check(set3, set1, set2)

    for num:= range ansSet {
        ans = append(ans, num)
    }

    return
}

// 哈希 - array
func twoOutOfThree2(nums1 []int, nums2 []int, nums3 []int) (ans []int) {
    hash := func(nums []int) (h [101]int) {
        for _, num := range nums {
            h[num] = 1
        }
        return
    }

    h1, h2, h3 := hash(nums1), hash(nums2), hash(nums3)
    for num:=1; num<=100; num++ {
        if h1[num]+h2[num]+h3[num] >= 2 {
            ans = append(ans, num)
        }
    }

    return
}

// 哈希 - mask(位运算)
func twoOutOfThree3(nums1 []int, nums2 []int, nums3 []int) []int {
    // num -> mask
    mask := map[int]int{}
    for i, nums := range [][]int{nums1, nums2, nums3} {
        for _, num := range nums {
            mask[num] |= (1<<i)
        }
    }

    // ans := make([]int, 0)
    ans := []int{}
    for num, m := range mask {
        // mask至少2个1 => num至少在2个数组中出现
        if m&(m-1) > 0 {
            ans = append(ans, num)
        }
    }

    return ans
}