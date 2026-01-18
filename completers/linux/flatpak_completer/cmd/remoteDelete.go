package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var remoteDeleteCmd = &cobra.Command{
	Use:     "remote-delete [OPTION…] NAME",
	Short:   "Delete a remote repository",
	GroupID: "repository",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remoteDeleteCmd).Standalone()

	remoteDeleteCmd.Flags().Bool("force", false, "Remove remote even if in use")
	remoteDeleteCmd.Flags().BoolP("help", "h", false, "Show help options")
	remoteDeleteCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	remoteDeleteCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	remoteDeleteCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	remoteDeleteCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	remoteDeleteCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(remoteDeleteCmd)

	// TODO flag

	// TODO positonal
}
