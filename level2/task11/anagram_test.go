package task11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	for _, tc := range []struct {
		input  []string
		output map[string][]string
	}{
		{
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"},
			output: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
	} {

		mp := GetAnagram(tc.input...)

		for k, v := range tc.output {
			if arr, ok := mp[k]; ok {
				for i := range arr {
					require.Equal(t, v[i], arr[i])
				}
			}
		}
	}
}
