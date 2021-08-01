// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package word

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"unicode"
)

//!+bench

//!-bench

//!+test
func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

//!-test

//!+bench
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

//!-bench

//!+example

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}

//!-example

/*
//!+random
import "math/rand"

//!-random
*/

//!+random
// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(addPunctuation(runes, rng))
}

var puncs = []rune{',', '.', ':', ' '}

func addPunctuation(runes []rune, rng *rand.Rand) []rune {
	n := rng.Intn(10)
	for i := 0; i < n; i++ {
		punc := puncs[rng.Intn(len(puncs))]
		pos := rng.Intn(len(runes) + 1) // If pos == len, append after last elem
		if len(runes) == pos {
			runes = append(runes, punc)
		} else {
			runes = append(runes[:pos+1], runes[pos:]...)
			runes[pos] = punc
		}
	}
	return runes
}

func nonPalindromize(p string, rng *rand.Rand) (np string, ok bool) {
	var letterPos []int
	var runes []rune
	for _, r := range p {
		if unicode.IsLetter(r) {
			letterPos = append(letterPos, len(runes))
		}
		runes = append(runes, r)
	}
	if len(letterPos) < 2 {
		return "", false
	}
	target := letterPos[0]
	for {
		r := rune(rng.Intn(0x1000))
		if unicode.IsLetter(r) && !isEqualLetter(r, runes[target]) {
			runes[target] = r
			break
		}
	}
	return string(runes), true
}

func isEqualLetter(a, b rune) bool {
	return unicode.ToLower(a) == unicode.ToLower(b)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		np, ok := nonPalindromize(p, rng)
		if !ok {
			continue
		}
		if IsPalindrome(np) {
			t.Errorf("IsPalindrome(%q) = true", np)
		}
	}
}

//!-random

/*
// Answer for Exercicse 11.1: Modify randomPalindrome to exercise
// IsPalindrome's handling of punctuation and spaces.

// WARNING: the conversion r -> upper -> lower doesn't preserve
// the value of r in some cases, e.g., µ Μ, ſ S, ı I

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomNoisyPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x200)) // random rune up to \u99
		runes[i] = r
		r1 := r
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			r = unicode.ToUpper(r)
			if unicode.ToLower(r) != r1 {
				fmt.Printf("cap? %c %c\n", r1, r)
			}
		}
		runes[n-1-i] = r
	}
	return "?" + string(runes) + "!"
}

func TestRandomNoisyPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	n := 0
	for i := 0; i < 1000; i++ {
		p := randomNoisyPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsNoisyPalindrome(%q) = false", p)
			n++
		}
	}
	fmt.Fprintf(os.Stderr, "fail = %d\n", n)
}
*/
