package main

import "testing"

//instantiating struct type
type normalizeTestCase struct {
	//key keyType
	input string
	want string
}

func TestNormalize(t *testing.T) {
	//array of test cases with the normalizeTestCase structure defined above
	testCases := []normalizeTestCase {
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"1234567892", "1234567892"},
		{"(123)456-7892", "1234567892"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := normalize(tc.input)

			if actual != tc.want {
				t.Errorf("got %s; want %s", actual, tc.want)
			}
		})
	}
}