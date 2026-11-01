package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_updatePathCmd = &cobra.Command{
	Use:   "update-path",
	Short: "Update the stored path for this instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_updatePathCmd).Standalone()

	repository_updatePathCmd.Flags().BoolP("help", "h", false, "Print help")
	repositoryCmd.AddCommand(repository_updatePathCmd)
}
