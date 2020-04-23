package main

import "testing"

func TestSameNecklace(t *testing.T) {

	necklaces := []struct {
		a        string
		b        string
		expected bool
	}{
		{"nicole", "icolen", true},
		{"nicole", "lenico", true},
		{"nicole", "coneli", false},
		{"aabaaaaabaab", "aabaabaabaaa", true},
		{"abc", "cba", false},
		{"xxyyy", "xxxyy", false},
		{"xyxxz", "xxyxz", false},
		{"x", "x", true},
		{"x", "xx", false},
		{"x", "", false},
		{"", "", true},
	}

	for _, tt := range necklaces {
		t.Run(tt.a, func(t *testing.T) {
			got := sameNecklace(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("'%v' '%v', got %v, wanted %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestRepeats(t *testing.T) {

	testcases := []struct {
		a        string
		expected int
	}{
		{"abc", 1},
		{"abcabcabc", 3},
		{"abcabcabcx", 1},
		{"aaaaaa", 6},
		{"a", 1},
		{"", 1},
	}

	for _, tt := range testcases {
		t.Run(tt.a, func(t *testing.T) {
			got := repeats(tt.a)
			if got != tt.expected {
				t.Errorf("'%v', got %v, wanted %v", tt.a, got, tt.expected)
			}
		})
	}
}

func TestCanoicalize(t *testing.T) {

	testcases := []struct {
		a        string
		expected string
	}{
		{"abc", "abc"},
		{"abcabcabc", "abcabcabc"},
		{"abcabcabcx", "abcabcabcx"},
		{"aaaaaa", "aaaaaa"},
		{"fish", "fish"},
		{"apple", "apple"},
		{"mississippi", "imississipp"},
		{"racecar", "acecarr"},
		{"dcba", "adcb"},
	}

	for _, tt := range testcases {
		t.Run(tt.a, func(t *testing.T) {
			got := canonicalize(tt.a)
			if got != tt.expected {
				t.Errorf("'%v', got %v, wanted %v", tt.a, got, tt.expected)
			}
		})
	}
}
