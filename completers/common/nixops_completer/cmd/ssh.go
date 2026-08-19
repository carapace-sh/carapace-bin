package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var SshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Ssh",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(SshCmd).Standalone()
	rootCmd.AddCommand(SshCmd)
}
