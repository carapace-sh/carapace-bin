package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var manCmd = &cobra.Command{
	Use:    "man",
	Short:  "Generates GoReleaser's command line manpages",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manCmd).Standalone()

	rootCmd.AddCommand(manCmd)
}
