package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_review_templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Configure the template to use for review descriptions. This will list all available templates found in the repository and allow you to select one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_review_templateCmd).Standalone()

	help_reviewCmd.AddCommand(help_review_templateCmd)
}
