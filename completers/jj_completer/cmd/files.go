package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var filesCmd = &cobra.Command{
	Use:   "files [OPTIONS] [PATHS]...",
	Short: "List files in a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(filesCmd).Standalone()

	filesCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	filesCmd.Flags().StringP("revision", "r", "@", "The revision to list files in")
	rootCmd.AddCommand(filesCmd)

	carapace.Gen(filesCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(filesCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionRevFiles(filesCmd.Flag("revision").Value.String())
		}),
	)
}
