package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wrap_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "show installed and available versions of your projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wrap_statusCmd).Standalone()

	wrap_statusCmd.Flags().Bool("allow-insecure", false, "Allow insecure server connections.")
	applyCommonFlagCompletion(wrap_statusCmd)
	wrapCmd.AddCommand(wrap_statusCmd)
}
