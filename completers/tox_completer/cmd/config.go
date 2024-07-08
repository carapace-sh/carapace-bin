package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"c"},
	Short:   "show tox configuration",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	add_common_flags(configCmd)
	add_common_subcommand_flags(configCmd)
	add_env_filtering_flags(configCmd)
	add_env_select_flag(configCmd)

	configCmd.Flags().Bool("core", false, "show core options (by default is hidden unless -e ALL is passed) (default: False)")
	configCmd.Flags().StringS("k", "k", "", "list just configuration keys specified (default: [])")

	configCmd.Flags().Bool("develop", false, "install package in development mode (default: False)")
	configCmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")
	configCmd.Flags().StringP("skip-missing-interpreters", "s", "config", "don't fail tests for missing interpreters: {config,true,false} choice (default: config)")

	carapace.Gen(configCmd).FlagCompletion(carapace.ActionMap{
		"skip-missing-interpreters": carapace.ActionValues("config", "true", "false"),
	})

	rootCmd.AddCommand(configCmd)
}
