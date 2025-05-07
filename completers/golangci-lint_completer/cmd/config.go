package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config file information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.PersistentFlags().StringP("config", "c", "", "Read config from file path `PATH`")
	configCmd.PersistentFlags().Bool("no-config", false, "Don't read config file")
	rootCmd.AddCommand(configCmd)

	carapace.Gen(configCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
