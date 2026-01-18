package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remoteAddCmd = &cobra.Command{
	Use:     "remote-add [OPTION…] NAME LOCATION",
	Short:   "Add a remote repository",
	GroupID: "repository",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remoteAddCmd).Standalone()

	remoteAddCmd.Flags().Bool("authenticator-install", false, "Autoinstall authenticator")
	remoteAddCmd.Flags().String("authenticator-name", "", "Name of authenticator")
	remoteAddCmd.Flags().String("authenticator-option", "", "Authenticator option")
	remoteAddCmd.Flags().String("collection-id", "", "Collection ID")
	remoteAddCmd.Flags().String("comment", "", "A one-line comment for this remote")
	remoteAddCmd.Flags().String("default-branch", "", "Default branch to use for this remote")
	remoteAddCmd.Flags().String("description", "", "A full-paragraph description for this remote")
	remoteAddCmd.Flags().Bool("disable", false, "Disable the remote")
	remoteAddCmd.Flags().String("filter", "", "Set path to local filter FILE")
	remoteAddCmd.Flags().Bool("from", false, "LOCATION specifies a configuration file, not the repo location")
	remoteAddCmd.Flags().String("gpg-import", "", "Import GPG key from FILE (- for stdin)")
	remoteAddCmd.Flags().BoolP("help", "h", false, "Show help options")
	remoteAddCmd.Flags().String("homepage", "", "URL for a website for this remote")
	remoteAddCmd.Flags().String("icon", "", "URL for an icon for this remote")
	remoteAddCmd.Flags().Bool("if-not-exists", false, "Do nothing if the provided remote exists")
	remoteAddCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	remoteAddCmd.Flags().Bool("no-authenticator-install", false, "Don't autoinstall authenticator")
	remoteAddCmd.Flags().Bool("no-enumerate", false, "Mark the remote as don't enumerate")
	remoteAddCmd.Flags().Bool("no-follow_redirect", false, "Don't follow the redirect set in the summary file")
	remoteAddCmd.Flags().Bool("no-gpg-verify", false, "Disable GPG verification")
	remoteAddCmd.Flags().Bool("no-use-for-deps", false, "Mark the remote as don't use for deps")
	remoteAddCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	remoteAddCmd.Flags().String("prio", "", "Set priority (default 1, higher is more prioritized)")
	remoteAddCmd.Flags().String("subset", "", "The named subset to use for this remote")
	remoteAddCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	remoteAddCmd.Flags().String("title", "", "A nice name to use for this remote")
	remoteAddCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	remoteAddCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(remoteAddCmd)

	// TODO flags
	carapace.Gen(remoteAddCmd).FlagCompletion(carapace.ActionMap{
		// "authenticator-name":   carapace.ActionValues(),
		// "authenticator-option": carapace.ActionValues(),
		// "collection-id":        carapace.ActionValues(),
		// "comment":              carapace.ActionValues(),
		"default-branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) < 2 || remoteAddCmd.Flag("from").Changed {
				return carapace.ActionValues()
			}
			return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: c.Args[1], Branches: true})
		}),
		// "description":          carapace.ActionValues(),
		"filter":     carapace.ActionFiles(),
		"gpg-import": carapace.ActionFiles(),
		// "homepage":             carapace.ActionValues(),
		// "icon":                 carapace.ActionValues(),
		// "installation":         carapace.ActionValues(),
		// "prio":                 carapace.ActionValues(),
		// "subset":               carapace.ActionValues(),
		// "title":                carapace.ActionValues(),
	})

	carapace.Gen(remoteAddCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if remoteAddCmd.Flag("from").Changed {
				return carapace.ActionFiles()
			}
			return git.ActionRepositorySearch(git.SearchOpts{}.Default())
		}),
	)
}
