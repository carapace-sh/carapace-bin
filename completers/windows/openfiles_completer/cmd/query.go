package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "query and display all open files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queryCmd).Standalone()
	rootCmd.AddCommand(queryCmd)
}
