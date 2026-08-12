package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Access the repository data store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_storeCmd).Standalone()

	repository_storeCmd.Flags().BoolP("help", "h", false, "Print help")
	repositoryCmd.AddCommand(repository_storeCmd)
}
