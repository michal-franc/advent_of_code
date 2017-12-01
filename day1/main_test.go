package main

import "testing"

func TestBasic5(t *testing.T) {
  actual := simpleSum("4424")
  expected := 8
  if ( actual != expected) {
    t.Errorf("For 4424 expected %v but got %v", expected, actual)
  }
}

func TestBasic3(t *testing.T) {
  actual := simpleSum("1111")
  expected := 4
  if ( actual != expected) {
    t.Errorf("For 1111 expected %v but got %v", expected, actual)
  }
}

func TestBasic2(t *testing.T) {
  actual := simpleSum("91212129")
  expected := 9
  if ( actual != expected) {
    t.Errorf("For 91212129 expected %v but got %v", expected, actual)
  }
}

func TestBasic1(t *testing.T) {
  actual := simpleSum("1234")
  expected := 0
  if ( actual != expected) {
    t.Errorf("For 1234 expected %v but got %v", expected, actual)
  }
}

func TestBasic(t *testing.T) {
  actual := simpleSum("1122")
  expected := 3
  if ( actual != expected) {
    t.Errorf("For 1122 expected %v but got %v", expected, actual)
  }
}
