package randname

import (
	"fmt"
	"testing"
)

func TestRandName(t *testing.T) {
	name := GetName()
	fmt.Println(name)
}
