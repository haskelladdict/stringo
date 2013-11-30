// lrs computes the longest repeating substring off a plain text file provided
// on the command line
package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "os"
  sp "github.com/haskelladdict/stringo/string_processing"
)


// usage displays standard usage information for this file
func usage() {
  fmt.Printf("usage: %s <filename>\n", os.Args[0])
  os.Exit(1)
}


// longest_repeated_substring returs the longest repeated substring
// in the list of items. 
// NOTE: items is assumed to be sorted
func longest_repeated_substring(items []string) string {

  var lrs string
  for i := 0; i < len(items)-1; i++ {
    lcp_len := lcp(items[i], items[i+1])
    if lcp_len > len(lrs) {
      lrs = items[i][:lcp_len]
    }
  }

  return lrs
}



// lcp returns the length of the common prefix of the two strings
// item1 and item2
func lcp(item1 string, item2 string) int {
  var i int
  for i = 0; i < len(item1); i++ {
    if i >= len(item2) {
      break;
    }

    if item1[i] != item2[i] {
      break;
    }
  }

  return i
}



// lrs_h is a helper function for computing the longest repeated
// substring
func lrs_h(raw_content string) string {

  content := sp.Preprocess_string(raw_content)
  prefix_array := sp.Make_prefix_array(content)
  sp.String_quicksort(prefix_array)
  return longest_repeated_substring(prefix_array)
}



// main entry point
func main() {

  if len(os.Args) <= 1 {
    usage()
  }

  // open file and read it into string
  raw_content, err := ioutil.ReadFile(os.Args[1])
  if err != nil {
    log.Fatal("Failed to open and read file.")
  }

  longest := lrs_h(string(raw_content))

  // print result
  fmt.Println("lrs length:", len(longest))
  fmt.Println("lrs       :", longest)
}

