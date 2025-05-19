package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload locally saved terminal session to asciinema.org",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uploadCmd).Standalone()

	rootCmd.AddCommand(uploadCmd)

	carapace.Gen(uploadCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
