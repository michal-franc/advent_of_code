package main

import "fmt"
import "io/ioutil"
import "strings"
import "strconv"
import "math"

func modulo(x int, y int) float64 {
  return math.Mod(float64(x), float64(y))
}

func divide(x int, y int) int {
  if x > y {
    return x / y
  } else {
    return y / x
  }
}

func FindEvenDivision(rawdata string) int {
  values := strings.Fields(rawdata)
  if len(values) <= 0 {
    return 0
  }

  result := -1 //default value

  for ii, i := range values {
    current, _ := strconv.Atoi(i)
    for ji, j := range values {
      next, _ := strconv.Atoi(j)
      if ii != ji {
        if modulo(next, current) == 0 {
          result = divide(next, current)
          break;
        }
      }
    }
  }

  return result
}

func main() {

  data, _ := ioutil.ReadFile("input.dat")
  lines:= strings.Split(string(data), "\n")

  result := 0

  for _, line := range lines {
    result += FindEvenDivision(line)
  }

  fmt.Println(result)
}
