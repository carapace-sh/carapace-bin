package golang

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
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
				return styles.Golang.Variable
			case "func":
				return styles.Golang.Function
			case "type":
				return styles.Golang.Type
			case "const":
				return styles.Golang.Constant
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

type MethodOrFieldOpts struct {
	Package string
	Symbol  string
}

// ActionMethodOrFields completes methods and fields of given symbol
//
//	Cache
//	Chdir
func ActionMethodOrFields(opts MethodOrFieldOpts) carapace.Action {
	return carapace.ActionExecCommand("go", "doc", "-short", opts.Package, opts.Symbol)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(fmt.Sprintf(`^func \([^ ]+ \*?%v\) (?P<method>[^( =]+).*`, opts.Symbol))

		methods := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				methods = append(methods, matches[1])
			}
		}

		found := false
		fields := make([]string, 0)
		for _, line := range lines {
			if strings.HasPrefix(line, fmt.Sprintf("type %v ", opts.Symbol)) {
				found = true
				continue
			}

			if !found {
				continue
			}

			if strings.HasPrefix(line, "}") {
				break
			}

			line = strings.SplitN(line, "//", 2)[0]
			line = strings.TrimSpace(line)
			if f := strings.Fields(line); len(f) > 1 {
				fields = append(fields, f[0])
			}
		}

		return carapace.Batch(
			carapace.ActionValues(methods...).Style(styles.Golang.Function),
			carapace.ActionValues(fields...).Style(styles.Golang.Field),
		).ToA()
	})
}
