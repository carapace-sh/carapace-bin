package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chgrp",
	Short: "change group ownership of a file",
	Long:  "https://keith.github.io/xcode-manpages/chgrp.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("H", "H", false, "If the file is a symbolic link, change the group of the link itself")
	rootCmd.Flags().BoolS("L", "L", false, "Traverse all symbolic links to directories")
	rootCmd.Flags().BoolS("P", "P", false, "Do not traverse any symbolic links to directories")
	rootCmd.Flags().BoolS("R", "R", false, "Recursively change the group of each file")
	rootCmd.Flags().BoolS("h", "h", false, "If the file is a symbolic link, change the group of the link itself")
	rootCmd.Flags().BoolS("v", "v", false, "Cause chgrp to be verbose, showing files as the group is modified")

	carapace.Gen(rootCmd).PositionalCompletion(
		os.ActionGroups(),
		carapace.ActionFiles(),
	)
}
