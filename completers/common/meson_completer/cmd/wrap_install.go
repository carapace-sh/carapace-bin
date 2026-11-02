package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wrap_installCmd = &cobra.Command{
	Use:   "install",
	Short: "install the specified project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wrap_installCmd).Standalone()

	wrap_installCmd.Flags().Bool("allow-insecure", false, "Allow insecure server connections.")
	applyCommonFlagCompletion(wrap_installCmd)
	wrapCmd.AddCommand(wrap_installCmd)
}
