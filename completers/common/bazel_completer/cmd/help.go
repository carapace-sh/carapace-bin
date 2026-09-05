package cmd

import (
	"github.com/carapace-sh/carapace"
	bazel "github.com/carapace-sh/carapace-bin/pkg/actions/tools/bazel"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Prints help for commands, or the index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	helpCmd.Flags().String("help_verbosity", "", "Select the verbosity of the help command.")
	helpCmd.Flags().StringP("long", "l", "", "Show full description of each option, instead of just its name.")
	helpCmd.Flags().String("short", "", "Show only the names of the options, not their types or meanings.")
	rootCmd.AddCommand(helpCmd)

	carapace.Gen(helpCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				bazel.ActionCommands(),
				carapace.ActionValues("startup_options", "target-syntax", "info-keys"),
			).ToA()
		}),
	)
}
