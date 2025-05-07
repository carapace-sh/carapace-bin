package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Print used config path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_pathCmd).Standalone()

	config_pathCmd.Flags().Bool("json", false, "Display as JSON")
	configCmd.AddCommand(config_pathCmd)
}
