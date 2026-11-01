package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var actionsCmd = &cobra.Command{
	Use:   "actions",
	Short: "List the actions a client or the server offers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(actionsCmd).Standalone()

	actionsCmd.Flags().Bool("json", false, "output actions as JSON")

	rootCmd.AddCommand(actionsCmd)
}
