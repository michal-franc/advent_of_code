package main

import "fmt"
import "io/ioutil"

func Btoi(r byte) int {
  return int(r) - '0'
}

func Rtoi(r rune) int {
  return int(r) - '0'
}

func nextIndex(index int, shift int, length int) int {
  candidate := index + shift
  outOfBound := (length - 1) - candidate
  if outOfBound < 0 {
    return (-outOfBound) - 1
  }

  return candidate
}

func simpleSum(values string) int {
  length := len(values)
  shift := length / 2
  sum := 0

  for index, num := range values {

    current := Rtoi(num)
    tocheck := Btoi(values[nextIndex(index, shift, length)])

    if current == tocheck {
      sum += current
    }
  }

  return sum
}

func main() {

  dat, _ := ioutil.ReadFile("input.dat")

  sum := simpleSum(string(dat))
  fmt.Println(sum)
}
