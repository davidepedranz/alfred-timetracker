package timetracker

import (
	"crypto/rand"
	"github.com/dvsekhvalnov/jose2go/base64url"
)

// TODO: this creates a string longer than the number of bytes
func randomStringURLSafe(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic("Cannot generate a random string")
	}
	return base64url.Encode(b)
}
