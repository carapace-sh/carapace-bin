package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chown",
	Short: "change file owner and group",
	Long:  "https://keith.github.io/xcode-manpages/chown.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("H", "H", false, "If the file is a symbolic link, change the user and group of the link itself")
	rootCmd.Flags().BoolS("L", "L", false, "Traverse all symbolic links to directories")
	rootCmd.Flags().BoolS("P", "P", false, "Do not traverse any symbolic links to directories")
	rootCmd.Flags().BoolS("R", "R", false, "Recursively change the user and group of each file")
	rootCmd.Flags().BoolS("h", "h", false, "If the file is a symbolic link, change the user and group of the link itself")
	rootCmd.Flags().BoolS("n", "n", false, "Do not recursively change the user and group of files")
	rootCmd.Flags().BoolS("v", "v", false, "Cause chown to be verbose, showing files as the owner is modified")

	carapace.Gen(rootCmd).PositionalCompletion(
		os.ActionUserGroup(),
		carapace.ActionFiles(),
	)
}
