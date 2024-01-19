// Accepted  41 ms  4.8 MB

import "strconv"

func isPalindrome(x int) bool {
    xString := strconv.Itoa(x)
    size := len(xString)
    
    for i := 0; i < size / 2; i++ {
        if xString[i] != xString[size - 1 - i] {
            return false
        }
    }
    
    return true
}