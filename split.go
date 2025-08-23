package piscine

func Split(s, sep string) []string {
	if sep == "" {
		// Separator empty, return s as single element slice
		return []string{s}
	}

	var result []string
	start := 0
	for i := 0; i <= len(s)-len(sep); {
		if s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			i += len(sep)
			start = i
		} else {
			i++
		}
	}
	// Append the last part after the last separator (or the whole string if no sep found)
	result = append(result, s[start:])
	return result
}
