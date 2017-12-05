package main

import "testing"

func TestMinMax(t *testing.T) {

  testCases := map[string] int {
    "44 3 4": 11,
    "11 1 1": 11,
    "9 12 12 5 7 29": 1,
    "7 2 3 4": 2,
    "5 12 2": 6,
    "11 4 2": 2,
    "50 3 7 100": 2,
  }

  for key, val := range testCases {
  actual := FindEvenDivision(key)
  expected := val
  if actual != expected {
    t.Errorf("For %v expected %v but got %v", key, expected, actual)
    }
  }
}
