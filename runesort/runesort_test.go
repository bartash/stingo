package runesort

import "testing"
import "fmt"

func TestSimple(t *testing.T) {
	output := SortString("bcad");
	expected := "abcd"
	if output != expected {
		errorString := fmt.Sprintf("Expected %v, got %v", expected, output)
		t.Error(errorString)
	}
}
