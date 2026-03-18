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

// --- TEST EXACT MATCH ---
func TestCast_ExactMatch(t *testing.T) {
	out := captureOutput(func() {
		cast("accio")
	})

	if out == "" {
		t.Errorf("expected description for exact match, got empty output")
	}
}

// --- TEST SUGGESTION ---
func TestCast_Suggestion(t *testing.T) {
	out := captureOutput(func() {
		cast("lumosx")
	})

	if out == "" {
		t.Fatalf("expected suggestions, got empty output")
	}

	if !contains(out, "lumos") {
		t.Errorf("expected 'lumos' in suggestions, got: %s", out)
	}
}

// --- helper contains ---
func contains(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}
