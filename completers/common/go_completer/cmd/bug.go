package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bugCmd = &cobra.Command{
	Use:   "bug",
	Short: "start a bug report",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bugCmd).Standalone()
	bugCmd.Flags().SetInterspersed(false)

	rootCmd.AddCommand(bugCmd)
}
