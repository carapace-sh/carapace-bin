package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var UnlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Unlock",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(UnlockCmd).Standalone()
	rootCmd.AddCommand(UnlockCmd)
}
