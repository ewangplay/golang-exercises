/*
Package counter implements the exercise below.

Exercise 7.1: Using the ideas from ByteCounter, implement counters for words
and for lines. You will find bufio.ScanWords useful.
*/
package counter

import (
	"bufio"
	"bytes"
)

// ByteCounter counts the bytes number of input data
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	l := len(p)
	*c += ByteCounter(l)
	return l, nil
}

// WordCounter counts the words number of input data
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	*c += WordCounter(count)
	return count, nil
}

// LineCounter counts the words number of input data
type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	*c += LineCounter(count)
	return count, nil
}
