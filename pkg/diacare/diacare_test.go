package diacare

import "testing"

func TestHello(t *testing.T) {
    got := Hello("ci")
    want := "hello ci"
    if got != want {
        t.Fatalf("Hello() = %q, want %q", got, want)
    }
}
