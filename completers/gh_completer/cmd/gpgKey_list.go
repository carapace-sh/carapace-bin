package cmd

import (
	"github.com/spf13/cobra"
)

var gpgKey_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists GPG keys in your GitHub account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	gpgKeyCmd.AddCommand(gpgKey_listCmd)
}
