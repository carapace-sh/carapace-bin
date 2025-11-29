package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var generate_direnvCmd = &cobra.Command{
	Use:   "direnv",
	Short: "Generate a .envrc file that integrates direnv with this devbox project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_direnvCmd).Standalone()

	generate_direnvCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	generate_direnvCmd.PersistentFlags().StringP("env", "e", "", "environment variables to set in the devbox environment")
	generate_direnvCmd.PersistentFlags().String("env-file", "", "path to a file containing environment variables to set in the devbox environment")
	generate_direnvCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	generate_direnvCmd.Flags().BoolP("force", "f", false, "force overwrite existing files")
	generate_direnvCmd.Flags().BoolP("print-envrc", "p", false, "output contents of devbox configuration to use in .envrc")
	generate_direnvCmd.Flag("print-envrc").Hidden = true
	generateCmd.AddCommand(generate_direnvCmd)

	// TODO environment
	carapace.Gen(generate_direnvCmd).FlagCompletion(carapace.ActionMap{
		"config":   carapace.ActionDirectories(),
		"env":      env.ActionNameValues(false),
		"env-file": carapace.ActionFiles(),
	})
}
