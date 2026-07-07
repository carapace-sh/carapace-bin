package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var themeDumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Print current theme in loadable format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(themeDumpCmd).Standalone()

	themeCmd.AddCommand(themeDumpCmd)
}
