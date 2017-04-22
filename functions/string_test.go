package functions

import (
	"bytes"
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	placeHolder := "[Error : %s (%s) option is mandatory with %s command]"
	finalValue := "[Error : a (b) option is mandatory with c command]"
	buf := bytes.NewBufferString("")

	fmt.Fprintf(buf, placeHolder, "a", "b", "c")
	if finalValue != buf.String() {
		t.Errorf("Unable to match the final value : %s", buf.String())
	}

}
