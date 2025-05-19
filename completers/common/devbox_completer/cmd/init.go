package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a directory as a devbox project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
