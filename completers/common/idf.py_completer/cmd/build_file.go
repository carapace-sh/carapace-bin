package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildFileCmd = &cobra.Command{
	Use:   "build-file",
	Short: "Build a standalone C file without project boilerplate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildFileCmd).Standalone()
	buildFileCmd.Flags().Bool("clean-cache", false, "Remove the cached container project before building")
	carapace.Gen(buildFileCmd).PositionalCompletion(carapace.ActionFiles(".c", ".cpp"))
	rootCmd.AddCommand(buildFileCmd)
}
