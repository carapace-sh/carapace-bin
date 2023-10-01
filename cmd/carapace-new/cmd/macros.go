package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions"
	spec "github.com/rsteube/carapace-spec"
	"github.com/spf13/cobra"
)

var macrosCmd = &cobra.Command{
	Use:   "--macros",
	Short: "list spec macros",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macrosCmd).Standalone()
	macrosCmd.Flags().SetInterspersed(false)

	carapace.Gen(macrosCmd).PositionalCompletion(
		ActionMacros(),
	)

	subcommands[macrosCmd.Name()] = macrosCmd

	carapace.Gen(macrosCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return spec.ActionMacro("$_" + c.Args[0]).Shift(1) // not fully working for bridge
		}),
	)
}

func ActionMacros() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0)
		for name := range actions.MacroMap {
			vals = append(vals, name, actions.MacroDescriptions[name])
		}
		return carapace.ActionValuesDescribed(vals...).MultiParts(".")
	})
}
