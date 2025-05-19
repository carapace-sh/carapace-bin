package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var manCmd = &cobra.Command{
	Use:    "man",
	Short:  "Generate man pages",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manCmd).Standalone()

	rootCmd.AddCommand(manCmd)
}
