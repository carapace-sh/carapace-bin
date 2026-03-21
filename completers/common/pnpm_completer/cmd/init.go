package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a package.json file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().Bool("bare", false, "Creates a package.json with only the required fields")
	initCmd.Flags().BoolP("help", "h", false, "Output usage information")
	initCmd.Flags().Bool("init-package-manager", false, "Pin the project to the current pnpm version via the packageManager field")
	initCmd.Flags().String("init-type", "", "Set the module system (commonjs or module)")

	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"init-type": carapace.ActionValues("commonjs", "module"),
	})
}
