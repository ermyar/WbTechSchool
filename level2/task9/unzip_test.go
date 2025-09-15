package task9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnzip(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected string
		err      error
	}{
		{input: "abacaba", expected: "abacaba", err: nil},
		{input: "a4bc2d5e", expected: "aaaabccddddde", err: nil},
		{input: "45", expected: "", err: ErrNoChar},
		{input: "", expected: "", err: nil},
		{input: "世世4", expected: "世世世世世", err: nil},
		{input: "\\4\\5", expected: "45", err: nil},
		{input: "qwe\\45", expected: "qwe44444", err: nil},
		{input: "\\n", expected: "", err: ErrUnexpectedChar},
		{input: "\\", expected: "", err: ErrUnexpectedChar},
	} {
		out, err := unzip(tc.input)
		assert.ErrorIs(t, err, tc.err)
		if tc.err == nil {
			assert.Equal(t, out, tc.expected)
		}
	}
}
