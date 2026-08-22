package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var findRevCmd = &cobra.Command{
	Use:   "find-rev",
	Short: "Translate between SVN revision numbers and tree-ish",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(findRevCmd).Standalone()

	findRevCmd.Flags().BoolP("after", "A", false, "Find closest match searching forward in history")
	findRevCmd.Flags().BoolP("before", "B", false, "Find commit corresponding to state at specified revision")
	rootCmd.AddCommand(findRevCmd)

	carapace.Gen(findRevCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
