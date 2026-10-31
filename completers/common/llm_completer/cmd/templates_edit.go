package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var templates_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the specified prompt template using the default $EDITOR",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(templates_editCmd).Standalone()

	templatesCmd.AddCommand(templates_editCmd)
}
