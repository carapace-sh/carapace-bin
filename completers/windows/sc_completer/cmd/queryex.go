package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queryexCmd = &cobra.Command{
	Use:   "queryex",
	Short: "display extended service status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queryexCmd).Standalone()
	rootCmd.AddCommand(queryexCmd)
}
