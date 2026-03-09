package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var markCmd = &cobra.Command{
	Use:   "mark <subcommand> [options] [<group-id>] <package-spec>...",
	Short: "change the reason of an installed package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(markCmd).Standalone()

	rootCmd.AddCommand(markCmd)
}
