package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/atuin"
	"github.com/spf13/cobra"
)

var scripts_newCmd = &cobra.Command{
	Use:   "new",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scripts_newCmd).Standalone()

	scripts_newCmd.Flags().StringP("description", "d", "", "")
	scripts_newCmd.Flags().BoolP("help", "h", false, "Print help")
	scripts_newCmd.Flags().String("last", "", "Use the last command as the script content Optionally specify a number to use the last N commands")
	scripts_newCmd.Flags().Bool("no-edit", false, "Skip opening editor when using --last")
	scripts_newCmd.Flags().String("script", "", "")
	scripts_newCmd.Flags().StringP("shebang", "s", "", "")
	scripts_newCmd.Flags().StringSliceP("tags", "t", nil, "")
	scriptsCmd.AddCommand(scripts_newCmd)

	carapace.Gen(scripts_newCmd).FlagCompletion(carapace.ActionMap{
		"tags": atuin.ActionTags(),
	})

	carapace.Gen(scripts_newCmd).FlagCompletion(carapace.ActionMap{
		"script": atuin.ActionScripts(),
	})
}
