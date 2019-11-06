package main

import (
	"testing"
)

func TestCalculateMacComb(t *testing.T) {
	for _, testCase := range testCases {
		if num := calculateResult(testCase.input); num != testCase.expected {
			t.Fatalf("FAIL: %s(%s)\nExpected: %d\nActual: %d",
				testCase.description, testCase.input, testCase.expected, num)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}
