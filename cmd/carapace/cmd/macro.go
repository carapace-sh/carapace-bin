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
	"github.com/carapace-sh/carapace/pkg/execlog"
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
			mArgs := []string{"_carapace", "macro"}
			mArgs = append(mArgs, args...)
			mCmd := execlog.Command("carapace", mArgs...)
			mCmd.Stdout = cmd.OutOrStdout()
			mCmd.Stderr = cmd.ErrOrStderr()
			mCmd.Run()
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
			return spec.ActionMacro("$carapace." + c.Args[0]).Shift(1) // TODO macro prefix
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
	// TODO implement this in spec.Macro
	if m, ok := actions.Macros[name]; ok {
		path := strings.Replace(name, ".", "/", -1)
		signature := ""
		if s := m.Signature(); s != "" {
			signature = fmt.Sprintf("(%v)", s)
		}

		fmt.Printf(`signature:   $carapace.%v%v
description: %v
reference:   https://pkg.go.dev/github.com/carapace-sh/carapace-bin/pkg/actions/%v#Action%v
`, name, signature, m.Description, filepath.Dir(path), filepath.Base(path))
	}
}
