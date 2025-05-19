package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var nar_catCmd = &cobra.Command{
	Use:   "cat",
	Short: "print the contents of a file inside a NAR file on stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nar_catCmd).Standalone()

	narCmd.AddCommand(nar_catCmd)

	carapace.Gen(nar_catCmd).PositionalCompletion(
		carapace.ActionFiles(".nar"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return nix.ActionNarFileContents(c.Args[0])
		}),
	)
}
