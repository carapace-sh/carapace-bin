package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_readmeCmd = &cobra.Command{
	Use:   "readme [filename]",
	Short: "Generate markdown readme file for this project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_readmeCmd).Standalone()

	generate_readmeCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	generate_readmeCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	generate_readmeCmd.Flags().Bool("save-template", false, "Save default template for the README file")
	generate_readmeCmd.Flags().StringP("template", "t", "", "Path to a custom template for the README file")
	generateCmd.AddCommand(generate_readmeCmd)

	// TODO environment
	carapace.Gen(generate_readmeCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})

	carapace.Gen(generate_readmeCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
