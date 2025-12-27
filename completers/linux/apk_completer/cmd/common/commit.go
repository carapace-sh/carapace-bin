package common

import "github.com/spf13/cobra"

func AddCommitFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("clean-protected", false, "Do not create .apk-new files in configuration directories")
	cmd.Flags().Bool("initramfs-diskless-boot", false, "Used by initramfs when it's recreating root tmpfs")
	cmd.Flags().Bool("no-commit-hooks", false, "Skip pre/post hook scripts")
	cmd.Flags().Bool("no-scripts", false, "Do not execute any scripts")
	cmd.Flags().Bool("overlay-from-stdin", false, "Read list of overlay files from stdin")
	cmd.Flags().BoolP("simulate", "s", false, "Simulate the requested operation without making any changes")
}
