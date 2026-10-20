package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_marketplace_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a marketplace from a URL, path, or GitHub repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_marketplace_addCmd).Standalone()

	plugin_marketplace_addCmd.Flags().Bool("claudeai", false, "Add the marketplace of this name that claude.ai hosts for you, by its listed name or its local…")
	plugin_marketplace_addCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_marketplace_addCmd.Flags().Bool("json", false, "Print one machine-readable result line as the last line on stdout (same exit codes)")
	plugin_marketplace_addCmd.Flags().String("scope", "", "Where to declare the marketplace: user (default), project, or local")
	plugin_marketplace_addCmd.Flags().StringSlice("sparse", nil, "Limit checkout to specific directories via git sparse-checkout (for monorepos)")
	plugin_marketplaceCmd.AddCommand(plugin_marketplace_addCmd)

	carapace.Gen(plugin_marketplace_addCmd).FlagCompletion(carapace.ActionMap{
		"scope": carapace.ActionValues("user", "project", "local"),
	})

	carapace.Gen(plugin_marketplace_addCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
