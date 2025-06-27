package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var file_listCmd = &cobra.Command{
	Use:   "files [OPTIONS] [PATHS]...",
	Short: "List files in a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_listCmd).Standalone()

	file_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	file_listCmd.Flags().StringP("revision", "r", "@", "The revision to list files in")
	file_listCmd.Flags().StringP("template", "T", "", "Render each file entry using the given template")
	fileCmd.AddCommand(file_listCmd)

	carapace.Gen(file_listCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(file_listCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionRevFiles(file_listCmd.Flag("revision").Value.String())
		}),
	)
}
