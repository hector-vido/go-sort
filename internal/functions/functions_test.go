package functions

import (
  "os"
  "slices"
  "testing"
)

// TestParam calls functions.AvoidParam with some args
// to verify if a filename is present
func TestAvoidParam(t *testing.T) {
  testCases := []struct {
    name, a, b  string
    expected int
  } {
    {
      name: "Find filename.txt",
      a: "-r",
      b: "filename.txt",
      expected: 2,
    },
    {
      name: "Find nothing",
      a: "-r",
      b: "-c",
      expected: -1,
    },
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      os.Args[1] = tc.a
      os.Args[2] = tc.b
      result := slices.IndexFunc(os.Args, AvoidParam("-"))
      if result != tc.expected  {
        t.Errorf("expecting %d, got %d", tc.expected, result)
      }
    })
  }
}

// TestFindParam calls functons.FindParam with some args
// and expects to find these parameters
func TestFindParam(t *testing.T) {
  testCases := []struct {
    name, a, b  string
    expected int
  } {
    {
      name: "Find -r",
      a: "filename.txt",
      b: "-r",
      expected: 2,
    },
  }
  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      os.Args[1] = tc.a
      os.Args[2] = tc.b
      result := slices.IndexFunc(os.Args, FindParam("-r"))
      if result != tc.expected  {
        t.Errorf("expecting %d, got %d", tc.expected, result)
      }
    })
  }
}

// TestSort calls SortAsc("d\nc\na\nb") and SortDesc("d\nc\na\nb")
// expecting to receive the parameters ordered.
func TestSort(t *testing.T) {
  testCases := []struct {
    name, expected string
    lines []byte
    function func(lines []byte) string
  } {
    {
      name: "SortAsc",
      lines: []byte("d\nc\na\nb"),
      expected: "a\nb\nc\nd",
      function: func(lines []byte) string { return SortAsc(lines) },
    },
    {
      name: "SortDesc",
      lines: []byte("d\nc\na\nb"),
      expected: "d\nc\nb\na",
      function: func(lines []byte) string { return SortDesc(lines) },
    },
  }
  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      result := tc.function(tc.lines)
      if result != tc.expected  {
        t.Errorf("expecting %s, got %s", tc.expected, result)
      }
    })
  }
}
