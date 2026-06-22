package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/common"
	"github.com/spf13/cobra"
)

var credentialsFetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch a credential value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(credentialsFetchCmd).Standalone()
	credentialsFetchCmd.Flags().BoolP("help", "h", false, "Show help")

	common.AddEnvironmentFlag(credentialsFetchCmd, "")

	carapace.Gen(credentialsFetchCmd).PositionalCompletion(
		carapace.ActionValues().Usage("key path (e.g. secret_key_base)"),
	)
}
