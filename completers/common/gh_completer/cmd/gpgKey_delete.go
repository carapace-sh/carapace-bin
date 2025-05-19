package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var gpgKey_deleteCmd = &cobra.Command{
	Use:   "delete <key-id>",
	Short: "Delete a GPG key from your GitHub account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gpgKey_deleteCmd).Standalone()

	gpgKey_deleteCmd.Flags().Bool("confirm", false, "Skip the confirmation prompt")
	gpgKey_deleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	gpgKey_deleteCmd.Flag("confirm").Hidden = true
	gpgKeyCmd.AddCommand(gpgKey_deleteCmd)

	carapace.Gen(gpgKey_deleteCmd).PositionalCompletion(
		gh.ActionGpgKeys(gh.HostOpts{}),
	)
}
