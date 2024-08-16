package main

import (
  "errors"
  "fmt"
  "io"
  "log"
  "os"
  "strings"
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
  fmt.Print(sortAsc(data) + "\n")
}

func sortAsc(data []byte) string {
  var tempLine string
  lines := strings.Split(strings.TrimSpace(string(data)), "\n")
  for i := 0; i < len(lines); i++ {
    for x := 0; x < len(lines); x++ {
      if lines[i] < lines[x] {
        tempLine = lines[i]
        lines[i] = lines[x]
        lines[x] = tempLine
      }
    }
  }
  return strings.Join(lines, "\n")
}
