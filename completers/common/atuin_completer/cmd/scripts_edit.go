package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/atuin"
	"github.com/spf13/cobra"
)

var scripts_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scripts_editCmd).Standalone()

	scripts_editCmd.Flags().StringP("description", "d", "", "")
	scripts_editCmd.Flags().BoolP("help", "h", false, "Print help")
	scripts_editCmd.Flags().Bool("no-edit", false, "Skip opening editor")
	scripts_editCmd.Flags().Bool("no-tags", false, "Remove all tags from the script")
	scripts_editCmd.Flags().String("rename", "", "Rename the script")
	scripts_editCmd.Flags().String("script", "", "")
	scripts_editCmd.Flags().StringP("shebang", "s", "", "")
	scripts_editCmd.Flags().StringSliceP("tags", "t", nil, "Replace all existing tags with these new tags")
	scriptsCmd.AddCommand(scripts_editCmd)

	carapace.Gen(scripts_editCmd).FlagCompletion(carapace.ActionMap{
		"script": atuin.ActionScripts(),
		"tags":   atuin.ActionTags(),
	})
}
