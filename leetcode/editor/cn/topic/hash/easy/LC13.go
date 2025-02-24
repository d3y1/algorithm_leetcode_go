// LC13 罗马数字转整数
// @author d3y1

package easy

func romanToInt(s string) int {
    romans := map[byte]int {
        'I' : 1,
        'V' : 5,
        'X' : 10,
        'L' : 50,
        'C' : 100,
        'D' : 500,
        'M' : 1000,
    }

    result := 0
    n := len(s)
    for i:=1; i<n; i++ {
        x, y := romans[s[i-1]], romans[s[i]]
        if x < y {
            result -= x
        }else{
            result += x
        }
    }

    result += romans[s[n-1]]

    return result
}