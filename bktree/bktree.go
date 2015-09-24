package bktree

import (
	"fmt"
	_ "os"
)

// Function call to create a new Bktree
// It returns a pointer to a new Bktree structure
func NewTree() *Bktree {
	return &Bktree{
		root: nil,
	}
}

// Function call to create a new bktreeNode
// It returns a pointer to a new bktreeNode struct
func newbktreeNode(s string) *bktreeNode {
	return &bktreeNode{
		str:   s,
		child: make(map[int]*bktreeNode),
	}
}

// Given a string and a node, insert the string
func (b *Bktree) insert(root *bktreeNode, s string) {

	d := levenshtein(root.str, s)

	if d == 0 {
		return
	}

	for distance, node := range root.child {
		if d == distance {
			b.insert(node, s)
			return
		}
	}

	// if distance is new
	ch := newbktreeNode(s)
	root.child[d] = ch
}

// Insert a node in a Bktree
// It inserts the string at the right place in Bktree struct
func (b *Bktree) Insert(s string) {
	if b.root == nil {
		b.root = newbktreeNode(s)
		return
	} else {
		b.insert(b.root, s)
		return
	}
}

// Internal function to traverse the tree
func traverse(node *bktreeNode) {
	if node == nil {
		fmt.Printf("Empty \n")
		return
	}

	fmt.Printf("Parent: %s\n", node.str)
	for distance, node := range node.child {
		fmt.Printf("%d %s\n", distance, node.str)
		traverse(node)
	}
	fmt.Printf("Returning %s\n", node.str)
}

// Internal function to traverse the tree, to use it
// change name of function to Traverse()
func (b *Bktree) traverse1() {
	traverse(b.root)
}

// Private function to find a string
func find(node *bktreeNode, s string, l int) (res []string) {
	d := levenshtein(node.str, s)
	min_d := d - l
	max_d := d + l

	if d <= l {
		res = append(res, node.str)
	}

	for dis, n := range node.child {
		if dis >= min_d && dis <= max_d {
			res = append(res, find(n, s, l)...)
		}
	}
	return
}

// Function to find all strings within the tolerance levenshtein distance
func (b *Bktree) Find(s string, l int) []string {
	return find(b.root, s, l)
}
