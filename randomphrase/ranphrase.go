package randomphrase

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// usage is a map of string to struct{} to keep track of used words.
// An empty struct{} uses no memnory.
type usage map[string]struct{}

// buildUnique generates a password with n unique words.
func GenerateUnique(n int) (string, error) {

	if n > len(dict) {
		return "", fmt.Errorf("cannot generate password with unique words since length '%v' is greater than dictionary size", n)
	}

	used := make(usage)
	res := []string{}

	l := len(dict)

	for range n {

		word := dict[rand.Intn(l)]

		for alreadyUsed(used, word) {
			word = dict[rand.Intn(l)]
		}
		res = append(res, fmt.Sprintf("%v", word))
	}
	return strings.Join(res, "-"), nil
}

// CodeSpace returns the size of the random code space.
func CodeSpace(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("cannot generate password with negative length")
	}
	m := len(dict)
	return int(math.Pow(float64(m), float64(n))), nil
}

// GenerateWithFormat generates a password with n words and a separator.
// If upper is true, the first letter of each word is capitalized.
func GenerateWithFormat(n int, sep string, upper bool) (string, error) {
	return generateWithFormat(dict, n, sep, upper)
}

func generateWithFormat(d []string, n int, sep string, upper bool) (string, error) {
	result := build(d, n, upper)
	return strings.Join(result, sep), nil
}

// build generates a password with n words.
func Generate(n int) (string, error) {

	if n == 0 {
		return "", fmt.Errorf("cannot generate zero length password")
	}

	res := build(dict, n, false)

	return strings.Join(res, "-"), nil
}

func build(words []string, n int, upper bool) []string {
	result := make([]string, n)
	l := len(words)

	for i := range n {
		word := words[rand.Intn(l)]
		if upper {
			word = strings.ToUpper(string(word[0])) + word[1:]
		}
		result[i] = word
	}
	return result
}

func alreadyUsed(m usage, word string) bool {
	_, ok := m[word]
	if !ok {
		m[word] = struct{}{}
	}
	return ok
}
