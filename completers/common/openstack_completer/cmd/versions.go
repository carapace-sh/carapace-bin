package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionsCmd).Standalone()

	rootCmd.AddCommand(versionsCmd)
}
