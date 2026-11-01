package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "report available sleep states",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aCmd).Standalone()
	rootCmd.AddCommand(aCmd)
}
