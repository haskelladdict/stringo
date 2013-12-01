// lrs computes the longest repeating substring off a plain text file provided
// on the command line
package string_processing

import (
  "math"
  "regexp"
)



// Preprocess_string removes newlines and extraneous spaces from the input 
// string
func Preprocess_string(content string) string {

  // strip BOM if present
  re := regexp.MustCompile("\xEF\xBB\xBF")
  content = re.ReplaceAllString(content, "")
  re = regexp.MustCompile("([\n]+|[\r\n]+|[ ]{2,})")
  return re.ReplaceAllString(content, " ")
}



// Make_prefix_array creates a prefix array of the given string
func Make_prefix_array(content string) []string {

  prefix_array := make([]string, len(content))
  for i := 0; i < len(content); i++ {
    prefix_array[i] = content[i:]
  }
  return prefix_array
}



// Char_at is a helper function for string sorting. It returns the
// pos' character of byte slice item as int or -1 of we've have reaced
// the end of the slice
func Char_at(item string, pos int) int {
  if len(item) <= pos {
    return -1
  } else {
    return int(item[pos])
  }
}



// msd_radix_sort implements a most significant digit radix sort
// algorithm
func Msd_radix_sort(items []string) {

  aux := make([]string, len(items))
  msd_radix_sort_h(items, aux, 0, len(items)-1, 0)
}



// msd_radix_sort_h implements the actual recursive code used for 
// implementing msd_radix sort
func msd_radix_sort_h(items []string, aux []string, low int, high int,
  d int) {

  if low >= high {
    return
  }

  R := math.MaxUint8
  count := make([]int, R+2)

  // count keys
  for i := low; i <= high; i++ {
    count[Char_at(items[i],d)+2]++
  }

  // build cummulative counts
  for i := 0; i <= R; i++ {
    count[i+1] += count[i]
  }

  // build aux array sorted based on dth digit
  for i := low; i <= high; i++ {
    aux[count[Char_at(items[i],d)+1]] = items[i]
    count[Char_at(items[i],d)+1]++
  }

  // copy aux into a
  for i := low; i <= high; i++ {
    items[i] = aux[i-low]
  }

  // recursively search subarrays
  for i := 0; i < R; i++ {
    msd_radix_sort_h(items, aux, low+count[i], low+count[i+1]-1, d+1)
  }
}



// string_quicksort implements a 3-way string quicksort for efficient
// sorting of an array of strings
func String_quicksort(items []string) {
  string_quicksort_h(items, 0, len(items)-1, 0)
}



// string_quicksort_h is a recursive helper function implementing
// string_quicksort
func string_quicksort_h(items []string, low int, high int, pos int) {
  if low >= high {
    return
  }

  pivot := Char_at(items[low], pos)
  i := low + 1
  lt := low
  gt := high

  // partition elements around pivot
  for i <= gt {
    comp := Char_at(items[i], pos)
    if comp < pivot {
      items[lt], items[i] = items[i], items[lt]
      lt++
      i++
    } else if comp > pivot {
      items[i], items[gt] = items[gt], items[i]
      gt--
    } else {
      i++
    }
  }

  string_quicksort_h(items, low, lt-1, pos)
  if (pivot >= 0) {
    string_quicksort_h(items, lt, gt, pos+1)
  }
  string_quicksort_h(items, gt+1, high, pos)
}



// Binary_search searches for an entry in the sorted array items 
// with word as the prefix.
// XXX: items has to be sorted
func Binary_search(items []string, word string) int {

  low := 0
  high := len(items)
  word_len := len(word)

  for low <= high {

    middle := low + (high-low)/2
    item := string(items[middle][:word_len])
    if item == word {
      return middle
    } else if item < word {
      low = middle + 1
    } else {
      high = middle - 1
    }
  }

  return -1
}



// is_prefix checks if word is a prefix of target
func is_prefix(word, target string) bool {

  if len(target) < len(word) {
    return false
  }

  item := target[:len(word)]
  if word == item {
    return true
  } else {
    return false
  }
}



// Scan_array scans a sorted array of strings for occurences
// of word as prefix before and after the index loc. 
// The string at loc is assumed to contain word as prefix.
func Scan_array(items []string, word string, loc int) []string {

  matches := make([]string, 0)

  // scan backward
  for i := loc; i > 0; i-- {
    if is_prefix(word, items[i]) {
      matches = append(matches, items[i])
    } else {
      break
    }
  }

  // scan forward
  for i := loc+1; i < len(items); i++ {
    if is_prefix(word, items[i]) {
      matches = append(matches, items[i])
    } else {
      break
    }
  }

  return matches
}
