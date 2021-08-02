package demo

import "testing"

func TestFuncA1(t *testing.T) {
	t.Logf("funcA error description")
	t.Fail()
	t.Log("continuing other tests..")
}
func TestFuncA2(t *testing.T) {
	t.Errorf("funcA error description") // Errorf = Fail + Logf
	t.Log("continuing other tests..")
}

func TestFuncB1(t *testing.T) {
	t.Logf("funcB error description")
	t.FailNow()
	t.Log("continuing other tests..") // This won't log as the test would stop immediately
}

func TestFuncB2(t *testing.T) {
	t.Fatalf("funcB error description") // Fatalf = Logf + FailNow()
	t.Log("continuing other tests..")   // This won't log
}
