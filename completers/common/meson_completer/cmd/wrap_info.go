package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wrap_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "show available versions of a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wrap_infoCmd).Standalone()

	wrap_infoCmd.Flags().Bool("allow-insecure", false, "Allow insecure server connections.")
	applyCommonFlagCompletion(wrap_infoCmd)
	wrapCmd.AddCommand(wrap_infoCmd)
}
