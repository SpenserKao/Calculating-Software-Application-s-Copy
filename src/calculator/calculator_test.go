package main

//!+bench

import (
	"testing"
)

//!-bench

//!+test
func TestValidRec(t *testing.T) {
	var tests = []struct {
		computerID    string
		userID        string
		applicationID string
		computerType  string
		want          bool
	}{
		{"1", "1", "374", "LAPTOP", true},
		{"1", "1", "374", "", false},
		{"2", "1", "", "LAPTOP", false},
		{"3", "", "374", "DESKOP", false},
		{"", "2", "", "DESKTOP", false},
		{"1", " ", "", "DESKTOP", false},
	}
	for _, test := range tests {
		if got := ValidRec(test.computerID, test.userID, test.applicationID, test.computerType); got != test.want {
			t.Errorf("ValidRec(%q, %q, %q, %q) = %v", test.computerID, test.userID, test.applicationID, test.computerType, got)
		}
	}
}
