package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "Sync plus rebase the current branch onto the updated p4 remote branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebaseCmd).Standalone()

	rebaseCmd.Flags().Bool("import-labels", false, "Import p4 labels")
	rootCmd.AddCommand(rebaseCmd)
}
