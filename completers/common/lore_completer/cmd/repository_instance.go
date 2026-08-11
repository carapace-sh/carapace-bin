package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_instanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "Instance management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_instanceCmd).Standalone()

	repository_instanceCmd.Flags().BoolP("help", "h", false, "Print help")
	repositoryCmd.AddCommand(repository_instanceCmd)
}
