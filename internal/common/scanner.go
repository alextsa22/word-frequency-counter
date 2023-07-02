package common

// ScanWords is a split function for a bufio.Scanner.
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	for ; start < len(data); start++ {
		if isLatinLetter(data[start]) {
			break
		}
	}
	for i := start; i < len(data); i++ {
		if !isLatinLetter(data[i]) {
			return i, data[start:i], nil
		}
	}
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	return start, nil, nil
}

func isLatinLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}
