package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cryptCmd = &cobra.Command{
	Use:   "crypt",
	Short: "Use Charm encryption.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cryptCmd).Standalone()

	rootCmd.AddCommand(cryptCmd)
}
