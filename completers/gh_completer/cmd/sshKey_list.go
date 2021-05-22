package cmd

import (
	"github.com/spf13/cobra"
)

var sshKey_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists SSH keys in your GitHub account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	sshKeyCmd.AddCommand(sshKey_listCmd)
}
