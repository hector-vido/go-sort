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
  var file_name string
  args_len := len(os.Args)
  for i := 1; i < args_len; i++ {
    if !strings.HasPrefix(os.Args[i], "-") && file_name == "" {
      file_name = os.Args[i]
      _, err := os.Stat(file_name)
      if errors.Is(err, os.ErrNotExist) {
        log.Fatal(err)
      }
    } else {
      log.Fatal(fmt.Sprintf("Unrecognized argument \"%s\"", os.Args[i]))
    }
  }
  f, err := os.Open(file_name)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()
  data, err := io.ReadAll(f)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Print(sort_asc(data) + "\n")
}

func sort_asc(data []byte) string {
  var temp_line string
  lines := strings.Split(strings.TrimSpace(string(data)), "\n")
  for i := 0; i < len(lines); i++ {
    for x := 0; x < len(lines); x++ {
      if lines[i] < lines[x] {
        temp_line = lines[i]
        lines[i] = lines[x]
        lines[x] = temp_line
      }
    }
  }
  return strings.Join(lines, "\n")
}
