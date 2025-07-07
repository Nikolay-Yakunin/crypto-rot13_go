package crypto

import (
	"testing"
)

func TestCryptRot13(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"abc", "nop"},
		{"ABC", "NOP"},
		{"Hello, World!", "Uryyb, Jbeyq!"},
		{"12345", "12345"}}

	for _, c := range cases {
		result := CryptRot13(c.input)
		if result != c.expected {
			t.Errorf("CryptRot13(%q) = %q; exp %q", c.input, result, c.expected)
		}
	}
}
