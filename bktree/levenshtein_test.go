package bktree

import (
	"testing"
)

func TestLevenshtein(t *testing.T) {
	a, b := "cook", "book"

	if levenshtein(a, b) != 1 {
		t.Fatalf("Error %s %s != %d", a, b, 1)
	}

	a, b = "Hello", "World"

	if levenshtein(a, b) != 4 {
		t.Fatalf("Error %s %s != %d", a, b, 4)
	}

	a, b = "", "Hello"

	if levenshtein(a, b) != 5 {
		t.Fatalf("Error %s %s != %d", a, b, 5)
	}

	a, b = "Hello", ""

	if levenshtein(a, b) != 5 {
		t.Fatalf("Error %s %s != %d", a, b, 5)
	}

	if levenshtein("a", "b") != 1 {
		t.Fatalf("Error %s %s != %d", "a", "b", 1)
	}

	if levenshtein("aa", "bb") != 2 {
		t.Fatalf("Error %s %s != %d", "aa", "bb", 2)
	}

	if levenshtein("abcd", "dcba") != 4 {
		t.Fatalf("Error %s %s != %d", "abcd", "dcba", 4)
	}
}
