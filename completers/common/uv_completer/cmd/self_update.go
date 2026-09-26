package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var self_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update uv",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(self_updateCmd).Standalone()

	self_updateCmd.Flags().Bool("dry-run", false, "Run without performing the update")
	self_updateCmd.Flags().String("token", "", "A GitHub token for authentication. A token is not required but can be used to reduce the chance of encountering rate limits")
	selfCmd.AddCommand(self_updateCmd)
}
