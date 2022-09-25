package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestEchoHello(t *testing.T) {
	got := extractStdout(t, EchoTest)
	want := "Test"
	if got != want {
		t.Errorf("failed to test. got: %s, want: %s", got, want)
	}

}

func extractStdout(t *testing.T, fnc func()) string {
	t.Helper()

	orgStdout := os.Stdout
	defer func() {
		os.Stdout = orgStdout
	}()
	r, w, _ := os.Pipe()
	os.Stdout = w
	fnc()
	w.Close()
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("failed to read buf: %v", err)
	}
	return strings.TrimRight(buf.String(), "\n")
}
