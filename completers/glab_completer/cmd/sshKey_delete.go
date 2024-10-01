package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sshKey_deleteCmd = &cobra.Command{
	Use:   "delete <key-id>",
	Short: "Deletes a single SSH key specified by the ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshKey_deleteCmd).Standalone()

	sshKey_deleteCmd.Flags().StringP("page", "p", "", "Page number.")
	sshKey_deleteCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	sshKeyCmd.AddCommand(sshKey_deleteCmd)
}
