// Accepted  7 ms  6.54 MB

func romanToInt(s string) int {
	values := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900}
	total, i, l := 0, 0, len(s)

	for i < l {
		if i < l-1 {
			if v, ok := values[s[i:i+2]]; ok {
				total += v
				i += 2
			} else {
				total += values[string(s[i])]
				i += 1
			}
		} else {
			total += values[string(s[i])]
			i += 1
		}
	}

	return total
}