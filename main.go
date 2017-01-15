package main // import "github.com/camptocamp/duration"

import (
  "fmt"
  "os"
  "time"
)

func main() {
  arg := os.Args[1]
  duration, _ := time.ParseDuration(arg)
  fmt.Println(time.Duration.Seconds(duration))
}
