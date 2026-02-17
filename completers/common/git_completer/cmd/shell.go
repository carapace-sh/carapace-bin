package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Restricted login shell for Git-only SSH access",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellCmd).Standalone()

	shellCmd.Flags().BoolS("c", "c", false, "run command")
	rootCmd.AddCommand(shellCmd)
}
