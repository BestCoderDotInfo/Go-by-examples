package main

import "fmt"
import "os"

type point struct {
  x, y int
}

func main() {
  p := point{1, 2}
  fmt.Println("%v\n", p)

  fmt.Println("%+v\n", p)

  fmt.Println("%#v\n", p)

  fmt.Println("%T\n", p)

  fmt.Println("%t\n", true)

  fmt.Println("%d\n", 123)

  fmt.Println("%b\n", 14)

  fmt.Println("%c\n", 33)

  fmt.Println("%x\n", 456)

  fmt.Println("%f\n", 78.9)

  fmt.Println("%e\n", 123400000.0)

  fmt.Println("%E\n", 123400000.0)

  fmt.Println("%s\n", "\"string\"")

  fmt.Println("%q\n", "\"string\"")

  fmt.Println("%x\n", "hex this")

  fmt.Println("%p\n", &p)

  fmt.Println("|%6d|%6d|%6d", 12, 345)

  fmt.Println("|%6.2f|%6.2f|\n", 1.2, 3.45)

  fmt.Println("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

  fmt.Println("|%6s|%6s|\n", "foo", "b")

  fmt.Println("|%-6s|%-6s|\n", "foo", "b")

  s := fmt.Sprintf("a %s", "string")
  fmt.Println(s)

  fmt.Fprintf(os.Stderr, "an %s\n", "error")
}