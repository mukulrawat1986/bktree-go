package bktree

import (
	_ "fmt"
	"testing"
)

func TestInset(t *testing.T) {
	bk := NewTree()

	bk.Insert("ABC")
	if bk.root.str != "ABC" {
		t.Fatalf("Error")
	}

	bk.Insert("DEF")
	if bk.root.child[3].str != "DEF" {
		t.Fatalf("Error")
	}

	bk.Insert("ABA")
	if bk.root.child[1].str != "ABA" {
		t.Fatalf("Error")
	}

	bk.Insert("ABB")
	if bk.root.child[1].child[1].str != "ABB" {
		t.Fatalf("Error")
	}
}

func TestFind(t *testing.T) {
	bk := NewTree()
	bk.Insert("Book")
	bk.Insert("Books")
	bk.Insert("Cake")
	bk.Insert("Boo")
	bk.Insert("Cape")
	bk.Insert("Cart")
	bk.Insert("Boon")
	bk.Insert("Cook")
	res := bk.Find("Caqe", 1)
	if res[0] != "Cake" && res[1] != "Cape" {
		t.Fatalf("Error")
	}
}
