package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selfUpdateCmd = &cobra.Command{
	Use:     "self-update",
	Short:   "Updates composer.phar to the latest version",
	Aliases: []string{"selfupdate"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selfUpdateCmd).Standalone()

	selfUpdateCmd.Flags().Bool("1", false, "Force an update to the stable channel, but only use 1.x versions")
	selfUpdateCmd.Flags().Bool("2", false, "Force an update to the stable channel, but only use 2.x versions")
	selfUpdateCmd.Flags().Bool("2.2", false, "Force an update to the stable channel, but only use 2.2.x LTS versions")
	selfUpdateCmd.Flags().Bool("clean-backups", false, "Delete old backups during an update. This makes the current version of composer the only backup available after the update")
	selfUpdateCmd.Flags().Bool("no-progress", false, "Do not output download progress.")
	selfUpdateCmd.Flags().Bool("preview", false, "Force an update to the preview channel")
	selfUpdateCmd.Flags().BoolP("rollback", "r", false, "Revert to an older installation of composer")
	selfUpdateCmd.Flags().Bool("set-channel-only", false, "Only store the channel as the default one and then exit")
	selfUpdateCmd.Flags().Bool("snapshot", false, "Force an update to the snapshot channel")
	selfUpdateCmd.Flags().Bool("stable", false, "Force an update to the stable channel")
	selfUpdateCmd.Flags().Bool("update-keys", false, "Prompt user for a key update")

	rootCmd.AddCommand(selfUpdateCmd)
}
