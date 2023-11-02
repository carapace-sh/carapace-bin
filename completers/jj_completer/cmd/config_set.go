package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:     "set",
	Short:   "Update config file to set the given option to a given value",
	Aliases: []string{"s"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()

	config_setCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_setCmd.Flags().Bool("repo", false, "Target the repo-level config")
	config_setCmd.Flags().Bool("user", false, "Target the user-level config")
	configCmd.AddCommand(config_setCmd)
}
