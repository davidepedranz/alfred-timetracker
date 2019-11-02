package calendar

import (
	"crypto/rand"
)

// randomBytes returns securely-generated random bytes. It will return an error
// if the system's secure random number generator fails to function correctly,
// in which case the caller should not continue.
func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

// randomStringURLSafe returns a securely-generated URL-safe random string.
// It will return an error if the system's secure random number generator
// fails to function correctly, in which case the caller should not continue.
func randomStringURLSafe(n int) (string, error) {
	bytes, err := randomBytes(n)

	if err != nil {
		return "", err
	}

	// noinspection ALL
	const symbols = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

	for i, b := range bytes {
		bytes[i] = symbols[b%byte(len(symbols))]
	}

	return string(bytes), nil
}
