package counter

import (
	"fmt"
	"testing"
)

func TestByteCounter(t *testing.T) {
	var c ByteCounter

	n, err := c.Write([]byte("hello"))
	if err != nil {
		t.Fail()
	}
	if n != len("hello") {
		t.Fail()
	}
	t.Logf("Total Bytes: %v", c)

	fmt.Fprintf(&c, ",%s", "world")
	t.Logf("Total Bytes: %v", c)
}

func TestWordCounter(t *testing.T) {
	var c WordCounter

	n, err := c.Write([]byte("This is a word counter test."))
	if err != nil {
		t.Fail()
	}
	if n != 6 {
		t.Fail()
	}

	t.Logf("Total Words: %v", c)

	fmt.Fprintf(&c, "%s", "Keep simple, keep stupid, then you will be freedom!")
	t.Logf("Total Words: %v", c)
}

func TestLineCounter(t *testing.T) {
	var c LineCounter

	n, err := c.Write([]byte("This is a line counter test.\nCan you guess it?\nhaha~~~"))
	if err != nil {
		t.Fail()
	}
	if n != 3 {
		t.Fail()
	}

	t.Logf("Total Lines: %v", c)

	fmt.Fprintf(&c, "%s", "Keep simple, keep stupid, then you will be freedom!\nok, it's over.")
	t.Logf("Total Lines: %v", c)
}
