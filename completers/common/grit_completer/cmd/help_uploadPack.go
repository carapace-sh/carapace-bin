package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_uploadPackCmd = &cobra.Command{
	Use:    "upload-pack",
	Short:  "Serve a fetch or clone over stdin/stdout (run by ssh and grit-http-server)",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_uploadPackCmd).Standalone()

	helpCmd.AddCommand(help_uploadPackCmd)
}
