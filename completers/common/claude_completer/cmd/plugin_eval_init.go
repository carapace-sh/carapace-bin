package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_eval_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Author an eval suite under the eval dir (evals/ unless --eval-dir or the manifest says otherwise)…",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_eval_initCmd).Standalone()

	plugin_eval_initCmd.Flags().Bool("bare", false, "Write a blank template (prompt.md + graders/criteria.md) instead of running the interview")
	plugin_eval_initCmd.Flags().String("eval-dir", "", "Directory (below the current directory) to write cases into")
	plugin_eval_initCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_eval_initCmd.Flags().BoolP("interactive", "i", false, "Run the authoring interview (already the default in a terminal); requires an interactive terminal")
	plugin_evalCmd.AddCommand(plugin_eval_initCmd)

	carapace.Gen(plugin_eval_initCmd).FlagCompletion(carapace.ActionMap{
		"eval-dir": carapace.ActionDirectories(),
	})
}
