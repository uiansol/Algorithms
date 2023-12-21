// Accepted  33 ms  6 MB

func firstUniqChar(s string) int {
	type Location struct {
		qt int
		i  int
	}

	letters := make(map[rune]Location)

	for i, v := range s {
		if _, ok := letters[v]; ok {
			letters[v] = Location{qt: letters[v].qt + 1}
		} else {
			letters[v] = Location{qt: 1, i: i}
		}

	}

	fmt.Println(letters)

	idx := len(s)
	for _, v := range letters {
		if v.qt == 1 && v.i < idx {
			idx = v.i
		}
	}

	if idx == len(s) {
		idx = -1
	}

	return idx
}