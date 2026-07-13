package cmd

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions"
	spec "github.com/carapace-sh/carapace-spec"
	"github.com/spf13/cobra"
)

var macroCmd = &cobra.Command{
	Use:   "--macro [macro] [...args]",
	Short: "list or execute macros",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			printMacros()
		case 1:
			printMacro(args[0])
		default:
			mCmd := &cobra.Command{
				DisableFlagParsing: true,
			}
			carapace.Gen(mCmd).Standalone()
			carapace.Gen(mCmd).PositionalAnyCompletion(
				spec.NewAction([]string{"$carapace." + args[0]}).Parse(mCmd),
			)
			mCmd.SetArgs(append([]string{"_carapace", "export", ""}, args[1:]...))
			mCmd.SetOut(cmd.OutOrStdout())
			mCmd.SetErr(cmd.ErrOrStderr())
			mCmd.Execute()
		}
	},
}

func init() {
	carapace.Gen(macroCmd).Standalone()
	macroCmd.Flags().SetInterspersed(false)

	carapace.Gen(macroCmd).PositionalCompletion(
		carapace.ActionMultiPartsN("(", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionExecCommand("carapace", "--macro")(func(output []byte) carapace.Action {
					lines := strings.Split(string(output), "\n")

					vals := make([]string, 0)
					for _, line := range lines[:len(lines)-1] {
						if fields := strings.Fields(line); len(fields) > 1 {
							vals = append(vals, fields[0], strings.Join(fields[1:], " "))
						} else {
							vals = append(vals, fields[0], "")
						}
					}
					return carapace.ActionValuesDescribed(vals...).Invoke(carapace.Context{}).ToMultiPartsA(".")
				})
			default:
				if m, ok := actions.Macros[c.Parts[0]]; ok {
					return carapace.ActionValues(m.Signature() + ")")
				}
				return carapace.ActionValues(")")
			}
		}),
	)

	carapace.Gen(macroCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return spec.NewAction([]string{"$carapace." + c.Args[0]}).Parse(macroCmd).Shift(1) // TODO macro prefix
		}),
	)
}

func printMacros() {
	maxlen := 0
	names := make([]string, 0)
	for name := range actions.Macros {
		names = append(names, name)
		if len := len(name); len > maxlen {
			maxlen = len
		}
	}

	sort.Strings(names)
	for _, name := range names {
		description := ""
		if m, ok := actions.Macros[name]; ok {
			description = m.Description
		}
		fmt.Printf("%-"+strconv.Itoa(maxlen)+"v %v\n", name, description)
	}
}

func printMacro(name string) {
	// extract macro name from full spec (e.g. "color.XtermColorNames ||| $filter([Green])")
	macroName := strings.TrimSpace(strings.Split(name, " ||| ")[0])
	// TODO implement this in spec.Macro
	if m, ok := actions.Macros[macroName]; ok {
		path := strings.ReplaceAll(macroName, ".", "/")
		signature := ""
		if s := m.Signature(); s != "" {
			signature = fmt.Sprintf("(%v)", s)
		}

		fmt.Printf(`signature:   $carapace.%v%v
description: %v
reference:   https://pkg.go.dev/github.com/carapace-sh/carapace-bin/pkg/actions/%v#Action%v
`, macroName, signature, m.Description, filepath.Dir(path), filepath.Base(path))
	}
}
