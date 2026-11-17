package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var terminfoRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the ssh guard, leaving the terminfo entry installed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminfoRemoveCmd).Standalone()

	terminfoRemoveCmd.Flags().Bool("dry-run", false, "report what would change without writing anything")
	terminfoRemoveCmd.Flags().Bool("json", false, "output what the run did as JSON")
	terminfoRemoveCmd.Flags().Bool("system", false, "remove the guard that every account on this machine shares")

	terminfoCmd.AddCommand(terminfoRemoveCmd)
}
