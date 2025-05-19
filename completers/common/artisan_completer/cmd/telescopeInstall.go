package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var telescopeInstallCmd = &cobra.Command{
	Use:   "telescope:install",
	Short: "Install all of the Telescope resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(telescopeInstallCmd).Standalone()

	rootCmd.AddCommand(telescopeInstallCmd)
}
