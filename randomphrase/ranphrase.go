package randomphrase

import (
	"fmt"
	"math/rand"
	"strings"
)

// Dictionary of words.
var a = []string{
	"apple", "anaconda", "almost", "above", "alone", "allied", "abandon", "abide", "ability", "able", "abrupt", "absolute", "abstract", "abundant", "academic", "accelerate", "acceptance", "accomplish", "accurate", "azure",
	"bath", "bed", "beyond", "brought", "banana", "brave", "battery", "bakery", "bison", "behold", "bumps", "boiled", "butter", "butterfly",
	"cat", "cucumber", "cook", "cake", "code", "cauldron", "crave", "cave",
	"dog", "day", "date", "drain", "door", "donut", "daily", "degree", "draw",
	"elephant", "equal", "elegant", "einstein", "equinox", "element", "east", "each", "ears",
	"figure", "face", "foot", "flake", "frosty", "forge", "found",
	"giraffe", "green", "grass", "ghost", "grapes",
	"help", "hospital", "home", "hologram", "horse", "hooked",
	"ingot", "input", "infant", "item", "internal", "inward", "invent",
	"jazz", "justice", "juicy", "junior", "jack", "jail", "jelly", "jigsaw", "jolt", "jolly", "jungle", "jupiter",
	"kangaroo", "keen", "knee", "knight", "knot",
	"lost", "limited", "lazy", "lance", "leap", "lunar", "luminous", "lucky", "lucid",
	"mango", "many", "madrigal", "meerkat", "magic", "mystic", "mystery",
	"never", "north", "norway", "nacent", "noble", "nifty", "noodle", "nose",
	"operation", "open", "output", "origin", "orange", "ocean", "octopus", "over",
	"plum", "problem", "police", "pickle", "pine", "palm", "pass", "panic", "points", "power",
	"quench", "quartz", "quack", "quasar", "quite", "quotidian",
	"ranch", "really", "recalcitrant", "reticent", "redolent",
	"sport", "space", "south", "staple", "slant", "string", "solar", "speak", "share", "swim", "stops",
	"truth", "tango", "trench", "talent", "taciturn", "tautology", "tincture", "trains",
	"umbrella", "unusual", "under", "utmost", "uptick", "unreal",
	"violin", "verb", "verge", "valiant", "voter", "visit",
	"weather", "wheels", "wet", "wash", "wonder", "windows", "west", "wheat", "waste", "waters", "wants", "wool", "weeks",
	"xylophone", "xray", "xanadu", "xenophobic",
	"yet", "yarn", "yielding", "yesterday", "young", "yodel", "yonder", "yoga", "yogurt", "yolk", "youth", "yummy", "yeast",
	"zeal", "zebra", "zoo", "zoom", "zinc", "zone", "zero", "zest", "zany", "zombie", "zygote",
}

// usage is a map of string to struct{} to keep track of used words.
// An empty struct{} uses no memnory.
type usage map[string]struct{}

// buildUnique generates a password with n unique words.
func GenerateUnique(n int) (string, error) {

	if n > len(a) {
		return "", fmt.Errorf("cannot generate password with unique words since length '%v' is greater than dictionary size", n)
	}

	used := make(usage)
	res := []string{}

	for range n {

		word := a[rand.Intn(len(a))]

		for alreadyUsed(used, word) {
			word = a[rand.Intn(len(a))]
		}
		res = append(res, fmt.Sprintf("%v", word))
	}
	return strings.Join(res, "-"), nil
}

// GenerateWithFormat generates a password with n words and a separator.
// If upper is true, the first letter of each word is capitalized.
func GenerateWithFormat(n int, sep string, upper bool) (string, error) {
	return generateWithFormat(a, n, sep, upper)
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

	res := build(a, n, false)

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
