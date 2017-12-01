package main

import "fmt"
import "io/ioutil"

func Btoi(r byte) int {
  return int(r) - '0'
}

func Rtoi(r rune) int {
  return int(r) - '0'
}

func simpleSum(values string) int {
  sum := 0
  current := -1
  first := Btoi(values[0])
  last := Btoi(values[len(values)-1])

  if(first == last) {
    sum += first
  }

  for _, num := range values {

    numI := Rtoi(num)

    if(current == numI) {
      sum += numI
    }

    current = numI
  }

  return sum
}

func main() {

  dat, _ := ioutil.ReadFile("input.dat")

  sum := simpleSum(string(dat))
  fmt.Println(sum)
}
