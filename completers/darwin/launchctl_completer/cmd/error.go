package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var errorCmd = &cobra.Command{
	Use:   "error",
	Short: "Print a description of an error",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(errorCmd).Standalone()
	rootCmd.AddCommand(errorCmd)
}
