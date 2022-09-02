package pw

import (
	"crypto/rand"
	"math/big"
)

// UpperChars ...
const UpperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// LowerChars ...
const LowerChars = "abcdefghijklmnopqrstuvwxyz"

// SpecialChars ...
const SpecialChars = "!@#$%^&*()-_=+,.?/:;{}[]~"

// NumberChars ...
const NumberChars = "0123456789"

// AllChars ...
const AllChars = UpperChars + LowerChars + SpecialChars + NumberChars

// Option type
type Option func() string

func randomFromChars(length int, chars string) string {
	max := len(chars) - 1
	p := ""

	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
		p = p + string(chars[n.Int64()])
	}

	return p
}

// WithUpperCase option to use upper case characters in the password
func WithUpperCase() Option {
	return func() string {
		return UpperChars
	}
}

// WithLowerCase option to use lower case characters in the password
func WithLowerCase() Option {
	return func() string {
		return LowerChars
	}
}

// WithNumbers option to use number characters in the password
func WithNumbers() Option {
	return func() string {
		return NumberChars
	}
}

// WithSpecial option to use special characters in the password
func WithSpecial() Option {
	return func() string {
		return SpecialChars
	}
}

// NewPassword to create a new password with length. Defaults characters to include
// all characters. Options can be passed to customize the password.
func NewPassword(length int, options ...Option) string {
	if len(options) != 0 {
		c := ""

		for _, option := range options {
			c = c + option()
		}

		return randomFromChars(length, c)
	}

	return randomFromChars(length, AllChars)
}
