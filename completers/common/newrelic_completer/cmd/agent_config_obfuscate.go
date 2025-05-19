package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agent_config_obfuscateCmd = &cobra.Command{
	Use:   "obfuscate",
	Short: "Obfuscate a configuration value using a key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_config_obfuscateCmd).Standalone()
	agent_config_obfuscateCmd.Flags().StringP("key", "k", "", "the key to use when obfuscating the clear-text value")
	agent_config_obfuscateCmd.Flags().StringP("value", "v", "", "the value, in clear text, to be obfuscated")
	agent_configCmd.AddCommand(agent_config_obfuscateCmd)
}
