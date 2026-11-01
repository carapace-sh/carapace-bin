package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var remote_help_setUrlCmd = &cobra.Command{
	Use:   "set-url",
	Short: "Set the fetch or push URL for an existing remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_help_setUrlCmd).Standalone()

	remote_helpCmd.AddCommand(remote_help_setUrlCmd)
}
