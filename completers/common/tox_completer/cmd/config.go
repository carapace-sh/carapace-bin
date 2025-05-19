package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
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

	configCmd.Flags().Bool("core", false, "show core options (by default is hidden unless -e ALL is passed) (default: False)")
	configCmd.Flags().Bool("develop", false, "install package in development mode (default: False)")
	configCmd.Flags().StringS("k", "k", "", "list just configuration keys specified (default: [])")
	configCmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")
	configCmd.Flags().StringP("skip-missing-interpreters", "s", "config", "don't fail tests for missing interpreters: {config,true,false} choice (default: config)")
	addCommonSubcommandFlags(configCmd)
	addEnvFilteringFlags(configCmd)
	addEnvSelectFlag(configCmd)
	rootCmd.AddCommand(configCmd)

	carapace.Gen(configCmd).FlagCompletion(carapace.ActionMap{
		"skip-missing-interpreters": carapace.ActionValues("config", "true", "false").StyleF(style.ForKeyword),
	})
}
