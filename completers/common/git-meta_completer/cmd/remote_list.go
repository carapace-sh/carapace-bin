package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var remote_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured metadata remotes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_listCmd).Standalone()

	remote_listCmd.Flags().BoolP("help", "h", false, "Print help")
	remoteCmd.AddCommand(remote_listCmd)
}
