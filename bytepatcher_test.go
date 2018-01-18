package main

import "testing"

func TestGoodSigCheck(t *testing.T) {

	// sigCheck(input []byte, sig []byte)
	good := sigCheck([]byte{31, 32, 33}, []byte{31, 32, 33})
	if !good {
		t.Errorf("good sigCheck failed, got: %v, want: %v", good, true)
	}
}

func TestBadSigCheck(t *testing.T) {

	// sigCheck(input []byte, sig []byte)
	bad := sigCheck([]byte{31, 32, 33}, []byte{33, 33, 33})
	if bad {
		t.Errorf("bad sigCheck failed, got: %v, want: %v", bad, false)
	}
}
