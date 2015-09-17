package fibo

import "testing"

type fibtest struct {
  input int
  expected_o int
}

var fts = []fibtest {
  {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13},
}

func TestFib(t *testing.T) {
  var actual int
  for _, tc := range fts {
    actual = Fib(tc.input);
    if actual != tc.expected_o {
      t.Errorf("Fib (%d) = Expected (%d) ----- Actual (%d)",tc.input,tc.expected_o,actual);

    }
  }
}
