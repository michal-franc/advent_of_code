package main

import "testing"

func TestMinMax(t *testing.T) {

  testCases := map[string] int {
    "44 2 4": 42,
    "11 1 1": 10,
    "9 12 12 1 0 29": 29,
    "1 2 3 4": 3,
    "1 12 2": 11,
    "1 1 2": 1,
  }

  for key, val := range testCases {
  actual := FindMixMaxDiff(key)
  expected := val
  if actual != expected {
    t.Errorf("For %v expected %v but got %v", key, expected, actual)
    }
  }
}
