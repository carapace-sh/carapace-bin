package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Print the location of the global configuration file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_pathCmd).Standalone()

	config_pathCmd.Flags().Bool("dir", false, "Print the configuration directory instead of the configuration file.")
	configCmd.AddCommand(config_pathCmd)
}
