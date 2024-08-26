package main

import (
  "errors"
  "flag"
  "fmt"
  "io"
  "log"
  "os"

  "github.com/hector-vido/go-sort/internal/functions"
)

var orderFlag bool

func init() {
  flag.BoolVar(&orderFlag, "r", false, "Reverse the sort order")
}

func main() {

  var data []byte
  var err error
  var f *os.File
  
  flag.Parse()

  var nArg = flag.NArg()

  if nArg == 1 {

    fileName := flag.Arg(0)
    if _, err = os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
      log.Fatal(err)
    }

    if f, err = os.Open(fileName); err != nil {
      log.Fatal(err)
    }
    defer f.Close()

    if data, err = io.ReadAll(f); err != nil {
      log.Fatal(err)
    }

  } else if nArg > 1 {
      log.Fatal(fmt.Sprintf("Unrecognized argument \"%s\"", flag.Arg(1)))
  } else {
    data, err = io.ReadAll(os.Stdin)
    if err != nil {
      log.Fatal(err)
    }
  }

  if orderFlag {
    fmt.Printf("%s\n", functions.SortDesc(data))
  } else {
    fmt.Printf("%s\n", functions.SortAsc(data))
  }
}
