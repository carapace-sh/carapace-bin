package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/common"
	"github.com/spf13/cobra"
)

var credentialsDiffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Enroll or disenroll from credential diffs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(credentialsDiffCmd).Standalone()
	credentialsDiffCmd.Flags().Bool("disenroll", false, "Disenroll from credential diffs")
	credentialsDiffCmd.Flags().Bool("enroll", false, "Enroll in credential diffs")
	credentialsDiffCmd.Flags().BoolP("help", "h", false, "Show help")

	common.AddEnvironmentFlag(credentialsDiffCmd, "")

	carapace.Gen(credentialsDiffCmd).PositionalCompletion(
		carapace.ActionValues().Usage("content path"),
	)
}
