package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var remoteModifyCmd = &cobra.Command{
	Use:     "remote-modify [OPTION…] NAME",
	Short:   "Modify a remote repository",
	GroupID: "repository",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remoteModifyCmd).Standalone()

	remoteModifyCmd.Flags().Bool("authenticator-install", false, "Autoinstall authenticator")
	remoteModifyCmd.Flags().String("authenticator-name", "", "Name of authenticator")
	remoteModifyCmd.Flags().String("authenticator-option", "", "Authenticator options")
	remoteModifyCmd.Flags().String("collection-id", "", "Collection ID")
	remoteModifyCmd.Flags().String("comment", "", "A one-line comment for this remote")
	remoteModifyCmd.Flags().String("default-branch", "", "Default branch to use for this remote")
	remoteModifyCmd.Flags().String("description", "", "A full-paragraph description for this remote")
	remoteModifyCmd.Flags().Bool("disable", false, "Disable the remote")
	remoteModifyCmd.Flags().Bool("enable", false, "Enable the remote")
	remoteModifyCmd.Flags().Bool("enumerate", false, "Mark the remote as enumerate")
	remoteModifyCmd.Flags().String("filter", "", "Set path to local filter FILE")
	remoteModifyCmd.Flags().Bool("follow-redirect", false, "Follow the redirect set in the summary file")
	remoteModifyCmd.Flags().String("gpg-import", "", "Import GPG key from FILE (- for stdin)")
	remoteModifyCmd.Flags().Bool("gpg-verify", false, "Enable GPG verification")
	remoteModifyCmd.Flags().BoolP("help", "h", false, "Show help options")
	remoteModifyCmd.Flags().String("homepage", "", "URL for a website for this remote")
	remoteModifyCmd.Flags().String("icon", "", "URL for an icon for this remote")
	remoteModifyCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	remoteModifyCmd.Flags().Bool("no-authenticator-install", false, "Don't autoinstall authenticator")
	remoteModifyCmd.Flags().Bool("no-enumerate", false, "Mark the remote as don't enumerate")
	remoteModifyCmd.Flags().Bool("no-filter", false, "Disable local filter")
	remoteModifyCmd.Flags().Bool("no-follow-redirect", false, "Don't follow the redirect set in the summary file")
	remoteModifyCmd.Flags().Bool("no-gpg-verify", false, "Disable GPG verification")
	remoteModifyCmd.Flags().Bool("no-use-for-deps", false, "Mark the remote as don't use for deps")
	remoteModifyCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	remoteModifyCmd.Flags().String("prio", "", "Set priority (default 1, higher is more prioritized)")
	remoteModifyCmd.Flags().String("subset", "", "Set a new subset to use")
	remoteModifyCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	remoteModifyCmd.Flags().String("title", "", "A nice name to use for this remote")
	remoteModifyCmd.Flags().Bool("update-metadata", false, "Update extra metadata from the summary file")
	remoteModifyCmd.Flags().String("url", "", "Set a new url")
	remoteModifyCmd.Flags().Bool("use-for-deps", false, "Mark the remote as used for dependencies")
	remoteModifyCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	remoteModifyCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(remoteModifyCmd)

	// TODO flags
	carapace.Gen(remoteModifyCmd).FlagCompletion(carapace.ActionMap{
		// "authenticator-name":   carapace.ActionValues(),
		// "authenticator-option": carapace.ActionValues(),
		// "collection-id":        carapace.ActionValues(),
		// "comment":              carapace.ActionValues(),
		// "default-branch":       carapace.ActionValues(),
		// "description":          carapace.ActionValues(),
		// "filter":               carapace.ActionValues(),
		// "gpg-import":           carapace.ActionValues(),
		// "homepage":             carapace.ActionValues(),
		// "icon":                 carapace.ActionValues(),
		// "installation":         carapace.ActionValues(),
		// "prio":                 carapace.ActionValues(),
		// "subset":               carapace.ActionValues(),
		// "title":                carapace.ActionValues(),
		// "url":                  carapace.ActionValues(),
	})

	// TODO positional
}
