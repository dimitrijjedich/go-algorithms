package algorithms

import (
	"github.com/dimitrijjedich/go-algorithms/algorithms"
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		name           string
		needle         string
		haystack       string
		expectedResult int
		expectedError  string
	}{
		{
			name:           "Exact Match",
			needle:         "abc",
			haystack:       "abc",
			expectedResult: 0,
			expectedError:  "",
		},
		{
			name:           "Match at Start",
			needle:         "abc",
			haystack:       "abcdef",
			expectedResult: 0,
			expectedError:  "",
		},
		{
			name:           "Match in Middle",
			needle:         "def",
			haystack:       "abcdefghi",
			expectedResult: 3,
			expectedError:  "",
		},
		{
			name:           "Match at End",
			needle:         "ghi",
			haystack:       "abcdefghi",
			expectedResult: 6,
			expectedError:  "",
		},
		{
			name:           "No Match",
			needle:         "xyz",
			haystack:       "abcdefghi",
			expectedResult: -1,
			expectedError:  "needle not found",
		},
		{
			name:           "Empty Needle",
			needle:         "",
			haystack:       "abcdefghi",
			expectedResult: -1,
			expectedError:  "needle can not be empty",
		},
		{
			name:           "Empty Haystack",
			needle:         "abc",
			haystack:       "",
			expectedResult: -1,
			expectedError:  "haystack can not be empty",
		},
		{
			name:           "Needle Longer Than Haystack",
			needle:         "abcdefgh",
			haystack:       "abc",
			expectedResult: -1,
			expectedError:  "haystack can not be longer than needle",
		},
		{
			name:           "Overlapping Matches",
			needle:         "aba",
			haystack:       "abababa",
			expectedResult: 0,
			expectedError:  "",
		},
		{
			name:           "Single Character Match",
			needle:         "a",
			haystack:       "abcdef",
			expectedResult: 0,
			expectedError:  "",
		},
		{
			name:           "Single Character No Match",
			needle:         "z",
			haystack:       "abcdef",
			expectedResult: -1,
			expectedError:  "needle not found",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result, errorMessage := algorithms.Search([]rune(testCase.needle), []rune(testCase.haystack))
			if result != testCase.expectedResult {
				t.Errorf("Search(%q, %q) = result %d; want %d", testCase.needle, testCase.haystack, result, testCase.expectedResult)
			}
			if errorMessage == nil && testCase.expectedError != "" {
				t.Errorf("Search(%q, %q) = error(%v); want nil", testCase.needle, testCase.haystack, errorMessage.Error())
			}
			if (errorMessage != nil) && (errorMessage.Error() != testCase.expectedError) {
				t.Errorf("Search(%q, %q) = error(%v); want %v", testCase.needle, testCase.haystack, errorMessage, testCase.expectedError)
			}
		})
	}
}
