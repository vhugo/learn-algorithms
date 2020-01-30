package euclid

import "testing"

func TestEuclidsAlgorithm(t *testing.T) {
	for _, tc := range []struct {
		name           string
		p, q, expected uint
	}{
		{
			name:     "q is zero",
			p:        255,
			q:        0,
			expected: 255,
		},
		{
			name:     "q is not zero",
			p:        255,
			q:        10,
			expected: 5,
		},
		{
			name:     "exercise 1.1.24",
			p:        1111111,
			q:        1234567,
			expected: 1,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := gcd(tc.p, tc.q)

			if tc.expected != got {
				t.Fatalf("expected %d, got %d", tc.expected, got)
			}
		})
	}
}
