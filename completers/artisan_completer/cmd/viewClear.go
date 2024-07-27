package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var viewClearCmd = &cobra.Command{
	Use:   "view:clear",
	Short: "Clear all compiled view files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(viewClearCmd).Standalone()

	rootCmd.AddCommand(viewClearCmd)
}
