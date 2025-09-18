class Solution(object):
    def lengthOfLongestSubstring(self, s):
        #trying to use 'pointers' and hashmaps for this one
        hashmap = {}
        left = 0
        right = 0
        maxL = 0
        for right in range(len(s)):
            if s[right] in hashmap:
                left = max(left, hashmap[s[right]] + 1)
            hashmap[s[right]] = right
            maxL = max(maxL, right - left + 1)
        return maxL

#really enjoyed this one, the sorta 'sliding box' defined by the left and right 'pointers' are really different than anything ive done before
#15ms beats 85.85%, but as with 2. the exact execution time is quite random due to python. I do believe this is pretty close to the fastest way possible besides small things