package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var file_editCmd = &cobra.Command{
	Use:   "edit [OPTIONS] <FILE>",
	Short: "Edit the contents of a file in a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_editCmd).Standalone()

	file_editCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	file_editCmd.Flags().Bool("restore-descendants", false, "Preserve the content (not the diff) when rebasing descendants")
	file_editCmd.Flags().StringP("revision", "r", "@", "The revision to edit the file in")
	fileCmd.AddCommand(file_editCmd)

	carapace.Gen(file_editCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevsets(jj.RevOpts{}.Default()),
	})

	carapace.Gen(file_editCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionRevFiles(file_editCmd.Flag("revision").Value.String())
		}),
	)
}
