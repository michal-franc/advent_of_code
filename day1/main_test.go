package main

import "testing"

//1122 = 3
func simpleSum(a []int) int {
  sum := 0

  for i := range a {
    sum = sum + i
  }
  return sum
}

func TestBasic3(t *testing.T) {
  actual := simpleSum([]int {1,1,1,1})
  expected := 4
  if ( actual != expected) {
    t.Errorf("For 1111 expected %v but got %v", expected, actual)
  }
}

func TestBasic2(t *testing.T) {
  actual := simpleSum([]int {9,1,2,1,2,1,2,9})
  expected := 9
  if ( actual != expected) {
    t.Errorf("For 91212129 expected %v but got %v", expected, actual)
  }
}

func TestBasic1(t *testing.T) {
  actual := simpleSum([]int {1,2,3,4})
  expected := 0
  if ( actual != expected) {
    t.Errorf("For 1234 expected %v but got %v", expected, actual)
  }
}

func TestBasic(t *testing.T) {
  actual := simpleSum([]int {1,1,2,2})
  expected := 3
  if ( actual != expected) {
    t.Errorf("For 1122 expected %v but got %v", expected, actual)
  }
}
