package main

import "fmt"
import "math"

func main()  {
    target := 600851475143
    f := float64(target)
    var i float64

    for target%2 == 0 {
        fmt.Println("2")
        target = target/2
    }

    for i=3; i<=math.Sqrt(f); i++ {
        for target%int(i) == 0 {
            fmt.Println(i)
            target = target/int(i)
            i+=2
        }
    }

    if target>2 {
        fmt.Println(target)
    }
}