package functions

import (
  "os"
  "slices"
  "strings"
)

func SortAsc(data []byte) string {
  lines := strings.Split(strings.TrimSpace(string(data)), "\n")
  slices.Sort(lines)
  return strings.Join(lines, "\n")
}

func SortDesc(data []byte) string {
  lines := strings.Split(strings.TrimSpace(string(data)), "\n")
  slices.Sort(lines)
  slices.Reverse(lines)
  return strings.Join(lines, "\n")
}

func AvoidParam(prefix string) func(string) bool {
  return func(a string) bool {
    return a != os.Args[0] && !strings.HasPrefix(a, prefix)
  }
}

func FindParam(prefix string) func(string) bool {
  return func(a string) bool {
    return a != os.Args[0] && strings.HasPrefix(a, prefix)
  }
}
