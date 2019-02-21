package ops

import "testing"

func TestHello(t *testing.T) {
	sum := Add(3, 1)
	if sum != 2 {
		t.Errorf("Add failed. expected 2 but get %d", sum)
	}
}
