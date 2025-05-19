package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var validateConfigCmd = &cobra.Command{
	Use:   "+validate-config",
	Short: "validate a Ghostty config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validateConfigCmd).Standalone()

	validateConfigCmd.Flags().String("config-file", "", "validate a specific target config file")
	validateConfigCmd.Flags().Bool("help", false, "show help")
	rootCmd.AddCommand(validateConfigCmd)

	carapace.Gen(validateConfigCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
	})
}
