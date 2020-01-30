package euclid

// gcd returns the greatest common divisor
func gcd(p, q uint) uint {
	if q == 0 {
		return p
	}
	r := p % q
	return gcd(q, r)
}
