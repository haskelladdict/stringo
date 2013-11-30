// lrs_test implements a few small tests and benchmarks for the
// lowest repeated substring code

package main


import (
  "io/ioutil"
  "log"
  "testing"
)


// test and benchmark 1
var input_string_1 string
func init() {
  input_bytes_1, err := ioutil.ReadFile("test_files/moby_dick.txt")
  if err != nil {
    log.Fatal("Failed to open and read file.")
  }

  input_string_1 = string(input_bytes_1)
}



// TestLrs_1 tests the proper determination of the longest repeated
// substring
func TestLrs_1(t *testing.T) {

  target_string := ",--  Such a funny, sporty, gamy, jesty, joky, hoky-poky lad, is the Ocean, oh!  Th"
  target_length := 82

  longest_string := lrs_h(input_string_1)

  if longest_string != target_string || len(longest_string) != target_length {
    t.Errorf("Failed to identify correct longest repeated string")
  }
}


// BenchmarkLrs_1 benchmarks the computation of the lrs
func BenchmarkLrs_1(t *testing.B) {
  lrs_h(input_string_1)
}




// test 2
var input_string_2 string
func init() {
  input_bytes_2, err := ioutil.ReadFile("test_files/short_test.txt")
  if err != nil {
    log.Fatal("Failed to open and read file.")
  }

  input_string_2 = string(input_bytes_2)
}



// TestLrs_2 tests the proper determination of the longest repeated
// substring
func TestLrs_2(t *testing.T) {

  target_string := " have never had a "
  target_length := 18

  longest_string := lrs_h(input_string_2)

  if longest_string != target_string || len(longest_string) != target_length {
    t.Errorf("Failed to identify correct longest repeated string")
  }
}



// BenchmarkLrs_2 benchmarks the computation of the lrs
func BenchmarkLrs_2(t *testing.B) {
  lrs_h(input_string_2)
}


