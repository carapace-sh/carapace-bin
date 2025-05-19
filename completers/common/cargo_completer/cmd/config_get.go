package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getCmd).Standalone()

	config_getCmd.Flags().String("format", "", "Display format")
	config_getCmd.Flags().BoolP("help", "h", false, "Print help")
	config_getCmd.Flags().String("merged", "", "Whether or not to merge config values")
	config_getCmd.Flags().Bool("show-origin", false, "Display where the config value is defined")
	configCmd.AddCommand(config_getCmd)

	carapace.Gen(config_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("toml", "json", "json-value"),
		"merged": carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
	})

	// TODO config keys
}
