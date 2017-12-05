package main

import "fmt"
import "io/ioutil"
import "strings"
import "strconv"

func Btoi(r byte) int {
  return int(r) - '0'
}

func Rtoi(r rune) int {
  return int(r) - '0'
}

func FindMinMaxDiff(rawdata string) int {
  values := strings.Fields(rawdata)
  if len(values) <= 0 {
    return 0
  }

  min, _ := strconv.Atoi(values[0])
  max, _ := strconv.Atoi(values[0])

  for _, x := range values {
    val, _ := strconv.Atoi(x)
    if val > max {
      max = val
    }

    if val < min {
      min = val
    }
  }

  return max - min
}

func main() {

  data, _ := ioutil.ReadFile("input.dat")
  lines:= strings.Split(string(data), "\n")

  result := 0

  for _, line := range lines {
    result += FindMinMaxDiff(line)
  }

  fmt.Println(result)
}
