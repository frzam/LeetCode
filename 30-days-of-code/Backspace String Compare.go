func backspaceCompare(S string, T string) bool {
	var n string
	for i := 0; i < len(S); i++ {
		if S[i] != '#' {
			n += string(S[i])
		} else if S[i] == '#' && len(n) > 0 {
			n = n[:len(n)-1]
		}
	}
	S = n
	n = ""
	for i := 0; i < len(T); i++ {
		if T[i] != '#' {
			n += string(T[i])
		} else if T[i] == '#' && len(n) > 0 {
			n = n[:len(n)-1]
		}
	}
	T = n

	return S == T
}