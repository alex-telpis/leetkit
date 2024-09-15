package testutil

import "testing"

func CheckErr(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Errorf("want non-empty error, got nil")
	}
}

func CheckNilErr(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("unexpected error '%v'", err)
	}
}

func CheckVal[A comparable](t *testing.T, want A, got A) {
	t.Helper()

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func CheckSlice[A comparable](t *testing.T, want []A, got []A) {
	t.Helper()

	if len(want) != len(got) {
		t.Errorf("want %v, got %v", want, got)
		return
	}

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("want %v, got %v", want, got)
			return
		}
	}
}

func CheckPanic(t *testing.T, f func()) {
	t.Helper()

	defer func() {
		if r := recover(); r == nil {
			t.Error("want panic")
		}
	}()
	f()
}
