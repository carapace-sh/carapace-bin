package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var file_deleteCmd = &cobra.Command{
	Use:   "delete [OPTIONS] <FILESETS>...",
	Short: "Delete files from the given revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_deleteCmd).Standalone()

	file_deleteCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	file_deleteCmd.Flags().Bool("restore-descendants", false, "Preserve the content (not the diff) when rebasing descendants")
	file_deleteCmd.Flags().StringP("revision", "r", "@", "The revision to delete the files in")
	fileCmd.AddCommand(file_deleteCmd)

	carapace.Gen(file_deleteCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevsets(jj.RevOpts{}.Default()),
	})

	carapace.Gen(file_deleteCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionRevFiles(file_deleteCmd.Flag("revision").Value.String()).FilterArgs()
		}),
	)
}
