package golang

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionSymbols completes symbols of given package
//
//	Action
//	ActionCallback
func ActionSymbols(pkg string) carapace.Action {
	return carapace.ActionExecCommand("go", "doc", "-short", pkg)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^ *(?P<type>var|func|type|const) (?P<symbol>[^( =]+).*`) // TODO incomplete (e.g. generics))

		styleFor := func(s string) string {
			switch s {
			case "var":
				return style.Yellow
			case "func":
				return style.Blue
			case "type":
				return style.Magenta
			case "const":
				return style.Green
			default:
				return style.Default
			}
		}

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[2], styleFor(matches[1]))
			}
		}
		return carapace.ActionStyledValues(vals...)
	})
}

// ActionMethodOrFields completes methods and (TODO fields) of given symbol
//
//	Cache
//	Chdir
func ActionMethodOrFields(pkg, symbol string) carapace.Action {
	return carapace.ActionExecCommand("go", "doc", "-short", pkg, symbol)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(fmt.Sprintf(`^func \([^ ]+ \*?%v\) (?P<method>[^( =]+).*`, symbol)) // TODO only methods for now (fields missing)

		methods := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				methods = append(methods, matches[1])
			}
		}

		found := false
		fields := make([]string, 0)
		for _, line := range lines {
			if strings.HasPrefix(line, fmt.Sprintf("type %v ", symbol)) {
				found = true
				continue
			}

			if !found {
				continue
			}

			if strings.HasPrefix(line, "}") {
				break
			}

			if line := strings.TrimSpace(line); line != "" && !strings.HasPrefix(line, "//") {
				fields = append(fields, strings.Fields(line)[0])
			}
		}

		return carapace.Batch(
			carapace.ActionValues(methods...).Style(style.Blue),
			carapace.ActionValues(fields...).Style(style.Yellow),
		).ToA()
	})
}
