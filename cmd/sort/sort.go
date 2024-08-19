package main

import (
  "errors"
  "fmt"
  "io"
  "log"
  "os"
  "slices"

  "github.com/hector-vido/go-sort/internal/functions"
)

func main() {

  var data []byte;
  filePathIdx := slices.IndexFunc(os.Args, functions.AvoidParam("-"))
  if filePathIdx != -1 {

    invalidParameterIdx := slices.IndexFunc(os.Args[filePathIdx + 1:], functions.AvoidParam("-"))
    if invalidParameterIdx != -1 {
      log.Fatal(fmt.Sprintf("Unrecognized argument \"%s\"", os.Args[invalidParameterIdx + filePathIdx + 1]))
    }

    fileName := os.Args[filePathIdx]
    _, err := os.Stat(fileName)
    if errors.Is(err, os.ErrNotExist) {
      log.Fatal(err)
    }
    f, err := os.Open(fileName)
    if err != nil {
      log.Fatal(err)
    }
    defer f.Close()
    data, err = io.ReadAll(f)
    if err != nil {
      log.Fatal(err)
    }
  }

  if slices.IndexFunc(os.Args, functions.FindParam("-r")) != -1 {
    fmt.Printf("%s\n", functions.SortDesc(data))
  } else {
    fmt.Printf("%s\n", functions.SortAsc(data))
  }
}
