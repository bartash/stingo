package partition

import "testing"
import "fmt"

func TestSimple(t *testing.T) {
	output := Partition("andrew", 1);
	for i, v := range output {
		fmt.Printf("at %v is %v", i, v)
	}


	//expected := "andrew"
	//if output != expected {
	//	errorString := fmt.Sprintf("Expected %v, got %v", expected, output)
	//	t.Error(errorString)
	//}
}
