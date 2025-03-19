package fruuid

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

var alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-" // len 64

// GenerateWithAlpha Return a random string of length 'n' based on the supplied string of letters.
// Length of string should be >0 or an error will be returned.
// Supplied letters should be of length > 0 .
// The number of unique strigns is len(alpha) ^ n.
// For comparison a random UUID (type 4) has a random space of 5.3 x 10 ^36, with a supplied random string of length 64
// you would require a string of length 21 to achieve a random space of 8.507059173×10³⁷
func GenerateWithAlpha(n int, al string) (string, error) {
	if n < 1 {
		return "", fmt.Errorf("length of random string must be >= 1")
	}
	if len(al) < 1 {
		return "", fmt.Errorf("length of dictionary must >= 1")
	}
	return generate(n, al)
}

// Generate Return a random string of length 'n' based on the built in alphabet.
// The number of unique strigns is 64 ^ n.
// For comparison a random UUID (type 4) has a random space of 5.3 x 10 ^36, using the built in alphabet
// you would only require a string of length 21 to achieve a larger random space = 8.507059173×10³⁷ unique strings.
func Generate(n int) (string, error) {
	if n < 1 {
		return "", fmt.Errorf("length of random string must be >= 1")
	}
	return generate(n, alpha)
}

// GenerateUUID Return a random string yielding better than UUID size of random space = 8.507059173×10³⁷ unique strings.
func GenerateUUID() (string, error) {
	return generate(21, alpha)
}

func generate(n int, al string) (string, error) {
	if n < 1 {
		return "", fmt.Errorf("length of random string must be >= 1")
	}
	bytes := make([]byte, n)

	la := int64(len(al))
	for c := range n {
		bytes[c] = alpha[cryptoRandSecure(la)]
	}
	return string(bytes), nil

}

// Generate a cryptographically secure random integer across the half open interval [0, max]
func cryptoRandSecure(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Println(err)
	}
	return nBig.Int64()
}
