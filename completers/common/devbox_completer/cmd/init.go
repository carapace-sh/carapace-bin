package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [<dir>]",
	Short: "Initialize a directory as a devbox project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().Bool("auto", false, "Automatically detect packages to add")
	initCmd.Flags().Bool("dry-run", false, "Dry run for auto mode. Prints the config that would be used")
	initCmd.Flag("auto").Hidden = true
	initCmd.Flag("dry-run").Hidden = true
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
