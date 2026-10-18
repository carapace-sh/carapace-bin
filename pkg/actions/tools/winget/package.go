package winget

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

var columnRegexp = regexp.MustCompile(`\S+`)

// ActionUpgradablePackages completes packages with an available upgrade
//
//	7zip.7zip (7-Zip 26.02 (x64 edition))
//	Nushell.Nushell (Nushell)
func ActionUpgradablePackages() carapace.Action {
	return carapace.ActionExecCommand("winget", "list", "--accept-source-agreements", "--include-unknown", "--upgrade-available")(func(output []byte) carapace.Action {
		return carapace.ActionValuesDescribed(parsePackages(dropSummary(string(output)))...)
	}).Tag("upgradable packages")
}

// ActionSearchPackages completes packages available from the configured sources
//
//	7zip.7zip (7-Zip)
//	Nushell.Nushell (Nushell)
func ActionSearchPackages() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"search", "--accept-source-agreements"}
		if c.Value != "" {
			args = append(args, c.Value)
		}
		return carapace.ActionExecCommand("winget", args...)(func(output []byte) carapace.Action {
			return carapace.ActionValuesDescribed(parsePackages(string(output))...)
		})
	}).Tag("packages")
}

func parsePackages(output string) []string {
	lines := strings.Split(strings.ReplaceAll(output, "\r\n", "\n"), "\n")

	header := -1
	for i, line := range lines {
		if strings.TrimSpace(line) != "" {
			header = i
			break
		}
	}
	if header == -1 {
		return []string{}
	}

	starts := make([]int, 0)
	for _, match := range columnRegexp.FindAllStringIndex(lines[header], -1) {
		starts = append(starts, match[0])
	}
	if len(starts) < 2 {
		return []string{}
	}

	vals := make([]string, 0)
	for _, line := range lines[header+1:] {
		if strings.TrimSpace(line) == "" || strings.Trim(strings.TrimSpace(line), "- ") == "" {
			continue
		}

		name := column(line, starts, 0)
		id := column(line, starts, 1)
		if id == "" || strings.Contains(id, " ") {
			continue
		}
		vals = append(vals, id, name)
	}
	return vals
}

func dropSummary(output string) string {
	lines := strings.Split(strings.ReplaceAll(output, "\r\n", "\n"), "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		if strings.TrimSpace(lines[i]) != "" {
			return strings.Join(lines[:i], "\n")
		}
	}
	return output
}

func column(line string, starts []int, index int) string {
	if index >= len(starts) || starts[index] > len(line) {
		return ""
	}

	end := len(line)
	if index+1 < len(starts) && starts[index+1] < end {
		end = starts[index+1]
	}
	return strings.TrimSpace(line[starts[index]:end])
}
