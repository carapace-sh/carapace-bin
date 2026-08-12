package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload conda packages to various channels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uploadCmd).Standalone()

	uploadCmd.Flags().StringSlice("allow-insecure-host", nil, "List of hosts for which SSL certificate verification should be skipped")
	uploadCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")
	uploadCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	uploadCmd.Flags().Bool("offline", false, "Run without network access. Uploading always requires the network, so this makes `pixi upload` fail fast instead of attempting to connect")
	rootCmd.AddCommand(uploadCmd)

	carapace.Gen(uploadCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
	})
}
