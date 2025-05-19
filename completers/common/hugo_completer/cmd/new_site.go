package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var new_siteCmd = &cobra.Command{
	Use:   "site",
	Short: "Create a new site (skeleton)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(new_siteCmd).Standalone()
	new_siteCmd.Flags().Bool("force", false, "init inside non-empty directory")
	new_siteCmd.Flags().StringP("format", "f", "toml", "config & frontmatter format")
	newCmd.AddCommand(new_siteCmd)

	carapace.Gen(new_siteCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "toml", "yaml"),
	})

	carapace.Gen(new_siteCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
