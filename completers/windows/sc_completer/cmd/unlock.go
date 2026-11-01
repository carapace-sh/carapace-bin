package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "unlock the service database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unlockCmd).Standalone()
	rootCmd.AddCommand(unlockCmd)
}
