package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "Manage the npm configuration files",
	Aliases: []string{"c"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()
	configCmd.PersistentFlags().Bool("json", false, "output as json")
	configCmd.PersistentFlags().String("editor", "", "editor to use for config edit")
	configCmd.PersistentFlags().BoolP("global", "g", false, "operate in global mode")
	configCmd.PersistentFlags().StringP("location", "L", "", "config file location")
	configCmd.PersistentFlags().BoolP("long", "l", false, "show extended information")

	rootCmd.AddCommand(configCmd)

	carapace.Gen(configCmd).FlagCompletion(carapace.ActionMap{
		"location": carapace.ActionValues("global", "user", "project"),
	})
}
