package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var pr_templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Configure the template to use for review descriptions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_templateCmd).Standalone()

	pr_templateCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	prCmd.AddCommand(pr_templateCmd)

	carapace.Gen(pr_templateCmd).PositionalCompletion(
		carapace.ActionFiles().ChdirF(traverse.GitWorkTree),
	)
}
