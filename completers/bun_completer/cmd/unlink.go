package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unlinkCmd = &cobra.Command{
	Use:   "unlink",
	Short: "Globally unlink an npm package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unlinkCmd).Standalone()

	rootCmd.AddCommand(unlinkCmd)
}
