package partition

import "testing"
import "fmt"

func TestSimple(t *testing.T) {
	output := Partition("andrew", 1);
	for i, v := range output {
		fmt.Printf("at %v is %v", i, v)
	}
	var expected = []Pair  {
		Pair{"a", "dnrew"},
		Pair{"n", "adrew"},
		Pair{"d", "anrew"},
		Pair{"r", "adnew"},
		Pair{"e", "adnrw"},
		Pair{"w", "adnre"},
	}
	 if len(expected) != len(output){
		 t.Errorf("lengths differ")
	 }
	//expected := "andrew"
	//if output != expected {
	//	errorString := fmt.Sprintf("Expected %v, got %v", expected, output)
	//	t.Error(errorString)
	//}
}
