package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var review_templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Configure the template to use for review descriptions. This will list all available templates found in the repository and allow you to select one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(review_templateCmd).Standalone()

	review_templateCmd.Flags().BoolP("help", "h", false, "Print help")
	reviewCmd.AddCommand(review_templateCmd)
}
