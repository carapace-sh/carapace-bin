package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link [code]",
	Short: "Link multiple machines to your Charm account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkCmd).Standalone()

	rootCmd.AddCommand(linkCmd)
}
