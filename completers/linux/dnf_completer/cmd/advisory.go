package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var advisoryCmd = &cobra.Command{
	Use:   "advisory <subcommand> [options] [<advisory-spec>...]",
	Short: "manage advisories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(advisoryCmd).Standalone()

	rootCmd.AddCommand(advisoryCmd)
}
