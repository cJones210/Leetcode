func longestPalindrome(s string) string {
	lS := len(s)
	longstV := 0
	longstStrt := 0
	longstEnd := 1
	oddfail := false
	evenfail := false
	for i := 0; i < lS; i++ {
		// loops through every char in the string and expands out from each char by v comparing the two
		oddfail = false
		evenfail = false
		for v := 0; i-v >= 0; v++ {
			//this if below only finds the longest odd palindrome
			if i+v < lS && s[i+v] == s[i-v] && !oddfail {
				if 2*v+1 > longstV {
					longstV = 2*v + 1
					longstStrt = i - v
					longstEnd = i + v + 1
				}
			} else {
				oddfail = true
			}
			//finds longest even palindrome
			if i+1+v < lS && s[i-v] == s[i+1+v] && !evenfail {
				if 2*(v+1) > longstV {
					longstV = 2 * (v + 1)
					longstStrt = i - v
					longstEnd = i + 1 + v + 1 // +1 because slice end is exclusive
				}
			} else {
				evenfail = true
			}
		}
	}
	return s[longstStrt:longstEnd]
}

//This is not the fastest method

//even with the expanding out from i idea this code is abhoriently slow at o(n^2)
// i will probably redo this one and this will get commited over

//45ms beats 29.26%