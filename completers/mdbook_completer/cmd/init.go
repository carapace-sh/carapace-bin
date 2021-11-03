package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates the boilerplate structure and files for a new book",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().Bool("force", false, "Skips confirmation prompts")
	initCmd.Flags().BoolP("help", "h", false, "Prints help information")
	initCmd.Flags().Bool("theme", false, "Copies the default theme into your source folder")
	initCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
