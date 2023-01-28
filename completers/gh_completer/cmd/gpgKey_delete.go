package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var gpgKey_deleteCmd = &cobra.Command{
	Use:   "delete <key-id>",
	Short: "Delete a GPG key from your GitHub account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gpgKey_deleteCmd).Standalone()

	gpgKey_deleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	gpgKeyCmd.AddCommand(gpgKey_deleteCmd)

	carapace.Gen(gpgKey_deleteCmd).PositionalCompletion(
		gh.ActionGpgKeys(gh.HostOpts{}),
	)
}
