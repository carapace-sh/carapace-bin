package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "archiveutil",
	Short: "command-line interface for the Archive Utility application",
	Long:  "https://man.freebsd.org/cgi/man.cgi?archiveutil",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionValues("extractXcodeAEA"))
}
