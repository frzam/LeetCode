func isHappy(n int) bool {
	x := n
	s := n
	for n > 9 {
		s = 0
		for n > 0 {
			r := n % 10
			s += r * r
			n = n / 10
		}
		if s == 1 {
			return true
		} else if s == x {
			return false
		}
		n = s
	}
	if s == 7 || s == 1 {
		return true
	}
	return false
}