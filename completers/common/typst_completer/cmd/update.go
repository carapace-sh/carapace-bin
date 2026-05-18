package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Self update the Typst CLI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().String("backup-path", "", "Custom path to the backup file created on update and used by `--revert`, defaults to system-dependent location")
	updateCmd.Flags().Bool("force", false, "Forces a downgrade to an older version (required for downgrading)")
	updateCmd.Flags().Bool("revert", false, "Reverts to the version from before the last update (only possible if `typst update` has previously ran)")
	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"backup-path": carapace.ActionFiles(),
	})
}
