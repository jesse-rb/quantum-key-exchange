package qubit

import "testing"

func TestQubitClass(t *testing.T) {
	var expected_testVal1 int8 = 0;
	var expected_testVal2 int8 = 1;
	
	for i := 0; i < 100; i++ {
		var qubit *Qubit = NewQubit()
		var measured int8 = qubit.Measure(0)
		if measured != expected_testVal1 && expected_testVal2 != 1 { // measured has to be 1 or 0
			t.Fatalf("Qubit.Measure(polarization) -> expected: %v or %v, got: %v", expected_testVal1, expected_testVal2, measured)
		}
	}
}