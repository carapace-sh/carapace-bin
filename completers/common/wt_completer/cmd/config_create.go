package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create configuration file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_createCmd).Standalone()

	config_createCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_createCmd.Flags().Bool("project", false, "Create project config (.config/wt.toml) instead of user config")
	configCmd.AddCommand(config_createCmd)
}
