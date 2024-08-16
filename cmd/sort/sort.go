package main

import (
  "errors"
  "fmt"
  "io"
  "log"
  "os"
  "strings"

  "github.com/hector-vido/go-sort/internal/functions"
)

func main() {
  var fileName string
  argsLen := len(os.Args)
  for i := 1; i < argsLen; i++ {
    if !strings.HasPrefix(os.Args[i], "-") && fileName == "" {
      fileName = os.Args[i]
      _, err := os.Stat(fileName)
      if errors.Is(err, os.ErrNotExist) {
        log.Fatal(err)
      }
    } else {
      log.Fatal(fmt.Sprintf("Unrecognized argument \"%s\"", os.Args[i]))
    }
  }
  f, err := os.Open(fileName)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()
  data, err := io.ReadAll(f)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", functions.SortAsc(data))
}
