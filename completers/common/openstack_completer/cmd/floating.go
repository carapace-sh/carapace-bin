package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var floatingCmd = &cobra.Command{
	Use:   "floating",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(floatingCmd).Standalone()

	rootCmd.AddCommand(floatingCmd)
}
