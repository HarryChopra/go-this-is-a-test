package sum

import "testing"

func TestSum(t *testing.T) {
	result := Sum(1, 2, 3, -5)
	if result != 2 {
		t.Errorf("Wanted 2 but received %d", result)
	}
}
