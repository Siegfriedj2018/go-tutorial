package test

import (
	pokedexcli "github.com/siegfriedj2018/bootdotdev/pokedexcli/pokedex"
	"testing"
)

func TestCleanInput(t *testing.T) {
	// This is the main testing function for pokedex file
	cases := []struct {
		input string
		expected []string
	}{
		{
			input:	"  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:	"Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander","bulbasaur","pikachu"},
		},
		{
			input:	"  hello   ",
			expected: []string{"hello"},
		},
		{
			input:	"",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := pokedexcli.CleanInput(c.input)
		//Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error mesage
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("clean input length != expected length, %v != %v", actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice 
			// if they don't match, use t.Errorf to print an error mesage
			// and fail the test
			if word != expectedWord {
				t.Errorf("input: %s\nexpected: %s", word, expectedWord)
			}
		}
	}
}