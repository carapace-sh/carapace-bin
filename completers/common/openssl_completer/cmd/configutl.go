package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configutlCmd = &cobra.Command{
	Use:     "configutl",
	Short:   "openssl configuration file helper command",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configutlCmd)

	configutlCmd.Flags().StringS("config", "config", "", "Config file to deal with (the default one if omitted)")
	configutlCmd.Flags().BoolS("noheader", "noheader", false, "Don't print the information about original config")
	configutlCmd.Flags().StringS("out", "out", "", "Output to filename rather than stdout")
	rootCmd.AddCommand(configutlCmd)

	carapace.Gen(configutlCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"out":    carapace.ActionFiles(),
	})
}
