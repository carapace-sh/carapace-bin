package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sparkInstallCmd = &cobra.Command{
	Use:   "spark:install",
	Short: "Install all of the Spark resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparkInstallCmd).Standalone()

	rootCmd.AddCommand(sparkInstallCmd)
}
