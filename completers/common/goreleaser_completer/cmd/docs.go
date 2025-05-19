package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docsCmd = &cobra.Command{
	Use:    "docs",
	Short:  "Generates GoReleaser's command line docs",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docsCmd).Standalone()

	rootCmd.AddCommand(docsCmd)
}
