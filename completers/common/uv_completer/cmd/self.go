package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selfCmd = &cobra.Command{
	Use:   "self",
	Short: "Manage the uv executable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selfCmd).Standalone()

	rootCmd.AddCommand(selfCmd)
}
