package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var auth_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display identity information for the current user or specified user IDs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_infoCmd).Standalone()

	auth_infoCmd.Flags().BoolP("help", "h", false, "Print help")
	auth_infoCmd.Flags().Bool("with-token", false, "Include cached tokens in the output")
	authCmd.AddCommand(auth_infoCmd)

	carapace.Gen(auth_infoCmd).PositionalAnyCompletion(
		action.ActionIdentities(auth_infoCmd),
	)
}
