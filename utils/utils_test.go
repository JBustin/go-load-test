package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Dedupe(t *testing.T) {
	mySlice := []string{"a", "b", "a", "c", "d", "e", "e", "f"}

	assert.Equal(
		t,
		[]string{"a", "b", "c", "d", "e", "f"},
		Dedupe(mySlice),
		"Dedupe slice with any duplicate entries",
	)

	mySlice = []string{"a", "b", "c", "d", "e", "f"}

	assert.Equal(
		t,
		[]string{"a", "b", "c", "d", "e", "f"},
		Dedupe(mySlice),
		"Dedupe slice without any duplicate entries",
	)
}

func Test_Fill(t *testing.T) {
	mySlice := []string{"a", "b"}

	assert.Equal(
		t,
		[]string{"a", "b", "a", "b", "a", "b"},
		Fill(mySlice, 6),
		"Fill slice by repeating existing values",
	)

	mySlice = []string{"a", "b", "c"}

	assert.Equal(
		t,
		[]string{"a", "b"},
		Fill(mySlice, 2),
		"Fill slice by keeping only expected size",
	)
}
