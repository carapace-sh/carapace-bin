package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonInstallCmd = &cobra.Command{
	Use:   "horizon:install",
	Short: "Install all of the Horizon resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonInstallCmd).Standalone()

	rootCmd.AddCommand(horizonInstallCmd)
}
