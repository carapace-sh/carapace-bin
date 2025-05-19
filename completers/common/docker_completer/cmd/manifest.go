package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var manifestCmd = &cobra.Command{
	Use:     "manifest COMMAND",
	Short:   "Manage Docker image manifests and manifest lists",
	GroupID: "management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifestCmd).Standalone()

	rootCmd.AddCommand(manifestCmd)
}
