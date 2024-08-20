package functions

import (
  "os"
  "slices"
  "strings"
  "testing"
)

// TestAvoidParamFile calls functions.AvoidParam with "-r filename.txt" args
// and expects do find "filename.txt" in position 2.
func TestAvoidParamFile(t *testing.T) {
  os.Args[1] = "-r"
  os.Args[2] = "filename.txt"
  filePathIdx := slices.IndexFunc(os.Args, AvoidParam("-"))
  if filePathIdx == -1 {
    t.Fatalf(`"sort -r filename.txt" = AvoidParam() looking for "filename.txt" but not found.`)
  }
}

// TestAvoidParamStdin calls functions.AvoidParam with "-r -c" args
// and expects to find no files.
func TestAvoidParamStdin(t *testing.T) {
  os.Args[1] = "-r"
  os.Args[2] = "-c"
  noIdx := slices.IndexFunc(os.Args, AvoidParam("-"))
  if noIdx != -1 {
    t.Fatalf(`"sort -r -c" = AvoidParam() should found nothing, but found something.`)
  }
}

// TestFindParam calls functons.FindParam with "filename.txt -r" args
// and expectos to find "-r" in position 2.
func TestFindParam(t *testing.T) {
  os.Args[1] = "filename.txt"
  os.Args[2] = "-r"
  parameterIdx := slices.IndexFunc(os.Args, FindParam("-r"))
  if parameterIdx == -1 {
    t.Fatalf(`"sort filename.txt -r" = FindParam() should found "-r", but found none.`)
  }
}

// TestSortAcs calls functions.SortAsc("d\nc\na\nb") and expects
// to receive them ordered from a to d.
func TestSortAsc(t *testing.T) {
  lines := SortAsc([]byte("d\nc\na\nb"))
  if lines != "a\nb\nc\nd" {
    t.Fatalf(`SortAsc("d, c, a, b")" should order as "a, b, c, d" but ordered as "%s"`, strings.ReplaceAll(lines, "\n", ", "))
  }
}

// TestSortDesc calls functions.SortAsc("d\nc\na\nb") and expects
// to receive them ordered from d to a.
func TestSortDesc(t *testing.T) {
  lines := SortDesc([]byte("d\nc\na\nb"))
  if lines != "d\nc\nb\na" {
    t.Fatalf(`SortDesc("d, c, a, b")" should order as "d, c, b, a" but ordered as "%s"`, strings.ReplaceAll(lines, "\n", ", "))
  }
}
