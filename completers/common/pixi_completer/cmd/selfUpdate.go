package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selfUpdateCmd = &cobra.Command{
	Use:   "self-update",
	Short: "Update pixi to the latest version or a specific version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selfUpdateCmd).Standalone()

	selfUpdateCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")
	selfUpdateCmd.Flags().Bool("dry-run", false, "Only show release notes, do not modify the binary")
	selfUpdateCmd.Flags().Bool("force", false, "Force download the desired version when not exactly same with the current. If no desired version, always replace with the latest version")
	selfUpdateCmd.Flags().String("from-url", "", "The github releases URL, useful when behind a proxy, or using custom Pixi release")
	selfUpdateCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	selfUpdateCmd.Flags().Bool("no-release-note", false, "Skip printing the release notes")
	selfUpdateCmd.Flags().Bool("offline", false, "Run without network access. Updating always requires the network, so this makes `pixi self-update` fail fast instead of attempting to connect")
	selfUpdateCmd.Flags().String("version", "", "The desired version (to downgrade or upgrade to)")
	rootCmd.AddCommand(selfUpdateCmd)

	carapace.Gen(selfUpdateCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
	})
}
