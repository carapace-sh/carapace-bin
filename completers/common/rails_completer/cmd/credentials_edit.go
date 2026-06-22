package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/common"
	"github.com/spf13/cobra"
)

var credentialsEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit encrypted credentials",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(credentialsEditCmd).Standalone()
	credentialsEditCmd.Flags().BoolP("help", "h", false, "Show help")

	common.AddEnvironmentFlag(credentialsEditCmd, "")

	carapace.Gen(credentialsEditCmd).PositionalCompletion(
		carapace.ActionValues().Usage("content path"),
	)
}
