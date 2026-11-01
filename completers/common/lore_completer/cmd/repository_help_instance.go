package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_instanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "Instance management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_instanceCmd).Standalone()

	repository_helpCmd.AddCommand(repository_help_instanceCmd)
}
