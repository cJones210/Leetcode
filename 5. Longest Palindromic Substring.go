func longestPalindrome(s string) string {
	//trying to learn go for this one
	lS := len(s)
	if lS <= 1 {
		return s
	}
	oddPalindrome := s[0:1]
	evenPalindrome := s[0:1]
	longest := s[0:1]
	iterateExpand := func(leftPointer int, rightPointer int) string {
		for rightPointer < lS && leftPointer >= 0 && s[rightPointer] == s[leftPointer] {
			rightPointer++
			leftPointer--
		}
		return s[leftPointer+1 : rightPointer]
	}

	for i := 0; i < lS-1; i++ {
		oddPalindrome = iterateExpand(i, i)
		evenPalindrome = iterateExpand(i, i+1)
		if len(oddPalindrome) > len(longest) {
			longest = oddPalindrome
		}
		if len(evenPalindrome) > len(longest) {
			longest = evenPalindrome
		}
	}
	return longest
}

//second time doing this problem and i only came to this solution after looking at the discussions tab and seeing people talk about using pointers
//this is sooo much better then my code before and while it still technically runs in o(n^2) time, 0ms makes me not really care

//0ms beats 100%