package golang

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
)

type SymbolOpts struct {
	Package    string
	Unexported bool
}

// ActionSymbols completes symbols of given package
//
//	Action
//	ActionCallback
func ActionSymbols(opts SymbolOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"doc", "-short"}
		if opts.Unexported {
			args = append(args, "-u")
		}
		args = append(args, opts.Package)
		return carapace.ActionExecCommand("go", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^ *(?P<type>var|func|type|const) (?P<symbol>[^( =\[]+).*`) // TODO incomplete (e.g. generics))

			variables := make([]string, 0)
			functions := make([]string, 0)
			types := make([]string, 0)
			constants := make([]string, 0)
			for _, line := range lines {
				if matches := r.FindStringSubmatch(line); matches != nil {
					switch matches[1] {
					case "var":
						variables = append(variables, matches[2])
					case "func":
						functions = append(functions, matches[2])
					case "type":
						types = append(types, matches[2])
					case "const":
						constants = append(constants, matches[2])
					default:
					}
				}
			}
			return carapace.Batch(
				carapace.ActionValues(variables...).Style(styles.Golang.Variable).Tag("variables"),
				carapace.ActionValues(functions...).Style(styles.Golang.Function).Tag("functions"),
				carapace.ActionValues(types...).Style(styles.Golang.Type).Tag("types"),
				carapace.ActionValues(constants...).Style(styles.Golang.Constant).Tag("constants"),
			).ToA()
		})
	})
}

type MethodOrFieldOpts struct {
	Package    string
	Symbol     string
	Unexported bool
}

// ActionMethodOrFields completes methods and fields of given symbol
//
//	Cache
//	Chdir
func ActionMethodOrFields(opts MethodOrFieldOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"doc", "-short"}
		if opts.Unexported {
			args = append(args, "-u")
		}
		args = append(args, opts.Package, opts.Symbol)
		return carapace.ActionExecCommand("go", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			rFunc := regexp.MustCompile(fmt.Sprintf(`^func \([^ ]+ \*?%v(\[[^)]+)?\) (?P<method>[^( =]+).*`, opts.Symbol))
			rType := regexp.MustCompile(fmt.Sprintf(`^type %v[\[ ].*\{$`, opts.Symbol))

			methods := make([]string, 0)
			for _, line := range lines {
				if matches := rFunc.FindStringSubmatch(line); matches != nil {
					methods = append(methods, matches[2])
				}
			}

			found := false
			fields := make([]string, 0)
			for _, line := range lines {
				if rType.MatchString(line) {
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
				carapace.ActionValues(methods...).Style(styles.Golang.Function).Tag("functions"),
				carapace.ActionValues(fields...).Style(styles.Golang.Field).Tag("fields"),
			).ToA()
		})
	})
}

// ActionSymbolTypes completes symbol types
//
//	t (static text segment symbol)
//	R (read-only data segment symbol)
func ActionSymbolTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"T", "text (code) segment symbol",
		"t", "static text segment symbol",
		"R", "read-only data segment symbol",
		"r", "static read-only data segment symbol",
		"D", "data segment symbol",
		"d", "static data segment symbol",
		"B", "bss segment symbol",
		"b", "static bss segment symbol",
		"C", "constant address",
		"U", "referenced but undefined symbol",
	)
}
