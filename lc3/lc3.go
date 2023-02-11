package lc3

import "fmt"

//func lengthOfLongestSubstring(s string) int {
//	j := 0
//	maxlen := 0
//	substr := ""
//	for i := 0; i < len(s); i++ {
//		fmt.Println(string(s[i]))
//		for {
//			if strings.Contains(substr, string(s[i])) {
//				fmt.Println("包含", substr)
//				if j < i {
//					j++
//					substr = s[j:i]
//				}
//			} else {
//				break
//			}
//		}
//		substr = s[j : i+1]
//		if maxlen < len(substr) {
//			maxlen = len(substr)
//		}
//	}
//	return maxlen
//}

func lengthOfLongestSubstring(s string) int {
	j := 0
	l := len(s)
	ml := 0
	sub := ""
	for i := 0; i < l; i++ {
		for k, v := range sub {
			if uint8(v) == s[i] {
				if j < i {
					j = j + k
					sub = s[j:i]
					fmt.Println("k, j", k, j, sub)
				}
			}
		}
		sub = s[j:i]
		if ml < len(sub) {
			ml = len(sub)
		}
	}
	fmt.Println(sub, ml)
	return ml
}
