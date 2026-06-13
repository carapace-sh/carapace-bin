package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cd",
	Short: "Change the shell working directory",
	Long:  "https://man7.org/linux/man-pages/man1/cd.1p.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("@", "@", false, "present a file with extended attributes as a directory containing the file attributes")
	rootCmd.Flags().BoolS("L", "L", false, "force symbolic links to be followed")
	rootCmd.Flags().BoolS("P", "P", false, "use the physical directory structure without following symbolic links")
	rootCmd.Flags().BoolS("e", "e", false, "if the -P option is supplied, exit with a non-zero status if the current working directory cannot be determined successfully")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
