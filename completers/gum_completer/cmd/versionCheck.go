package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCheckCmd = &cobra.Command{
	Use:   "version-check",
	Short: "Semver check current gum version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCheckCmd).Standalone()

	rootCmd.AddCommand(versionCheckCmd)
}
