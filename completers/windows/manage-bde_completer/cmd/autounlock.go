package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autounlockCmd = &cobra.Command{
	Use:   "autounlock",
	Short: "enable or disable automatic unlocking",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autounlockCmd).Standalone()
	rootCmd.AddCommand(autounlockCmd)
}
