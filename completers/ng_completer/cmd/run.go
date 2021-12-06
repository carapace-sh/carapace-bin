package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs an Architect target with an optional custom builder configuration defined in your project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().StringP("configuration", "c", "", "One or more named builder configurations")

	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"configuration": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Args) < 1 {
				return carapace.ActionValues()
			}
			return action.ActionBuilderConfigurations(c.Args[0]).Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	// TODO positional completion
}
