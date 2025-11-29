package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var version_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update devbox launcher and binary",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(version_updateCmd).Standalone()

	versionCmd.AddCommand(version_updateCmd)
}
