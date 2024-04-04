package main

import (
	"testing"
)

func TestCheckRobots(t *testing.T) {
	testCases := []struct {
		domain string
		output string
	}{
		{"www.nytimes.com", "www.nytimes.com,false,true"},
		{"facebook.com", "facebook.com,true,true"},
		{"google.com", "google.com,false,false"},
		{"adfjahsdfasdkfjhawwwwww.com", "adfjahsdfasdkfjhawwwwww.com,Unable to retrieve robots.txt. Maybe the domain doesn't exist?,"},
	}
	for _, tc := range testCases {
		results := checkRobots(tc.domain)

		if len(results) < 1 {
			t.Fatal("Expected at least one line of output, got none")
		}

		if results[0] != tc.output {
			t.Fatalf("Expected output to be %s, got: %s", tc.output, results[0])
		}
	}
}
