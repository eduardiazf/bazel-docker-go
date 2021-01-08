package shared

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 2)

	if result == 4 {
		t.Logf("Sum: 2 + 2 = %d", result)
	}
}
