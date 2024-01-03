package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	total := Add(2, 3)
	if total != 4 {
		t.Errorf("Sum was incorrect, Actual: %d, Expected %d", total, 4)
	}
	t.Log("running TestAdd")
}
