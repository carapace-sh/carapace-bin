package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resume server(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_resumeCmd).Standalone()

	serverCmd.AddCommand(server_resumeCmd)
}
