package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var templates_pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Output the path to the templates directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(templates_pathCmd).Standalone()

	templatesCmd.AddCommand(templates_pathCmd)
}
