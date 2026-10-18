package winget

import (
	"fmt"
	"strings"
	"testing"
)

func TestParsePackages(t *testing.T) {
	tests := map[string]struct {
		output   string
		expected string
	}{
		"upgradable": {
			output: dropSummary(strings.Join([]string{
				fmt.Sprintf("%-36s %-19s %-10s %-12s %s", "Name", "ID", "Version", "Available", "Source"),
				strings.Repeat("-", 90),
				fmt.Sprintf("%-36s %-19s %-10s %-12s %s", "7-Zip 26.02 (x64 edition)", "7zip.7zip", "26.02.00.0", "26.03", "winget"),
				fmt.Sprintf("%-36s %-19s %-10s %-12s %s", "Nushell", "Nushell.Nushell", "0.114.1", "0.116.0", "winget"),
				"2 upgrades available.",
			}, "\n")),
			expected: "7zip.7zip|7-Zip 26.02 (x64 edition)|Nushell.Nushell|Nushell",
		},
		"localized": {
			output: dropSummary(strings.Join([]string{
				fmt.Sprintf("%-4s %-8s %s", "Nome", "ID", "Versão"),
				strings.Repeat("-", 20),
				fmt.Sprintf("%-4s %-8s %s", "abc", "a.b", "1.0"),
				"1 atualização disponível.",
			}, "\n")),
			expected: "a.b|abc",
		},
		"search": {
			output: strings.Join([]string{
				fmt.Sprintf("%-10s %-15s %-8s %-12s %s", "Name", "Id", "Version", "Match", "Source"),
				strings.Repeat("-", 70),
				fmt.Sprintf("%-10s %-15s %-8s %-12s %s", "7-Zip", "7zip.7zip", "24.09", "", "winget"),
				fmt.Sprintf("%-10s %-15s %-8s %-12s %s", "Nushell", "Nushell.Nushell", "0.99.1", "Tag: shell", "winget"),
			}, "\n"),
			expected: "7zip.7zip|7-Zip|Nushell.Nushell|Nushell",
		},
		"carriage-return": {
			output:   "Name  ID\n------\nabc   a.b\n1 upgrades available.",
			expected: "a.b|abc",
		},
		"empty": {
			output:   "",
			expected: "",
		},
		"header-only": {
			output:   "Name  ID",
			expected: "",
		},
		"no-id": {
			output:   "Name  ID\n-----\nabc\n1 upgrades available.",
			expected: "",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if actual := strings.Join(parsePackages(test.output), "|"); actual != test.expected {
				t.Errorf("expected: '%v' actual: '%v'", test.expected, actual)
			}
		})
	}
}
