package main

import "testing"

func TestBasic(t *testing.T) {

  testCases := map[string]int {
    "1212": 6,
    "1221": 0,
    "123425": 4,
    "123123": 12,
    "12131415":4,
  }

  for key, val := range testCases {
  actual := simpleSum(key)
  expected := val
  if actual != expected {
    t.Errorf("For %v expected %v but got %v", key, expected, actual)
    }
  }
}
