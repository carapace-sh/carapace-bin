package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "Manage Python versions and installations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pythonCmd).Standalone()

	rootCmd.AddCommand(pythonCmd)
}
