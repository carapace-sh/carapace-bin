package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var terminfoSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Install Rex's terminfo entry and ssh guard",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminfoSetupCmd).Standalone()

	terminfoSetupCmd.Flags().Bool("dry-run", false, "report what would change without writing anything")
	terminfoSetupCmd.Flags().Bool("entry-only", false, "install the entry alone")
	terminfoSetupCmd.Flags().Bool("json", false, "output what the run did as JSON")
	terminfoSetupCmd.Flags().Bool("system", false, "install the entry and the guard for every account on this machine")

	terminfoCmd.AddCommand(terminfoSetupCmd)
}
