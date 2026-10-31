package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var templates_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the specified prompt template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(templates_showCmd).Standalone()

	templatesCmd.AddCommand(templates_showCmd)
}
