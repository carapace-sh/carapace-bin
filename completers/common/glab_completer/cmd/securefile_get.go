package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securefile_getCmd = &cobra.Command{
	Use:     "get <fileID>",
	Short:   "Get details of a project secure file. (GitLab 18.0 and later)",
	Aliases: []string{"show"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securefile_getCmd).Standalone()

	securefileCmd.AddCommand(securefile_getCmd)

	// TODO positional completion
}
