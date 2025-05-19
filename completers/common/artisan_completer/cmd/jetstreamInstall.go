package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var jetstreamInstallCmd = &cobra.Command{
	Use:   "jetstream:install [--dark] [--teams] [--api] [--verification] [--pest] [--ssr] [--composer [COMPOSER]] [--] <stack>",
	Short: "Install the Jetstream components and resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(jetstreamInstallCmd).Standalone()

	jetstreamInstallCmd.Flags().Bool("api", false, "Indicates if API support should be installed")
	jetstreamInstallCmd.Flags().String("composer", "", "Absolute path to the Composer binary which should be used to install packages")
	jetstreamInstallCmd.Flags().Bool("dark", false, "Indicate that dark mode support should be installed")
	jetstreamInstallCmd.Flags().Bool("pest", false, "Indicates if Pest should be installed")
	jetstreamInstallCmd.Flags().Bool("ssr", false, "Indicates if Inertia SSR support should be installed")
	jetstreamInstallCmd.Flags().Bool("teams", false, "Indicates if team support should be installed")
	jetstreamInstallCmd.Flags().Bool("verification", false, "Indicates if email verification support should be installed")
	rootCmd.AddCommand(jetstreamInstallCmd)
}
