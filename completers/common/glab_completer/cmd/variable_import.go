package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var variable_importCmd = &cobra.Command{
	Use:     "import",
	Short:   "Import variables from a JSON file or standard input.",
	Aliases: []string{"im"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variable_importCmd).Standalone()

	variable_importCmd.PersistentFlags().StringP("group", "g", "", "Select a group or subgroup. Ignored if a repository argument is set.")
	variable_importCmd.Flags().StringP("input-file", "i", "", "Read the variables JSON from this file instead of standard input.")
	variable_importCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	variable_importCmd.Flags().Bool("skip-existing", false, "Skip variables that already exist instead of failing.")
	variableCmd.AddCommand(variable_importCmd)

	carapace.Gen(variable_importCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(variable_importCmd),
	})
}
