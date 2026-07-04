package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "+ssh",
	Short: "Wrap ssh to automatically configure Ghostty terminal integration on remote hosts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshCmd).Standalone()
	sshCmd.Flags().SetInterspersed(false)

	sshCmd.Flags().Bool("cache", true, "Use the terminfo install cache")
	sshCmd.Flags().Bool("forward-env", true, "Enable TERM / SendEnv forwarding")
	sshCmd.Flags().Bool("help", false, "show help")
	sshCmd.Flags().String("ssh", "", "Path to the ssh binary")
	sshCmd.Flags().Bool("terminfo", true, "Install Ghostty terminfo on first connect")
	sshCmd.Flags().Bool("verbose", false, "Print +ssh status lines to stderr")
	rootCmd.AddCommand(sshCmd)

	carapace.Gen(sshCmd).FlagCompletion(carapace.ActionMap{
		"ssh": carapace.ActionFiles(),
	})

	carapace.Gen(sshCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("ssh"),
	)
}
