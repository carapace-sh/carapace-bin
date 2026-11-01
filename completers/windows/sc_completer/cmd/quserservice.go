package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quserserviceCmd = &cobra.Command{
	Use:   "quserservice",
	Short: "query for a local instance of a user service template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quserserviceCmd).Standalone()
	rootCmd.AddCommand(quserserviceCmd)
}
