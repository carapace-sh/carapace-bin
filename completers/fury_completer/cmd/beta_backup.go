package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var beta_backupCmd = &cobra.Command{
	Use:   "backup DIR",
	Short: "Save all files to a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(beta_backupCmd).Standalone()

	beta_backupCmd.Flags().String("kind", "", "Filter to one kind of package")
	betaCmd.AddCommand(beta_backupCmd)

	// TODO kind

	carapace.Gen(betaCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
