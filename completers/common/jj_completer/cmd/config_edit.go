package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_editCmd = &cobra.Command{
	Use:     "edit",
	Short:   "Start an editor on a jj config file",
	Aliases: []string{"e"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_editCmd).Standalone()

	config_editCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_editCmd.Flags().Bool("repo", false, "Target the repo-level config")
	config_editCmd.Flags().Bool("user", false, "Target the user-level config")
	config_editCmd.Flags().Bool("workspace", false, "Target the workspace-level config")
	configCmd.AddCommand(config_editCmd)
}
