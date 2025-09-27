func lengthOfLongestSubstring(s string) int {
    hashmap := make(map[byte]int)
    var left int
    var right int
    var maxL int
    for right=0; right < len(s); right++ {
        if _, found := hashmap[s[right]]; found {
            left = max(left, hashmap[s[right]] + 1)
        }
        hashmap[s[right]] = right
        maxL = max(maxL, right - left + 1)
    }
    return maxL
}
// i just completely copied my python solution but in go to try to practice more with go hashmaps

// 0ms beats 100%
