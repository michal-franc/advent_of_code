package main

import "fmt"
import "io/ioutil"
import "strings"
import "strconv"
import "math"

func Btoi(r byte) int {
  return int(r) - '0'
}

func Rtoi(r rune) int {
  return int(r) - '0'
}

func FindEvenDivision(rawdata string) int {
  values := strings.Fields(rawdata)
  if len(values) <= 0 {
    return 0
  }

  for ii, i := range values {
    current, _ := strconv.Atoi(i)
    for ji, j := range values {
      next, _ := strconv.Atoi(j)
      if ii != ji {
        if math.Mod(float64(current), float64(next)) == 0 {
          if next > current {
            return next / current
          } else {
            return current / next
          }
        }
      }
    }
  }

  return -1
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
