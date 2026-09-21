package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func run(t *testing.T, args ...string) string {
	t.Helper()
	var buf bytes.Buffer
	root := NewRootCmd(&buf)
	root.SetArgs(args)
	if err := root.Execute(); err != nil {
		t.Fatal(err)
	}
	return buf.String()
}

func TestGreet(t *testing.T) {
	if got := run(t, "greet", "world"); !strings.Contains(got, "hello, world") {
		t.Fatalf("got %q", got)
	}
}

func TestGreetUpper(t *testing.T) {
	if got := run(t, "greet", "world", "--upper"); !strings.Contains(got, "hello, WORLD") {
		t.Fatalf("got %q", got)
	}
}

func TestMissingArgIsAnError(t *testing.T) {
	root := NewRootCmd(&bytes.Buffer{})
	root.SetArgs([]string{"greet"})
	if err := root.Execute(); err == nil {
		t.Fatal("expected an error for a missing argument")
	}
}
