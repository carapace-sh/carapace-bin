package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_pathCmd = &cobra.Command{
	Use:     "path",
	Short:   "Print the paths to the config files",
	Aliases: []string{"p"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_pathCmd).Standalone()

	config_pathCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_pathCmd.Flags().Bool("repo", false, "Target the repo-level config")
	config_pathCmd.Flags().Bool("user", false, "Target the user-level config")
	config_pathCmd.Flags().Bool("workspace", false, "Target the workspace-level config")
	configCmd.AddCommand(config_pathCmd)
}
