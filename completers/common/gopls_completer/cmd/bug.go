package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bugCmd = &cobra.Command{
	Use:   "bug",
	Short: "report a bug in gopls",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bugCmd).Standalone()

	rootCmd.AddCommand(bugCmd)
}
