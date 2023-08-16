package goleetcode

// func main() {
// 	print(romanToInt("IV"))
// }

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	fourNineRomanMap := map[string]string{
		"V": "I",
		"X": "I",
		"L": "X",
		"C": "X",
		"D": "C",
		"M": "C",
	}
	totalSum := 0
	for i := range s {
		// Get the character from the back
		reverseIndex := len(s) - i - 1
		r := s[reverseIndex]
		num := romanMap[string(r)]
		fourNineInitial, isFourNine := fourNineRomanMap[string(r)]
		if reverseIndex != 0 && isFourNine && string(s[reverseIndex-1]) == fourNineInitial {
			deduct := romanMap[fourNineInitial]
			totalSum += num - deduct - deduct
			continue
		}
		totalSum += num
	}
	print(totalSum)
	print("\n\n")
	return totalSum
}

// @lc code=end
