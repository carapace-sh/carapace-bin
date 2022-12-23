package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var sshKey_deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete an SSH key from your GitHub account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshKey_deleteCmd).Standalone()

	sshKey_deleteCmd.Flags().BoolP("confirm", "y", false, "Skip the confirmation prompt")
	sshKeyCmd.AddCommand(sshKey_deleteCmd)

	carapace.Gen(sshKey_deleteCmd).PositionalCompletion(
		gh.ActionSshKeys(gh.HostOpts{}),
	)
}
