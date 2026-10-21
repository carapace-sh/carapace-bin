package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "login on a machine via SSH",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshCmd).Standalone()

	sshCmd.Flags().Bool("now", false, "Do not acquire a lock before fetching the state")

	rootCmd.AddCommand(sshCmd)

	carapace.Gen(sshCmd).PositionalAnyCompletion(
		carapace.ActionValues(),
	)
}
