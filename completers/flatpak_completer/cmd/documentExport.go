package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var documentExportCmd = &cobra.Command{
	Use:     "document-export [OPTION…] FILE ",
	Short:   "Export a file to apps",
	GroupID: "access",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(documentExportCmd).Standalone()

	documentExportCmd.Flags().BoolP("allow-delete", "d", false, "Give the app delete permissions")
	documentExportCmd.Flags().BoolP("allow-grant-permission", "g", false, "Give the app permissions to grant further permissions")
	documentExportCmd.Flags().BoolP("allow-read", "r", false, "Give the app read permissions")
	documentExportCmd.Flags().BoolP("allow-write", "w", false, "Give the app write permissions")
	documentExportCmd.Flags().StringP("app", "a", "", "Add permissions for this app")
	documentExportCmd.Flags().Bool("forbid-delete", false, "Revoke delete permissions of the app")
	documentExportCmd.Flags().Bool("forbid-grant-permission", false, "Revoke the permission to grant further permissions")
	documentExportCmd.Flags().Bool("forbid-read", false, "Revoke read permissions of the app")
	documentExportCmd.Flags().Bool("forbid-write", false, "Revoke write permissions of the app")
	documentExportCmd.Flags().BoolP("help", "h", false, "Show help options")
	documentExportCmd.Flags().BoolP("noexist", "n", false, "Don't require the file to exist already")
	documentExportCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	documentExportCmd.Flags().BoolP("transient", "t", false, "Make the document transient for the current session")
	documentExportCmd.Flags().BoolP("unique", "u", false, "Create a unique document reference")
	documentExportCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(documentExportCmd)

	carapace.Gen(documentExportCmd).FlagCompletion(carapace.ActionMap{
		"app": flatpak.ActionApplications(),
	})

	carapace.Gen(documentExportCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
