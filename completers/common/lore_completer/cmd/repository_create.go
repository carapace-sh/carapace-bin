package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a repository in the given directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_createCmd).Standalone()

	repository_createCmd.Flags().String("description", "", "Optional description of repository")
	repository_createCmd.Flags().BoolP("help", "h", false, "Print help")
	repository_createCmd.Flags().String("id", "", "Optional ID of repository")
	repository_createCmd.Flags().String("shared-store-path", "", "Use this path rather than the system default as the shared store location")
	repository_createCmd.Flags().Bool("use-shared-store", false, "Use the shared store rather than create a local immutable store")
	repositoryCmd.AddCommand(repository_createCmd)
}
