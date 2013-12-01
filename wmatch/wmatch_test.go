// wmatch_test implements a few small tests and benchmarks for the
// word match code

package main


import (
  "io/ioutil"
  "log"
  "testing"
)


// test and benchmark
var input_string string
func init() {
  input_bytes, err := ioutil.ReadFile("test_files/moby_dick.txt")
  if err != nil {
    log.Fatal("Failed to open and read file.")
  }

  input_string = string(input_bytes)
}



// TestWmatch_1 tests the proper number of occurences of string "whale"
func TestWmatch_1(t *testing.T) {

  target_length := 1263
  target_word := "whale"

  matches := get_matches(input_string, target_word)
  if len(matches) != target_length {
    t.Errorf("Incorrect number of matches for target %s", target_word)
  }
}


// TestWmatch_2 tests the proper number of exact occurences of string "whale"
func TestWmatch_2(t *testing.T) {

  target_length := 371
  target_word := "whale"

  exact_match = true
  matches := get_matches(input_string, target_word)
  exact_match = false
  if len(matches) != target_length {
    t.Errorf("Incorrect number of exact matches for target %s", target_word)
  }
}



// TestWmatch_3 tests the proper number of occurences of string "whale"
func TestWmatch_3(t *testing.T) {

  target_length := 246
  target_word := "harpoon"

  matches := get_matches(input_string, target_word)
  if len(matches) != target_length {
    t.Errorf("Incorrect number of matches for target %s", target_word)
  }
}



// TestWmatch_4 tests the proper number of exact occurences of string "whale"
func TestWmatch_4(t *testing.T) {

  target_length := 39
  target_word := "harpoon"

  exact_match = true
  matches := get_matches(input_string, target_word)
  exact_match = false
  if len(matches) != target_length {
    t.Errorf("Incorrect number of exact matches for target %s", target_word)
  }
}


// BenchmarkWmatch_1 benchmarks the computation of the lrs
func BenchmarkWmatch_1(t *testing.B) {

  target_word := "whale"
  get_matches(input_string, target_word)
}
