package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var SshForEachCmd = &cobra.Command{
	Use:   "ssh-for-each",
	Short: "Ssh For Each",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(SshForEachCmd).Standalone()
	rootCmd.AddCommand(SshForEachCmd)
}
