package bktree

// Define a type to implement a bktree node
// The bktreeNode holds two values:
// str of type string to hold a word from the dictionary
// child is a map of levenshtein distance to a node
type bktreeNode struct {
	str   string
	child map[int]*bktreeNode
}

// Define a type to implement a Bktree
// The Bktree structure holds a root which is a pointer to
// the bktreeNode type
type Bktree struct {
	root *bktreeNode
}
