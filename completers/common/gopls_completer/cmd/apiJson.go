package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apiJsonCmd = &cobra.Command{
	Use:   "api-json",
	Short: "print json describing gopls API",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiJsonCmd).Standalone()

	rootCmd.AddCommand(apiJsonCmd)
}
