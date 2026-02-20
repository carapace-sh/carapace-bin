package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_upload_anacondaCmd = &cobra.Command{
	Use:   "anaconda",
	Short: "Options for uploading to a Anaconda.org server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_upload_anacondaCmd).Standalone()

	help_uploadCmd.AddCommand(help_upload_anacondaCmd)
}
