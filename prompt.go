
func fixQuote(word string) string {
	// Convert to rune slice to handle potential unicode, 
	// though not strictly necessary for simple '
	runes := []rune(word)
	var result []rune
	n := len(runes)

	for i := 0; i < n; i++ {
		// Check for pattern: Quote + Space -> ' 
		if runes[i] == '\'' && i+1 < n && runes[i+1] == ' ' {
			result = append(result, runes[i])
			i++ // Skip the space
			continue
		}
		// Check for pattern: Space + Quote ->  '
		if runes[i] == ' ' && i+1 < n && runes[i+1] == '\'' {
			// Skip adding the space
			continue
		}
		result = append(result, runes[i])
	}
	return string(result)
}
