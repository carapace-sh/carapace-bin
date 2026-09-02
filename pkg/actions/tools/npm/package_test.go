package npm

import (
	"strings"
	"testing"
)

func TestParseSearchOutput(t *testing.T) {
	tests := map[string]struct {
		output   string
		expected string
	}{
		"name and description": {
			output:   "express\ta minimal web framework\n",
			expected: "express|a minimal web framework",
		},
		"multiple packages": {
			output:   "express\tframework\nlodash\tutilities\n",
			expected: "express|framework|lodash|utilities",
		},
		// A package with no description yields a line with only the name (npm
		// omits trailing empty fields); this previously panicked on fields[1].
		"missing description": {
			output:   "ty\n",
			expected: "ty|",
		},
		"mixed with and without description": {
			output:   "ty\nexpress\tframework\n",
			expected: "ty||express|framework",
		},
		"empty output": {
			output:   "",
			expected: "",
		},
		"only trailing newline": {
			output:   "\n",
			expected: "",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if actual := strings.Join(parseSearchOutput([]byte(tc.output)), "|"); actual != tc.expected {
				t.Errorf("expected: '%v' actual: '%v'", tc.expected, actual)
			}
		})
	}
}
