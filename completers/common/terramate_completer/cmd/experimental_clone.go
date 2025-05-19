package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var experimental_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clones a stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(experimental_cloneCmd).Standalone()

	experimental_cloneCmd.Flags().Bool("skip-child-stacks", false, "Clone ignores child stacks")
	experimentalCmd.AddCommand(experimental_cloneCmd)

	carapace.Gen(experimental_cloneCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		carapace.ActionDirectories(),
	)
}
