package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var templatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "Manage stored prompt templates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(templatesCmd).Standalone()

	rootCmd.AddCommand(templatesCmd)
}
