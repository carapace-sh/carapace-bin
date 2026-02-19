package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Truly convenient SSH",
	Long:  "The ssh kitten is a thin wrapper around the ssh command. It automatically enables shell integration on the remote host, re-uses existing connections to reduce latency, makes the kitty terminfo database available, etc. Its invocation is identical to the ssh command. For details on its usage, see Truly convenient SSH.",
}

func init() {
	rootCmd.AddCommand(sshCmd)
	carapace.Gen(sshCmd).Standalone()

	sshCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(sshCmd).PositionalAnyCompletion(
		bridge.ActionCarapace("ssh"),
	)
}
