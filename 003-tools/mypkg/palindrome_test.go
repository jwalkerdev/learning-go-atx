package mypkg

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		desc   string
		word   string
		result bool
	}{
		{"True - Odd char count", "racecar", true},
		{"False - Odd char count", "racingcar", false},
		{"True - Even char count", "noon", true},
		{"False - Even char count", "none", false},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			got := IsPalindrome(test.word)
			if got != test.result {
				t.Errorf("%s : want %v got %v", test.word, got, test.result)
			}
		})
	}
}
