package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Generate shell script aliases for this project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_aliasCmd).Standalone()

	generate_aliasCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	generate_aliasCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	generate_aliasCmd.Flags().Bool("no-prefix", false, "Do not use a prefix for the generated aliases")
	generate_aliasCmd.Flags().StringP("prefix", "p", "", "Prefix for the generated aliases")
	generateCmd.AddCommand(generate_aliasCmd)

	carapace.Gen(generate_aliasCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
		// TODO environment
	})
}
