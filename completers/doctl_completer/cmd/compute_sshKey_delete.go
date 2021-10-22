package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_sshKey_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Permanently delete an SSH key from your account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_sshKey_deleteCmd).Standalone()
	compute_sshKey_deleteCmd.Flags().BoolP("force", "f", false, "Delete the key without a confirmation prompt")
	compute_sshKeyCmd.AddCommand(compute_sshKey_deleteCmd)
}
