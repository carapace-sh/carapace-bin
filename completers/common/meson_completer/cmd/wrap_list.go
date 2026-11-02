package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wrap_listCmd = &cobra.Command{
	Use:   "list",
	Short: "show all available projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wrap_listCmd).Standalone()

	wrap_listCmd.Flags().Bool("allow-insecure", false, "Allow insecure server connections.")
	applyCommonFlagCompletion(wrap_listCmd)
	wrapCmd.AddCommand(wrap_listCmd)
}
