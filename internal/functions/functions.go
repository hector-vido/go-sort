package functions

import (
  "slices"
  "strings"
)

func SortAsc(data []byte) string {
  lines := strings.Split(strings.TrimSpace(string(data)), "\n")
  slices.Sort(lines)
  return strings.Join(lines, "\n")
}
