package main

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"zero", 0, true},
		{"single digit", 5, true},
		{"negative number", -121, false},
		{"palindrome 121", 121, true},
		{"palindrome 1221", 1221, true},
		{"palindrome 12321", 12321, true},
		{"not palindrome 123", 123, false},
		{"not palindrome 1234", 1234, false},
		{"large palindrome", 123454321, true},
		{"large not palindrome", 123456789, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("isPalindrome(%d) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsPalindromeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"single char", "a", true},
		{"palindrome aba", "aba", true},
		{"palindrome abba", "abba", true},
		{"palindrome 12321", "12321", true},
		{"palindrome with spaces", "a b a", true}, // compares characters positionally
		{"not palindrome abc", "abc", false},
		{"not palindrome abcd", "abcd", false},
		{"palindrome with decimal", "50.05", true},
		{"not palindrome 32.14", "32.14", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindromeString(tt.input)
			if result != tt.expected {
				t.Errorf("isPalindromeString(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", ""},
		{"single char", "a", "a"},
		{"hello", "hello", "olleh"},
		{"palindrome", "aba", "aba"},
		{"numbers", "12345", "54321"},
		{"unicode", "🚗✨", "✨🚗"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverse(tt.input)
			if result != tt.expected {
				t.Errorf("reverse(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGeneratePalindromesForDigits(t *testing.T) {
	tests := []struct {
		name        string
		digits      int
		expectedLen int
		firstFew    []int
	}{
		{"1 digit", 1, 9, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"2 digits", 2, 9, []int{11, 22, 33, 44, 55, 66, 77, 88, 99}},
		{"3 digits", 3, 90, []int{101, 111, 121, 131, 141, 151, 161, 171, 181, 191}},
		{"4 digits", 4, 90, []int{1001, 1111, 1221, 1331, 1441, 1551, 1661, 1771, 1881, 1991}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generatePalindromesForDigits(tt.digits)
			if len(result) != tt.expectedLen {
				t.Errorf("generatePalindromesForDigits(%d) length = %d, want %d", tt.digits, len(result), tt.expectedLen)
			}
			if len(result) >= len(tt.firstFew) {
				for i, expected := range tt.firstFew {
					if result[i] != expected {
						t.Errorf("generatePalindromesForDigits(%d)[%d] = %d, want %d", tt.digits, i, result[i], expected)
					}
				}
			}
		})
	}
}

func TestFormatPounds(t *testing.T) {
	tests := []struct {
		name     string
		pence    int
		expected string
	}{
		{"zero", 0, "0.00"},
		{"single pence", 1, "0.01"},
		{"ten pence", 10, "0.10"},
		{"one pound", 100, "1.00"},
		{"one pound one pence", 101, "1.01"},
		{"large amount", 12345, "123.45"},
		{"palindrome pence", 3223, "32.23"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatPounds(tt.pence)
			if result != tt.expected {
				t.Errorf("formatPounds(%d) = %q, want %q", tt.pence, result, tt.expected)
			}
		})
	}
}

func TestIsEffectivelyInteger(t *testing.T) {
	tests := []struct {
		name     string
		f        float64
		epsilon  float64
		expected bool
	}{
		{"exact integer", 5.0, 0.01, true},
		{"close to integer", 5.001, 0.01, true},
		{"not close enough", 5.02, 0.01, false},
		{"negative close", -3.001, 0.01, true},
		{"decimal", 3.14, 0.01, false},
		{"large epsilon", 5.1, 0.2, true},
		{"zero", 0.0, 0.01, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isEffectivelyInteger(tt.f, tt.epsilon)
			if result != tt.expected {
				t.Errorf("isEffectivelyInteger(%f, %f) = %v, want %v", tt.f, tt.epsilon, result, tt.expected)
			}
		})
	}
}

func TestGetPalindromicPencesInRange(t *testing.T) {
	tests := []struct {
		name     string
		minPence int
		maxPence int
		expected []int
		checkLen bool // if true, just check length matches (for large ranges)
	}{
		{"small range with palindromes", 10, 50, []int{11, 22, 33, 44}, false},
		{"range with 1-3 digits", 1, 999, nil, true}, // just check we get results
		{"empty range", 50, 10, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getPalindromicPencesInRange(tt.minPence, tt.maxPence)
			if tt.checkLen {
				if len(result) == 0 {
					t.Errorf("getPalindromicPencesInRange(%d, %d) returned empty slice, expected some results", tt.minPence, tt.maxPence)
				}
			} else {
				// For empty ranges, just check that result is empty
				if tt.minPence > tt.maxPence {
					if len(result) != 0 {
						t.Errorf("getPalindromicPencesInRange(%d, %d) returned %d results, expected 0 for empty range", tt.minPence, tt.maxPence, len(result))
					}
				} else if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("getPalindromicPencesInRange(%d, %d) = %v, want %v", tt.minPence, tt.maxPence, result, tt.expected)
				}
			}
		})
	}
}

func TestFindPalindromicFuelCosts(t *testing.T) {
	tests := []struct {
		name           string
		pricePerLitre  float64
		maxLitres      int
		expectedCount  int // check count since full results would be long
		checkSpecific  bool
		specificResult *Result
	}{
		{"standard price", 128.9, 100, 4, true, &Result{
			Litres:             25.0,
			CostPounds:         "32.23",
			LitresIsPalindrome: false, // 25 is not a palindrome
			Type:               "whole",
		}},
		{"zero price", 0, 10, 0, false, nil},
		{"very high price", 1000, 1, 0, false, nil}, // no results in range
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := FindPalindromicFuelCosts(tt.pricePerLitre, tt.maxLitres)
			if len(results) != tt.expectedCount {
				t.Errorf("FindPalindromicFuelCosts(%f, %d) returned %d results, want %d",
					tt.pricePerLitre, tt.maxLitres, len(results), tt.expectedCount)
			}

			if tt.checkSpecific && tt.specificResult != nil {
				found := false
				for _, result := range results {
					if result.Litres == tt.specificResult.Litres &&
						result.CostPounds == tt.specificResult.CostPounds &&
						result.LitresIsPalindrome == tt.specificResult.LitresIsPalindrome &&
						result.Type == tt.specificResult.Type {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("FindPalindromicFuelCosts(%f, %d) did not contain expected result: %+v",
						tt.pricePerLitre, tt.maxLitres, tt.specificResult)
				}
			}
		})
	}
}

func TestFindNearestPalindromicCost(t *testing.T) {
	tests := []struct {
		name          string
		pricePerLitre float64
		targetLitres  float64
		searchRadius  int
		expectResult  bool
	}{
		{"find near 25", 128.9, 25.0, 10, true},
		{"find near 30", 128.9, 30.0, 10, true},
		{"no result in radius", 128.9, 1000.0, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindNearestPalindromicCost(tt.pricePerLitre, tt.targetLitres, tt.searchRadius)
			if tt.expectResult && result == nil {
				t.Errorf("FindNearestPalindromicCost(%f, %f, %d) expected result but got nil",
					tt.pricePerLitre, tt.targetLitres, tt.searchRadius)
			} else if !tt.expectResult && result != nil {
				t.Errorf("FindNearestPalindromicCost(%f, %f, %d) expected nil but got result",
					tt.pricePerLitre, tt.targetLitres, tt.searchRadius)
			}
		})
	}
}

func TestParseFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"valid float", "32.23", 32.23},
		{"integer string", "50", 50.0},
		{"negative", "-10.5", -10.5},
		{"zero", "0.00", 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseFloat(tt.input)
			if result != tt.expected {
				t.Errorf("parseFloat(%q) = %f, want %f", tt.input, result, tt.expected)
			}
		})
	}
}
