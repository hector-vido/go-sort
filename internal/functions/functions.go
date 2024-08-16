package functions

import (
  "strings"
)

func SortAsc(data []byte) string {
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
