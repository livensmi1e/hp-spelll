package main

import (
	"bytes"
	"os"
	"testing"
)

// --- helper: capture stdout ---
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

// --- helper contains ---
func contains(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}

// --- TEST LEVENSHTEIN ---
func TestLevenshtein(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		{"", "", 0},
		{"a", "", 1},
		{"", "abc", 3},
		{"commit", "comit", 1},
		{"comm", "commit", 2},
		{"kitten", "sitting", 3},
	}

	for _, tt := range tests {
		got := levenshtein(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("levenshtein(%q, %q) = %d, want %d",
				tt.a, tt.b, got, tt.want)
		}
	}
}

// --- TEST THRESHOLD ---
func TestThreshold(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"a", 2},
		{"ab", 2},
		{"abc", 2},
		{"abcdef", 2},    // len/3 = 2
		{"abcdefghi", 3}, // len/3 = 3
	}

	for _, tt := range tests {
		got := threshold(tt.input)
		if got != tt.want {
			t.Errorf("threshold(%q) = %d, want %d",
				tt.input, got, tt.want)
		}
	}
}

// --- TEST DOCAST: EXACT MATCH ---
func TestDoCast_ExactMatch(t *testing.T) {
	res := doCast("accio")

	want := "summon object toward caster"
	if res != want {
		t.Errorf("expected %q, got %q", want, res)
	}
}

// --- TEST DOCAST: SUGGESTION ---
func TestDoCast_Suggestion(t *testing.T) {
	res := doCast("lumosx")

	if res == "" {
		t.Fatalf("expected suggestions, got empty string")
	}

	if !contains(res, "lumos") {
		t.Errorf("expected 'lumos' in suggestions, got: %s", res)
	}
}

// --- TEST DOCAST: MULTIPLE SUGGESTIONS FORMAT ---
func TestDoCast_MultipleSuggestions(t *testing.T) {
	res := doCast("lum")

	if !contains(res, "The most similar spell") {
		t.Errorf("unexpected format: %s", res)
	}
}

// --- TEST DOCAST: NO MATCH ---
func TestDoCast_NoMatch(t *testing.T) {
	res := doCast("zzzzzz")

	want := "Avada Kedavra"
	if res != want {
		t.Errorf("expected %q, got %q", want, res)
	}
}

// --- TEST CAST (I/O layer) ---
func TestCast_Print(t *testing.T) {
	out := captureOutput(func() {
		cast("hello")
	})

	want := "hello\n"
	if out != want {
		t.Errorf("expected %q, got %q", want, out)
	}
}
