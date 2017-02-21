package main

import (
    "testing"
)

var fenTests = []struct {
  fen      string // input
  expected bool // expected result
}{
  {"8/8/8/8/8/8/8/8 w - - 0 0", true},
  {"4k3/r6B/8/8/8/8/8/K6Q w - - 0 0", true},
  {"4k3/r6Bp/8p/8/8/8/8/K6Q w - - 0 0", false},
  {"pouet", false},
}

func TestIsValidFen(t *testing.T) {
for _, tt := range fenTests {
    actual := isValidFen(tt.fen)
        if actual != tt.expected {
            t.Errorf("isValidFen(%s): expected %t, actual %t", tt.fen, tt.expected, actual)
        }
    }
}

