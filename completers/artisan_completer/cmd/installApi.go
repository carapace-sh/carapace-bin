package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installApiCmd = &cobra.Command{
	Use:   "install:api [--composer [COMPOSER]] [--force] [--passport] [--without-migration-prompt]",
	Short: "Create an API routes file and install Laravel Sanctum or Laravel Passport",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installApiCmd).Standalone()

	installApiCmd.Flags().String("composer", "", "Absolute path to the Composer binary which should be used to install packages")
	installApiCmd.Flags().Bool("force", false, "Overwrite any existing API routes file")
	installApiCmd.Flags().Bool("passport", false, "Install Laravel Passport instead of Laravel Sanctum")
	installApiCmd.Flags().Bool("without-migration-prompt", false, "Do not prompt to run pending migrations")
	rootCmd.AddCommand(installApiCmd)
}
