package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_autolink_createCmd = &cobra.Command{
	Use:     "create <keyPrefix> <urlTemplate>",
	Short:   "Create a new autolink reference",
	Aliases: []string{"new"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_autolink_createCmd).Standalone()

	repo_autolink_createCmd.Flags().BoolP("numeric", "n", false, "Mark autolink as numeric")
	repo_autolinkCmd.AddCommand(repo_autolink_createCmd)
}
