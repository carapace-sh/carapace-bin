package but

// TODO
//
//lint:ignore U1000 wip
func hash(input string) string {
	var hash uint64 = 0
	for i := 0; i < len(input); i++ {
		hash = hash*31 + uint64(input[i])
	}

	firstChars := "ghijklmnopqrstuvwxyz"
	firstChar := string(firstChars[hash%20])
	hash /= 20

	secondChars := "0123456789abcdefghijklmnopqrstuvwxyz"
	secondChar := string(secondChars[hash%36])

	return firstChar + secondChar
}
