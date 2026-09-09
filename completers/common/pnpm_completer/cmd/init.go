package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	initCmd.Flags().Bool("init-package-manager", false, "Pin the latest pnpm version in package.json, through \"devEngines.packageManager\" and \"packageManager\", and auto-download pnpm when it is missing")
	initCmd.Flags().String("init-type", "", "Set the module system for the package. Defaults to \"module\"")
	initCmd.Flags().Bool("no-init-package-manager", false, "Scaffold the manifest without a pnpm version pin")

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"init-type": carapace.ActionValues("commonjs", "module"),
	})

	rootCmd.AddCommand(initCmd)
}
