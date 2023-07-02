package common

func BytesToLower(b []byte) []byte {
	var ret []byte
	for _, c := range b {
		ret = append(ret, letterToLower(c))
	}
	return ret
}

func letterToLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
