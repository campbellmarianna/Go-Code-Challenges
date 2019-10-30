package main

    import (
      "fmt"
    )

    func Sqrt(x float64) (float64, int) {
      z := float64(x/2)
      count := 1
      prev_z := float64(0)
      for prev_z != z {
        prev_z = z
        z -= (z*z - x) / (2*z)
        fmt.Println(z)
        count += 1
      }
      return z, count
    }

    func main() {
      fmt.Println(Sqrt(4))
    }