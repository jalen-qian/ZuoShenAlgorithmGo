package class_04

import "testing"

func TestCalculate(t *testing.T) {
	t.Log(Add(12, 18))
	t.Logf("256 - 34 = %d", Sub(256, 34))
	t.Logf("-11 * 25 = %d", Mul(-11, 25))
	t.Log(3 / 10)
}
