package main

import "fmt"
import "io/ioutil"

func Btoi(r byte) int {
  return int(r) - '0'
}

func Rtoi(r rune) int {
  return int(r) - '0'
}

func checkBoundOr0(values string) int {
  first := Btoi(values[0])
  last := Btoi(values[len(values)-1])

  if(first == last) {
    return first
  }

  return 0
}

func simpleSum(values string) int {
  sum := checkBoundOr0(values)

  last := -1
  for _, num := range values {

    current := Rtoi(num)

    if(current == last) {
      sum += current
    }

    last = current
  }

  return sum
}

func main() {

  dat, _ := ioutil.ReadFile("input.dat")

  sum := simpleSum(string(dat))
  fmt.Println(sum)
}
