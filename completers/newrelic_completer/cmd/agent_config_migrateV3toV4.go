package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agent_config_migrateV3toV4Cmd = &cobra.Command{
	Use:   "migrateV3toV4",
	Short: "migrate V3 configuration to V4 configuration format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_config_migrateV3toV4Cmd).Standalone()
	agent_config_migrateV3toV4Cmd.Flags().Bool("overwrite", false, "if set ti true and pathOutput file exists already the old file is removed ")
	agent_config_migrateV3toV4Cmd.Flags().StringP("pathConfiguration", "c", "", "path configuration")
	agent_config_migrateV3toV4Cmd.Flags().StringP("pathDefinition", "d", "", "path definition")
	agent_config_migrateV3toV4Cmd.Flags().StringP("pathOutput", "o", "", "path output")
	agent_configCmd.AddCommand(agent_config_migrateV3toV4Cmd)

	carapace.Gen(agent_config_migrateV3toV4Cmd).FlagCompletion(carapace.ActionMap{
		"pathConfiguration": carapace.ActionFiles(),
		"pathDefinition":    carapace.ActionFiles(),
		"pathOutput":        carapace.ActionFiles(),
	})
}
