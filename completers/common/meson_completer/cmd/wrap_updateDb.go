package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wrap_updateDbCmd = &cobra.Command{
	Use:   "update-db",
	Short: "Update list of projects available in WrapDB (Since 0.61.0)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wrap_updateDbCmd).Standalone()

	wrap_updateDbCmd.Flags().Bool("allow-insecure", false, "Allow insecure server connections.")
	applyCommonFlagCompletion(wrap_updateDbCmd)
	wrapCmd.AddCommand(wrap_updateDbCmd)
}
