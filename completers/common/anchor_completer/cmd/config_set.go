package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/anchor"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set configuration settings in the local Anchor.toml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()

	config_setCmd.Flags().BoolP("help", "h", false, "Print help")
	config_setCmd.Flags().StringP("keypair", "k", "", "Path to wallet keypair file to update the Anchor.toml file with")
	config_setCmd.Flags().StringP("url", "u", "", "Cluster to connect to (custom URL). Use -um, -ud, -ut, -ul for standard clusters")
	configCmd.AddCommand(config_setCmd)

	carapace.Gen(config_setCmd).FlagCompletion(carapace.ActionMap{
		"keypair": carapace.ActionFiles(".json"),
		"url":     anchor.ActionClusters(),
	})
}
