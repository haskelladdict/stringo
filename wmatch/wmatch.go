// wmatch looks for the given word in the given input file.
// The input file is assumed to be in plain text format. An optional
// flag specifies the number of context characters to add to the output
// before and after the match (default = 10).
package main

import (
  "flag"
  "fmt"
  "io/ioutil"
  "log"
  "os"
  sp "github.com/haskelladdict/stringo/string_processing"
)



// usage displays standard usage information for this file
func usage() {
  fmt.Printf("usage: %s [-c num_context] <word> <filename>\n", os.Args[0])
  os.Exit(1)
}



// command line parsing stuff
var num_context int
const (
  context_usage = "number of lines of context"
  context_default = 10
)

var exact_match bool
const (
  exact_match_usage = "number of lines of context"
  exact_match_default = false
)

var colored_output bool
const (
  colored_output_usage = "colored output"
  colored_output_default = true
)


func init() {
  flag.IntVar(&num_context, "n", context_default,
    context_usage + " (short)")
  flag.IntVar(&num_context, "num", context_default, context_usage)

  flag.BoolVar(&exact_match, "e", exact_match_default,
    exact_match_usage + " (short)")
  flag.BoolVar(&exact_match, "exact match", exact_match_default,
    exact_match_usage)

  flag.BoolVar(&colored_output, "c", colored_output_default,
    colored_output_usage + " (short)")
  flag.BoolVar(&colored_output, "color", colored_output_default,
    colored_output_usage)
}


// get_matches returns all matches for word filtered according
// to the provided command line flags
func get_matches(input, word string) []string {

  // process, create prefix array and sort
  content := sp.Preprocess_string(string(input))
  prefix_array := sp.Make_prefix_array(content)
  sp.String_quicksort(prefix_array)

  matches := make([]string, 0)
  length := len(content)
  i := sp.Binary_search(prefix_array, word)

  // found a match
  if i >= 0 {

    raw_matches := sp.Scan_array(prefix_array, word, i)
    if exact_match {
      for _, match := range raw_matches {
        match_index := length - len(match)

        // throw out matches starting in the middle of a word 
        if match_index != 0 && content[match_index-1] != ' ' {
          continue
        } else if match_index <= length-2 &&
            content[match_index+len(word)] != ' ' {
          continue
        } else {
          matches = append(matches, match)
        }
      }
    } else {
      matches = raw_matches
    }
  }

  return matches
}



// main entry point
func main() {

  flag.Parse()

  if len(flag.Args()) < 2 {
    usage()
  }
  word := flag.Args()[0]

  // open file and read it into string
  raw_content, err := ioutil.ReadFile(flag.Args()[1])
  if err != nil {
    log.Fatal("Failed to open and read file.")
  }

  content := sp.Preprocess_string(string(raw_content))
  matches := get_matches(string(raw_content), word)
  for _, match := range matches {

    match_index := len(content) - len(match)
    if colored_output {
      fmt.Print(content[match_index-num_context:match_index])
      fmt.Print("\033[1;31m", 
        content[match_index:match_index+len(word)], "\033[0m")
      fmt.Println(content[match_index+len(word):len(word)+match_index+num_context])
    } else {
      fmt.Println(content[match_index-num_context:len(word)+match_index+num_context])
    }
  }
}

