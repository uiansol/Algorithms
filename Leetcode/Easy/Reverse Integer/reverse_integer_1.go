// Accepted  0 ms  2.1 MB

func reverse(x int) int {
	numString := strconv.Itoa(x)
	reversed := reverseString(numString)
	num, _ := strconv.Atoi(reversed)

	if num < int(math.Pow(-2, 31)) || num > int(math.Pow(2, 31))-1 {
		return 0
	}

	return num
}

func reverseString(s string) string {
	rns := []rune(s)

	if rns[0] == '-' {
		for i, j := 1, len(rns)-1; i <= j; i, j = i+1, j-1 {
			rns[i], rns[j] = rns[j], rns[i]
		}
	} else {
		for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
			rns[i], rns[j] = rns[j], rns[i]
		}
	}

	return string(rns)
} 