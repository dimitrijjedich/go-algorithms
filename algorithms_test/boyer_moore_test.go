package algorithms

import (
	"github.com/dimitrijjedich/go-algorithms/algorithms"
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		name     string
		needle   string
		haystack string
		expected int
	}{
		{
			name:     "Exact Match",
			needle:   "abc",
			haystack: "abc",
			expected: 0,
		},
		{
			name:     "Match at Start",
			needle:   "abc",
			haystack: "abcdef",
			expected: 0,
		},
		{
			name:     "Match in Middle",
			needle:   "def",
			haystack: "abcdefghi",
			expected: 3,
		},
		{
			name:     "Match at End",
			needle:   "ghi",
			haystack: "abcdefghi",
			expected: 6,
		},
		{
			name:     "No Match",
			needle:   "xyz",
			haystack: "abcdefghi",
			expected: -1,
		},
		{
			name:     "Empty Needle",
			needle:   "",
			haystack: "abcdefghi",
			expected: -1,
		},
		{
			name:     "Empty Haystack",
			needle:   "abc",
			haystack: "",
			expected: -1,
		},
		{
			name:     "Needle Longer Than Haystack",
			needle:   "abcdefgh",
			haystack: "abc",
			expected: -1,
		},
		{
			name:     "Overlapping Matches",
			needle:   "aba",
			haystack: "abababa",
			expected: 0,
		},
		{
			name:     "Single Character Match",
			needle:   "a",
			haystack: "abcdef",
			expected: 0,
		},
		{
			name:     "Single Character No Match",
			needle:   "z",
			haystack: "abcdef",
			expected: -1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := algorithms.Search([]rune(testCase.needle), []rune(testCase.haystack))
			if result != testCase.expected {
				t.Errorf("Search(%q, %q) = %d; want %d", testCase.needle, testCase.haystack, result, testCase.expected)
			}
		})
	}
}
