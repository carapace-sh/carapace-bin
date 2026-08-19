package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sshForEachCmd = &cobra.Command{
	Use:   "ssh-for-each",
	Short: "execute a command on each machine via SSH",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshForEachCmd).Standalone()
	rootCmd.AddCommand(sshForEachCmd)
}
