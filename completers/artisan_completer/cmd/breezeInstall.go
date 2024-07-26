package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var breezeInstallCmd = &cobra.Command{
	Use:   "breeze:install [--dark] [--pest] [--ssr] [--typescript] [--composer [COMPOSER]] [--] <stack>",
	Short: "Install the Breeze controllers and resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(breezeInstallCmd).Standalone()

	breezeInstallCmd.Flags().String("composer", "", "Absolute path to the Composer binary which should be used to install packages")
	breezeInstallCmd.Flags().Bool("dark", false, "Indicate that dark mode support should be installed")
	breezeInstallCmd.Flags().Bool("pest", false, "Indicate that Pest should be installed")
	breezeInstallCmd.Flags().Bool("ssr", false, "Indicates if Inertia SSR support should be installed")
	breezeInstallCmd.Flags().Bool("typescript", false, "Indicates if TypeScript is preferred for the Inertia stack")
	rootCmd.AddCommand(breezeInstallCmd)
}
