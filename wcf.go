package wcf

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"
)

// Word is a data structure that will be used to compute top k or top lowest k
type Word struct {
	Key   string
	Count int
}

// CountLines counts the number of lines in the supplied input. It returns a valid number of lines or an error.
func CountLines(rd io.Reader) (int, error) {
	lc := 0
	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		if scanner.Text() != "" {
			lc++
		}
	}
	return lc, scanner.Err()
}

// CountWords counts the number of words and their frequency in the supplied input.
// It returns a valid number of words and a map representing the word frequency.
// @TODO, there is a bug, if a word is terminated by ..., it won't be counted. To fix.
func CountWords(rd io.Reader) (int, map[string]int, error) {
	scanner := bufio.NewScanner(rd)
	scanner.Split(bufio.ScanWords)
	wc := 0
	wf := make(map[string]int)
	for scanner.Scan() {
		word := scanner.Text()
		lastChar, size := utf8.DecodeLastRuneInString(word)
		if !unicode.IsDigit(lastChar) && !unicode.IsLetter(lastChar) {
			word = word[:len(word)-size]
		}
		if word != "" && word != ".." {
			wf[word]++
			wc++
		}
	}
	return wc, wf, scanner.Err()
}

// TopK produces the top k word by their frequency
func TopK(input map[string]int, k int) ([]Word, error) {
	if k > len(input) {
		return []Word{}, fmt.Errorf("k is greater than the total number of elements")
	}
	h := &MinHeap{}
	ans := make([]Word, k)
	for key, count := range input {
		w := Word{key, count}
		heap.Push(h, w)
		if h.Len() > k {
			// remove the smallest
			heap.Pop(h)
		}
	}
	// build the answer
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(Word)
	}
	return ans, nil
}
