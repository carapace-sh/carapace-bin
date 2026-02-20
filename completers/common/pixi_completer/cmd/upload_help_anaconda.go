package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_help_anacondaCmd = &cobra.Command{
	Use:   "anaconda",
	Short: "Options for uploading to a Anaconda.org server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_help_anacondaCmd).Standalone()

	upload_helpCmd.AddCommand(upload_help_anacondaCmd)
}
