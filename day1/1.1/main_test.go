package main

import "testing"

func TestBasic(t *testing.T) {

  testCases := map[string]int {
    "4424": 8,
    "1111": 4,
    "91212129": 9,
    "1234": 0,
    "1122": 3,
    "112": 1,
  }

  for key, val := range testCases {
  actual := simpleSum(key)
  expected := val
  if actual != expected {
    t.Errorf("For %v expected %v but got %v", key, expected, actual)
    }
  }
}
