package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fragmentsCmd = &cobra.Command{
	Use:   "fragments",
	Short: "Manage fragments that are stored in the database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fragmentsCmd).Standalone()

	rootCmd.AddCommand(fragmentsCmd)
}
