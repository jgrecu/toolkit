package toolkit

import "crypto/rand"

const (
	randomStringSource = "abcdefghijklmnopqrstuvwxzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"
)

// Tools is the type used to instanciate this module. Any variable of this type will have access
// to all the methods with the receiver *Tools
type Tools struct{}

// RandomString returns a string of random characters of lenth n, using randomStringSource
// as the source for the string
func (t *Tools) RandomString(n int) string {
	if n <= 0 {
		return ""
	}
	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}
