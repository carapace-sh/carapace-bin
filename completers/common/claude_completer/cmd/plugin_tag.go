package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Create a {name}--v{version} git tag for a plugin release, validating that plugin.json and any…",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_tagCmd).Standalone()

	plugin_tagCmd.Flags().Bool("dry-run", false, "Print what would be tagged without creating it")
	plugin_tagCmd.Flags().BoolP("force", "f", false, "Skip the dirty-working-tree and tag-already-exists checks")
	plugin_tagCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_tagCmd.Flags().StringP("message", "m", "", "Tag annotation message (use %s for the version)")
	plugin_tagCmd.Flags().Bool("push", false, "Push the tag to --remote after creating it")
	plugin_tagCmd.Flags().String("remote", "", "Remote to push to with --push")
	pluginCmd.AddCommand(plugin_tagCmd)
}
