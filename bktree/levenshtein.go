package bktree

// The function find Levenshtein distance between two strings
// It takes as argument two strings and returns an integer value
// which is the levenshtein distance between the two strings
func Levenshtein(a, b string) int{

    // Find the length of the strings
    la, lb := len(a), len(b)

    // Create array to store distance
    col := make([]int, la+1)

    var lastdiag, olddiag, i, j int
    var temp1, temp2, temp3 int

    for i = 1; i <= la; i++{
        col[i] = i
    }

    for i = 1; i<= lb; i++{
        col[0] = i

        for j, lastdiag = 1, i-1; j<=la; j++{
            olddiag = col[j]
            temp1 = col[j] + 1
            temp2 = col[j-1] + 1
            if a[j-1] == b[i-1]{
                temp3 = lastdiag + 0
            }else{
                temp3 = lastdiag + 1
            }

            col[j] = min(temp1, temp2, temp3)
            lastdiag = olddiag
        }
    }

    return col[la]
}

// A function to find the minimum of two integers
// Returns the minimum of the two integers
func min(a, b, c int) int {
    if a < b && a < c{
        return a
    }else if b < a && b <c{
        return b
    }else{
        return c
    }
}
