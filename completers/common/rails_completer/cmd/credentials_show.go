package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/common"
	"github.com/spf13/cobra"
)

var credentialsShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show decrypted credentials",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(credentialsShowCmd).Standalone()
	credentialsShowCmd.Flags().BoolP("help", "h", false, "Show help")

	common.AddEnvironmentFlag(credentialsShowCmd, "")
}
