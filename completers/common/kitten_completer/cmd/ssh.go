package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Truly convenient SSH",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshCmd).Standalone()

	sshCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	rootCmd.AddCommand(sshCmd)

	carapace.Gen(sshCmd).PositionalAnyCompletion(
		bridge.ActionCarapace("ssh"),
	)
}
